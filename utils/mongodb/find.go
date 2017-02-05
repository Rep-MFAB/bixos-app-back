package mongodb

import (
	"errors"

	"github.com/seijihirao/bixos-app-back/config"
)

// FindOne searches for a single entry as indicated by query.Find in mongodb filling the query.Output field with what is found
func FindOne(query Query) error {
	if query.Find == nil || query.Document == "" || query.Result == nil {
		return errors.New("Please provide a valid query.")
	}

	c := Session.DB(config.Config.Mongodb.Database).C(query.Document)

	q := c.Find(query.Find)
	if query.Hint != nil {
		q.Hint(query.Hint...)
	}
	if query.Sort != nil {
		q.Sort(query.Sort...)
	}
	err := q.One(query.Result)

	return err
}

// FindAll searches for  all the entries
func FindAll(query Query) error {
	if query.Find == nil || query.Document == "" || query.Results == nil {
		return errors.New("Please provide a valid query.")
	}

	c := Session.DB(config.Config.Mongodb.Database).C(query.Document)

	q := c.Find(query.Find)
	if query.Hint != nil {
		q.Hint(query.Hint...)
	}
	if query.Sort != nil {
		q.Sort(query.Sort...)
	}
	if query.Limit <= 0 {
		q.Skip(query.Skip)
	} else {
		q.Skip(query.Skip).Limit(query.Limit)
	}

	err := q.All(query.Results)

	return err
}
