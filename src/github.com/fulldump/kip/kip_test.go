package kip

import (
	. "gopkg.in/check.v1"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (w *World) Test_DaoHappyPath(c *C) {

	john := w.Users.Create()

	value := john.Value.(*User)

	value.Id = bson.NewObjectId()
	value.Name = "John"
	value.Age = 16
	value.Single = true

	w.Users.Insert(john)

	// Check
	item := &User{}
	w.Database.C(w.Users.Collection.Name).Find(bson.M{
		"_id": value.Id,
	}).One(item)

	c.Assert(value, DeepEquals, item)
}

func (w *World) Test_DaoInsertTwice(c *C) {

	john := w.Users.Create()

	c.Assert(w.Users.Insert(john), IsNil)
	c.Assert(w.Users.Insert(john), NotNil)
}

func (w *World) Test_DaoCallbackOnCreate(c *C) {

	// Define collection
	w.Kip.Define(&Collection{
		Name: "Users2",
		OnCreate: func() interface{} {
			return &User{
				Id:   bson.NewObjectId(),
				Name: "default name",
			}
		},
	})

	// Instantiate Dao
	users := w.Kip.NewDao("Users2", w.Database)

	// Creae user
	john := users.Create()

	c.Assert(john.Value.(*User).Name, Equals, "default name")

}

/**
 * Define two times the same collection should panic
 */
func (w *World) Test_KipDefineTwice(c *C) {

	// Capture panic
	defer func() {
		recover()
	}()

	// Duplicate definition
	w.Kip.Define(&Collection{
		Name: "Users",
	})

	c.Error("A panic should be thrown when a collection is defined twice")
}

func (w *World) Test_get_id_map(c *C) {

	id := get_id(map[string]interface{}{
		"_id": 123,
	})
	c.Assert(id, Equals, 123)

	oid := bson.NewObjectId()
	id = get_id(map[string]interface{}{
		"_id": oid,
	})
	c.Assert(id, Equals, oid)

}

func (w *World) Test_get_id_bson(c *C) {

	id := get_id(bson.M{
		"_id": 456,
	})

	c.Assert(id, Equals, 456)
}

func (w *World) Test_get_id_struct_tag(c *C) {

	type A struct {
		ID int `bson:"_id nocomment"`
	}

	id := get_id(&A{
		ID: 789,
	})

	c.Assert(id, Equals, 789)
}

func (w *World) Test_get_id_struct_fieldname(c *C) {

	type A struct {
		Id int
	}

	id := get_id(&A{
		Id: 789,
	})

	c.Assert(id, Equals, 789)
}

func (w *World) Test_DefinitionEnsureIndex(c *C) {

	w.Kip.Define(&Collection{
		Name: "IndexedUsers",
		OnCreate: func() interface{} {
			return &User{
				Id:    bson.NewObjectId(),
				Name:  "default name",
				Email: "user@email.com",
			}
		},
		Indexes: []mgo.Index{
			mgo.Index{
				Key:        []string{"email"},
				Unique:     true,
				DropDups:   true,
				Background: true, // See notes.
				Sparse:     true,
			},
		},
	})

	users := w.Kip.NewDao("IndexedUsers", w.Database)

	user1 := users.Create()
	err1 := users.Insert(user1)
	c.Assert(err1, IsNil)

	user2 := users.Create()
	err2 := users.Insert(user2)
	c.Assert(err2, NotNil)

	result := []interface{}{}
	users.Database.C(users.Collection.Name).Find(bson.M{}).All(&result)

	c.Assert(len(result), Equals, 1)
}
