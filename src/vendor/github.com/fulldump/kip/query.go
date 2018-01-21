package kip

import (
	mgo "gopkg.in/mgo.v2"
)

type Query struct {
	dao *Dao

	limit    *int
	selector interface{}
	projection interface{}
	skip     *int
	sort     []string
	snapshot bool
}

func (q *Query) Limit(n int) *Query {
	q.limit = &n
	return q
}

func (q *Query) Select(projection interface{}) *Query {
	q.projection = projection
	return q
}

func (q *Query) Skip(n int) *Query {
	q.skip = &n
	return q
}

func (q *Query) Snapshot() *Query {
	q.snapshot = true
	return q
}

func (q *Query) Sort(fields ...string) *Query {
	q.sort = fields
	return q
}

// Finalizers
func (q *Query) All(result interface{}) error {

	query, db := q.buildQuery()
	defer db.Close()

	return query.All(result)
}

func (q *Query) Count() (n int, err error) {

	query, db := q.buildQuery()
	defer db.Close()

	return query.Count()
}

func (q *Query) Iter() (*mgo.Iter, *Database) {

	query, db := q.buildQuery()

	return query.Iter(), db
}

func (q *Query) ForEach(f func(*Item)) error {

	query, db := q.buildQuery()
	defer db.Close()

	i := query.Iter()

	item := q.dao.Create()
	for i.Next(item.Value) {
		item.saved = true

		f(item)

		item = q.dao.Create()
	}

	return i.Close()
}

func (q *Query) One(result interface{}) error {

	query, db := q.buildQuery()
	defer db.Close()

	return query.One(result)
}

// Internal helper
func (q *Query) buildQuery() (query *mgo.Query, db *Database) {

	db = q.dao.Database.Clone()
	// NOTE: do not defer db.Close()

	query = db.C(q.dao.Collection.Name).Find(q.selector)

	if nil != q.limit {
		query = query.Limit(*q.limit)
	}

	if nil != q.skip {
		query = query.Skip(*q.skip)
	}

	if q.snapshot {
		query = query.Snapshot()
	}

	if nil != q.sort {
		query = query.Sort(q.sort...)
	}

	if nil != q.projection {
		query = query.Select(q.projection)
	}

	return
}
