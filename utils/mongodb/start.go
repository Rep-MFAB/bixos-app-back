package mongodb

import (
	"log"

	"github.com/seijihirao/bixos-app-back/config"
	"gopkg.in/mgo.v2"
)

// Session exports session to use in all mongodb operations
var Session *mgo.Session

// Start intializes a new mongodb connection and returns its session information.
func Start() *mgo.Session {
	var err error
	Session, err = mgo.Dial(config.Config.Mongodb.Host)
	if err != nil {
		log.Fatal("Starting Mongo: ", err)
	}
	// Optional. Switch the session to a monotonic behavior.
	Session.SetMode(mgo.Primary, true)

	log.Printf("Mongodb started on host: \033[95m%s\033[0m", config.Config.Mongodb.Host)

	return Session
}
