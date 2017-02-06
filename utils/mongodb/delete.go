package mongodb

import (
	"errors"
	"time"

	"github.com/seijihirao/bixos-app-back/config"
)

// Delete updates the database with a removedAt tag signaling that this entry should not be used anymore.
// If query.HardDelete is true, it removes the data from db, else it just updates the query with a removedAt tag.
func Delete(query Query) error {
	if query.Find == nil || query.Document == "" {
		return errors.New("please provide a valid query")
	}
	var err error
	if query.HardDelete == true {
		c := Session.DB(config.Config.Mongodb.Database).C(query.Document)
		err = c.Remove(query.Find)
	} else {
		query.Data = &M{"removedAt": time.Now()}
		err = Update(query)
	}

	return err
}
