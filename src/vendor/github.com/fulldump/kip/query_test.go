package kip

import (
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (w *World) Stub_20AgedUsers() {
	for i := 10; i >= -10; i-- {
		user_item := w.Users.Create()

		user := user_item.Value.(*User)
		user.Age = i

		user_item.Save()
	}
}

func (w *World) Test_Query(c *C) {

	w.Stub_20AgedUsers()

	filter := bson.M{
		"age": bson.M{
			"$gte": 1,
		},
	}

	query := w.Users.Find(filter).Sort("age").Skip(3).Limit(2)

	// Check Count
	{
		n, err := query.Count()

		c.Assert(n, Equals, 2)
		c.Assert(err, IsNil)
	}

	// Check One
	{
		result := bson.M{}
		err := query.One(&result)

		c.Assert(result["age"], DeepEquals, 4)
		c.Assert(err, IsNil)
	}

	// Check All
	{
		result := []bson.M{}
		err := query.All(&result)

		c.Assert(len(result), Equals, 2)

		c.Assert(result[0]["age"], DeepEquals, 4)
		c.Assert(result[1]["age"], DeepEquals, 5)

		c.Assert(err, IsNil)
	}

	// Check ForEach
	{

		results := []int{}

		query.ForEach(func(item *Item) {
			age := item.Value.(*User).Age
			results = append(results, age)
		})

		c.Assert(results, DeepEquals, []int{4, 5})
	}

}

func (w *World) Test_Query_x10(c *C) {
	for i := 0; i < 10; i++ {
		w.Users.Delete(bson.M{})
		w.Test_Query(c)
	}
}

func (w *World) Test_Query_FindWithFilter(c *C) {

	w.Stub_20AgedUsers()

	n, err := w.Users.Find(bson.M{"age": 5}).Count()

	c.Assert(n, Equals, 1)

	c.Assert(err, IsNil)

}

func (w *World) Test_Query_FindWithoutFilter(c *C) {

	w.Stub_20AgedUsers()

	n, err := w.Users.Find(nil).Count()

	c.Assert(n, Equals, 21)

	c.Assert(err, IsNil)

}

func (w *World) Test_Query_Skip(c *C) {

	w.Stub_20AgedUsers()

	n, err := w.Users.Find(nil).Skip(20).Count()

	c.Assert(n, Equals, 1)

	c.Assert(err, IsNil)

}

func (w *World) Test_Query_Limit(c *C) {

	w.Stub_20AgedUsers()

	n, err := w.Users.Find(nil).Limit(20).Count()

	c.Assert(n, Equals, 20)

	c.Assert(err, IsNil)

}

func (w *World) Test_Query_Projection(c *C) {

	w.Stub_20AgedUsers()

	result := bson.M{}

	err := w.Users.Find(nil).Select(bson.M{"age": 1, "_id": -1}).One(&result)

	expected := bson.M{
		"_id": result["_id"],
		"age": result["age"],
	}

	c.Assert(result, DeepEquals, expected)

	c.Assert(err, IsNil)

}
