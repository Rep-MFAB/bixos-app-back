package json

import (
    "net/http"
    "encoding/json"
    "../mongodb"
)

func Get(w http.ResponseWriter, r *http.Request) (mongodb.M){

    if r.Body == nil {
        http.Error(w, "Please send a request body", 400)
        return nil
    }

    var object mongodb.M
    var err = json.NewDecoder(r.Body).Decode(&object)
    if err != nil {
        http.Error(w, err.Error(), 400)
        return nil
    }
    return object
}