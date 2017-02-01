package mongodb

import (
	"log"
	"time"

	"../../config"
)

func Insert(document string, data M) error {
	data["removedAt"] = false
	data["updatedAt"] = time.Now()

	c := Session.DB(config.Config.Mongodb.Database).C(document)
	err := c.Insert(data)
	if err != nil {
		log.Fatal("Error while inserting data on "+document, err)
	}
	return err
}
