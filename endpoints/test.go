package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/seijihirao/bixos-app-back/config"
)

func init() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(config.Config)
	})
}
