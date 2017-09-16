package mdb

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

//GenerateSimpleDb return mongodb database provide by url
func GenerateSimpleDb(url string) (Db, error) {
	info, err := mgo.ParseURL(url)
	if err != nil {
		return nil, err
	}
	if info.Database == "" {
		return nil, fmt.Errorf("the url %v is illegal, must provide database name", url)
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		return nil, err
	}
	return session.DB(""), nil
}
