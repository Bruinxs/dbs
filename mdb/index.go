package mdb

import (
	"gopkg.in/mgo.v2"
)

func EnsureOrDropIndex(indexs map[string]map[string]*mgo.Index, C func(name string) *mgo.Collection) error {
	for cname, entity := range indexs {
		oldIndexs, err := C(cname).Indexes()
		exitst := map[string]bool{}

		//deal old index
		for _, index := range oldIndexs {
			if index.Name == "_id_" {
				continue
			}
			if entity[index.Name] == nil {
				err = C(cname).DropIndexName(index.Name)
				if err != nil {
					return err
				}
			} else {
				exitst[index.Name] = true
			}
		}

		//ensure new index
		for name, index := range entity {
			if !exitst[name] {
				index.Name = name
				err = C(cname).EnsureIndex(*index)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
