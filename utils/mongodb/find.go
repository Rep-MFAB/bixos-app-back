package mongodb

import (
    "../../config"
)
 
func FindOne(document string, what M) (error, M){
    c := Session.DB(config.Config.Mongodb.Database).C(document)
    result := M{}

    err := c.Find(what).One(&result)
    return err, result
}
