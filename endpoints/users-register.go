package endpoints

import (
    "net/http"
    "../utils/mongodb"
    "../utils/json"
    "fmt"
)

func init(){
    http.HandleFunc("/users/register", func(w http.ResponseWriter, r *http.Request) {
        var user = json.Get(w, r)
        if user == nil {
            return
        }

        // Searching if user already exists
        err, _ := mongodb.FindOne("Users", mongodb.M{"email": user["email"]})
        
        if err == nil{
            http.Error(w, "User already exists", 409)
            return
        }else if err.Error() != "not found" {
            http.Error(w, "Database Error", 500)
            return
        }

        // Inserting user
        err = mongodb.Insert("Users", user)
        if err != nil {
            http.Error(w, "Database Error", 500)
            return
        }

        fmt.Fprint(w, "Success")
    })
}