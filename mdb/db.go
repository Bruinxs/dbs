package mdb

import (
	"gopkg.in/mgo.v2"
)

var defaultMDB *mgo.Session

func AddMongodb(url string) {
	if defaultMDB != nil {
		panic("default mongodb db had be initialized")
	}
	var err error
	defaultMDB, err = mgo.Dial(url)
	if err != nil {
		panic(err)
	}
}

func C(name string) *mgo.Collection {
	if defaultMDB == nil {
		panic("default mongodb db is not initialization")
	}
	return defaultMDB.DB("").C(name)
}
