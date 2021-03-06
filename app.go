package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/seijihirao/bixos-app-back/config"
	_ "github.com/seijihirao/bixos-app-back/endpoints"
	"github.com/seijihirao/bixos-app-back/utils/mongodb"
)

func main() {
	session := mongodb.Start()
	defer session.Close()

	log.Printf("Serving app on host: \033[95mlocalhost:%d\033[0m", config.Config.Server.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Server.Port), nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while listenning to port %d: ", config.Config.Server.Port), err)
	}

}
