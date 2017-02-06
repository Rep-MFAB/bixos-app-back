package mongodb

import (
	"errors"
	"log"

	"github.com/seijihirao/bixos-app-back/config"
)

// Insert adds data to the mongodb
func Insert(query Query) error {
	if query.Document == "" || query.Data == nil {
		return errors.New("Please provide a valid query")
	}

	c := Session.DB(config.Config.Mongodb.Database).C(query.Document)
	err := c.Insert(query.Data)
	if err != nil {
		log.Fatal("Error while inserting data on "+query.Document, err)
	}
	return err
}
