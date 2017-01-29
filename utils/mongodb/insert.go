package mongodb

import (
    "log"
    "../../config"
)

func Insert(document string, data M) (error){
    c := Session.DB(config.Config.Mongodb.Database).C(document)
    err := c.Insert(data)
    if err != nil {
        log.Fatal("Error while inserting data on " + document, err)
    }
    return err
}