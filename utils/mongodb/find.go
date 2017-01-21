package mongodb

import (
    "gopkg.in/mgo.v2/bson"
    "../../config"
    "log"
)
 
func FindOne(document string, what bson.M) bson.M{
    c := Session.DB(config.Config.Mongodb.Database).C(document)
    result := bson.M{}

    err := c.Find(what).One(&result)
    if err != nil {
        log.Fatal("Error while finding data from " + document, err)
    }

    return result
}
