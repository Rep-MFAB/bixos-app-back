package mongodb

import (
	"errors"
	"time"

	"github.com/seijihirao/bixos-app-back/config"
)

// Delete updates the database with a removedAt tag signaling that this entry should not be used anymore.
func Delete(query Query) error {
	if query.Find == nil || query.Document == "" {
		return errors.New("Please provide a valid query.")
	}
	var err error
	if query.Data == nil {
		c := Session.DB(config.Config.Mongodb.Database).C(query.Document)
		err = c.Remove(query.Find)
	} else {
		query.Data = &M{"removedAt": time.Now()}
		err = Update(query)
	}

	return err
}
