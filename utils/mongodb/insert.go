package mongodb

import (
    "gopkg.in/mgo.v2/bson"
    "../../config"
    "log"
)

func Insert(document string, data bson.M) {
    c := Session.DB(config.Config.Mongodb.Database).C(document)
    err := c.Insert(data)
    if err != nil {
        log.Fatal("Error while inserting data on " + document, err)
    }
}