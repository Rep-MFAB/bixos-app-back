package mongodb

import (
	"github.com/seijihirao/bixos-app-back/config"
)

func FindOne(document string, query M) (error, M) {
	query["removedAt"] = false

	c := Session.DB(config.Config.Mongodb.Database).C(document)
	result := M{}
	err := c.Find(query).One(&result)
	if err.Error() != "not found" {
		return nil, nil
	}
	return err, result
}

func FindAll(document string, query M, limit int, skip int) (error, []M) {
	query["removedAt"] = false

	c := Session.DB(config.Config.Mongodb.Database).C(document)
	result := []M{}

	err := c.Find(query).Skip(skip).Limit(limit).All(&result)
	if limit < 0 {
		err = c.Find(query).Skip(skip).All(&result)
	}
	return err, result
}
