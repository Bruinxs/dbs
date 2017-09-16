package mdb

import (
	"gopkg.in/mgo.v2"
)

//Db call func C to get a collection object
type Db interface {
	C(name string) *mgo.Collection
}
