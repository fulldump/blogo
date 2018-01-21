package kip

import "gopkg.in/mgo.v2"

type Collection struct {
	Name     string
	Sample   interface{}
	OnCreate OnCreate
	Indexes  []mgo.Index
}

type OnCreate func() interface{}
