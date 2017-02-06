package endpoints

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/seijihirao/bixos-app-back/models/usermodel"
	"github.com/seijihirao/bixos-app-back/utils/crypt"
	"github.com/seijihirao/bixos-app-back/utils/json"
	"github.com/seijihirao/bixos-app-back/utils/mongodb"
)

func init() {
	http.HandleFunc("/users/register", func(w http.ResponseWriter, r *http.Request) {
		// Check if the path is strictly as specified.
		if r.URL.Path != "/users/register" {
			http.Error(w, "Endpoint not found.", 404)
			return
		}

		// Checks for the request method
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", "POST")
			http.Error(w, "Invalid request method.", 405)
			return
		}

		// Retrieve request variables into a map.
		var user = json.Get(w, r)
		if user == nil {
			http.Error(w, "No request body provided.", 400)
			return
		}

		// Creates a new user to save on the db
		var newUser usermodel.User

		// Adds the email and checks if it is correctly formated.
		if val, ok := user["email"]; ok {
			if email, ok := val.(string); ok {
				e, err := mail.ParseAddress(email)
				if err != nil || e.Name != "" {
					http.Error(w, "Invalid email.", 400)
					return
				}
				newUser.Account.Email = e.Address
			} else {
				http.Error(w, "Invalid data.", 400)
				return
			}
		} else {
			http.Error(w, "Missing formulary information.", 400)
			return
		}

		// Searching if user already exists
		query := mongodb.NewQuery("Users", mongodb.M{"account.email": newUser.Email})
		var dbUser usermodel.User
		query.Result = &dbUser
		err := mongodb.FindOne(query)
		if err != nil {
			if err != mongodb.ErrNotFound {
				http.Error(w, "Internal server error.", 500)
				return
			}
		} else {
			http.Error(w, "User email already exists.", 409)
			return
		}

		// Adds the hashed password with its salt.
		if val, ok := user["password"]; ok {
			if pwd, ok := val.(string); ok {
				var errPwd error
				newUser.Account.Password, newUser.Account.Salt, errPwd = crypt.Encrypt([]byte(pwd))
				if errPwd != nil {
					http.Error(w, "Internal server error.", 500)
					return
				}
			} else {
				http.Error(w, "Invalid data.", 400)
				return
			}
		} else {
			http.Error(w, "Missing formulary information.", 400)
			return
		}

		// Adds the user name.
		if val, ok := user["name"]; ok {
			if name, ok := val.(string); ok {
				newUser.Name = name
			} else {
				http.Error(w, "Invalid data.", 400)
				return
			}
		} else {
			http.Error(w, "Missing formulary information.", 400)
			return
		}

		// Adds the user ra
		if val, ok := user["ra"]; ok {
			if ra, ok := val.(string); ok {
				newUser.RA = ra
			} else {
				http.Error(w, "Invalid data.", 400)
				return
			}
		} else {
			http.Error(w, "Missing formulary information.", 400)
			return
		}

		// Generates a new 5 digits UID for the new user.
		newUser.UID = crypt.UUID(5)

		// Utilizes the same query as before, just changing the data (Oh the beauty)
		query.Data = newUser
		fmt.Println(query.Data, query.Document)

		// Inserts the new user in the Database
		err = mongodb.Insert(query)
		if err != nil {
			http.Error(w, "Internal server error.", 500)
			return
		}

		// If code gets here all is ok.
		w.WriteHeader(200)
		return
	})
}
