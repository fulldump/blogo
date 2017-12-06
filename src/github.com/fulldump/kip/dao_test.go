package kip

import (
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (w *World) Test_InstanceFindById_Ok(c *C) {

	// Prepare
	john := w.Users.Create()
	w.Users.Insert(john)

	// Run
	u, err := w.Users.FindById(john.GetId())

	// Check
	c.Assert(u.Value, DeepEquals, john.Value)
	c.Assert(err, IsNil)
}

func (w *World) Test_InstanceFindById_NotFound(c *C) {

	u, err := w.Users.FindById("invented id")

	// Check
	c.Assert(u, IsNil)
	c.Assert(err, IsNil)
}

func (w *World) Test_InstanceFindOne_Ok(c *C) {

	// Prepare
	john := w.Users.Create()
	john.Value.(*User).Name = "John Snow"
	w.Users.Insert(john)

	// Run
	u, err := w.Users.FindOne(bson.M{"name": "John Snow"})

	// Check
	c.Assert(u.Value, DeepEquals, john.Value)
	c.Assert(err, IsNil)
}

func (w *World) Test_InstanceFindOne_Fail(c *C) {

	// Run
	u, err := w.Users.FindOne(bson.M{"name": "John Snow"})

	// Check
	c.Assert(u, IsNil)
	c.Assert(err, IsNil)
}

func (w *World) Test_Dao_Delete_Ok(c *C) {

	u1 := w.Users.Create()
	u1.Value.(*User).Name = "a"
	u1.Save()

	u2 := w.Users.Create()
	u2.Value.(*User).Name = "b"
	u2.Save()

	u3 := w.Users.Create()
	u3.Value.(*User).Name = "a"
	u3.Save()

	n, err := w.Users.Delete(bson.M{"name": "a"})

	c.Assert(n, Equals, 2)
	c.Assert(err, IsNil)

	objects, _ := w.Users.Find(bson.M{"name": "a"}).Count()
	c.Assert(objects, Equals, 0)

}
