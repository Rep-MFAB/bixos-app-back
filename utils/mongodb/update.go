package mongodb

import (
	"errors"
	"log"

	"github.com/seijihirao/bixos-app-back/config"
)

// Update updates the fields of the indicated document.
func Update(query Query) error {
	if query.Find == nil || query.Document == "" || query.Data == nil {
		return errors.New("please provide a valid query")
	}

	c := Session.DB(config.Config.Mongodb.Database).C(query.Document)

	err := c.Update(query.Find, M{"$set": query.Data})
	if err != nil {
		log.Fatal("Error while updating data on "+query.Document, err)
	}
	return err
}
