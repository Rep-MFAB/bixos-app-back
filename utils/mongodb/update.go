package mongodb

import (
	"log"
	"time"

	"github.com/seijihirao/bixos-app-back/config"
)

func Update(document string, query M, data M) error {
	data["updatedAt"] = time.Now()

	c := Session.DB(config.Config.Mongodb.Database).C(document)
	err := c.Update(query, M{"$set": data})
	if err != nil {
		log.Fatal("Error while updating data on "+document, err)
	}
	return err
}
