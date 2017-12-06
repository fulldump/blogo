package kip

import (
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (w *World) Test_Item_Delete(c *C) {

	john := w.Users.Create()
	john.Save()
	john.Delete()

	// Check
	item := &User{}
	err := w.Database.C(w.Users.Collection.Name).Find(bson.M{
		"_id": john.GetId(),
	}).One(item)

	c.Assert(err.Error(), Equals, "not found")

}

func (w *World) Test_Item_Save(c *C) {

	john := w.Users.Create()
	john.Save()

	// Check
	item := &User{}
	err := w.Database.C(w.Users.Collection.Name).Find(bson.M{
		"_id": john.GetId(),
	}).One(item)

	c.Assert(err, IsNil)

}

func (w *World) Test_Item_Save_Twice(c *C) {

	john := w.Users.Create()

	c.Assert(john.Save(), IsNil)
	c.Assert(john.Save(), NotNil)
}

func (w *World) Test_Item_Update_Set(c *C) {

	john := w.Users.Create()
	john.Save()

	// Do a patch
	john.Patch(&Patch{
		Operation: "set",
		Key:       "name",
		Value:     "New name",
	})
	john.Save()

	// Check
	item := &User{}
	w.Database.C(w.Users.Collection.Name).Find(bson.M{
		"_id": john.GetId(),
	}).One(item)

	c.Assert(item.Name, Equals, "New name")

}

func (w *World) Test_Item_Update_AddToSet(c *C) {

	john := w.Users.Create()
	john.Save()

	// Do a patch
	john.Patch(&Patch{
		Operation: "add_to_set",
		Key:       "friends",
		Value:     "Fulano",
	})
	john.Patch(&Patch{
		Operation: "add_to_set",
		Key:       "friends",
		Value:     "Mengano",
	})
	john.Patch(&Patch{
		Operation: "add_to_set",
		Key:       "friends",
		Value:     "Fulano",
	})
	john.Save()

	// Check
	item := &User{}
	w.Database.C(w.Users.Collection.Name).Find(bson.M{
		"_id": john.GetId(),
	}).One(item)

	c.Assert(item.Friends, DeepEquals, []string{"Fulano", "Mengano"})

}

func (w *World) Test_Item_Update_AddToSet_Multi(c *C) {

	john := w.Users.Create()
	john.Save()

	// Do a patch
	john.Patch(&Patch{
		Operation: "add_to_set",
		Key:       "colors",
		Value:     "Blue",
	})
	john.Patch(&Patch{
		Operation: "add_to_set",
		Key:       "friends",
		Value:     "Mengano",
	})
	john.Save()

	// Check
	item := &User{}
	w.Database.C(w.Users.Collection.Name).Find(bson.M{
		"_id": john.GetId(),
	}).One(item)

	c.Assert(item.Friends, DeepEquals, []string{"Mengano"})
	c.Assert(item.Colors, DeepEquals, []string{"Blue"})

}

func (w *World) Test_Item_Update_Multi(c *C) {

	john := w.Users.Create()
	john.Save()

	// Do a patch
	john.Patch(&Patch{
		Operation: "add_to_set",
		Key:       "colors",
		Value:     "Blue",
	})
	john.Patch(&Patch{
		Operation: "set",
		Key:       "name",
		Value:     "My New Name",
	})
	john.Save()

	// Check
	item := &User{}
	w.Database.C(w.Users.Collection.Name).Find(bson.M{
		"_id": john.GetId(),
	}).One(item)

	c.Assert(item.Name, DeepEquals, "My New Name")
	c.Assert(item.Colors, DeepEquals, []string{"Blue"})

}

func (w *World) Test_Item_Where_LockAcquired(c *C) {

	john_item := w.Users.Create()
	john := john_item.Value.(*User)
	john.Name = "John"
	john.Age = 25
	john.Lock = false
	john_item.Save()

	// Try to get lock
	john_item.Patch(&Patch{
		Operation: "set",
		Key:       "lock",
		Value:     26,
	})
	john_item.Where(bson.M{
		"lock": false,
	})
	err := john_item.Save()

	c.Assert(err, IsNil)

	// Check
	item := &User{}
	w.Database.C(w.Users.Collection.Name).Find(bson.M{
		"_id": john.Id,
	}).One(item)

	c.Assert(item.Lock, DeepEquals, true)

}

func (w *World) Test_Item_Where_LockNotAcquired(c *C) {

	john_item := w.Users.Create()
	john := john_item.Value.(*User)
	john.Name = "John"
	john.Age = 25
	john.Lock = true
	john_item.Save()

	// Try to get lock
	john_item.Patch(&Patch{
		Operation: "set",
		Key:       "lock",
		Value:     26,
	})
	john_item.Where(bson.M{
		"lock": false,
	})
	err := john_item.Save()

	c.Assert(err, NotNil)

	// Check
	item := &User{}
	w.Database.C(w.Users.Collection.Name).Find(bson.M{
		"_id": john.Id,
	}).One(item)

	c.Assert(item.Lock, DeepEquals, true)

}
