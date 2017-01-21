package endpoints

import (
    "net/http"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
    "../utils/mongodb"
)

func init(){
    http.HandleFunc("/users/register", func(w http.ResponseWriter, r *http.Request) {
        if r.Body == nil {
            http.Error(w, "Please send a request body", 400)
            return
        }

        var user bson.M
        err := json.NewDecoder(r.Body).Decode(&user)
        if err != nil {
            http.Error(w, err.Error(), 400)
            return
        }

        mongodb.Insert("Users", user)        
    })
}