package mongodb

import (
	"gopkg.in/mgo.v2/bson"
)

// M defines a generic type for mongodb queries.
type M bson.M

// Query defines an object to easily customise options to query functions
type Query struct {
	// Always required
	Document string // Document to qury in
	Find     M      // Query options

	// Required for inserting and updating and optionaly for deleting
	Data interface{}

	// Required for single finds
	Result interface{}

	// Required for mutiple finds
	Results []interface{}

	// Optional
	Limit int
	Skip  int
	Sort  []string
	Hint  []string
}

// NewQuery returns a new Query object with Find and document initialized
func NewQuery(document string, find M) (Query, error) {
	return Query{Document: document, Find: find}, nil
}
