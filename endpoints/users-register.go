package endpoints

import (
	"fmt"
	"net/http"

	"../utils/json"
	"../utils/mongodb"
)

func init() {
	http.HandleFunc("/users/register", func(w http.ResponseWriter, r *http.Request) {
		var user = json.Get(w, r)
		if user == nil {
			return
		}

		// Searching if user already exists
		err, dbUser := mongodb.FindOne("Users", mongodb.M{"email": user["email"]})

		if err != nil {
			http.Error(w, "Database Error", 500)
			return
		} else if dbUser != nil {
			http.Error(w, "User already exists", 409)
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
