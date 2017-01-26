package main

import (
    _ "./endpoints"
    "./config"
    "./utils/mongodb"
    "net/http"
    "log"
    "fmt"
)

func main(){
    session := mongodb.Start()
    defer session.Close()

    err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Server.Port), nil)
    if err != nil {
        log.Fatal(fmt.Sprintf("Error while listenning to port %d:", config.Config.Server.Port), err)
    }
}