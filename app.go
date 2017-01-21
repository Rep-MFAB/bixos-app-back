package main

import (
    _ "./endpoints"
    "./utils/mongodb"
    "log"
    "net/http"
)

func main(){
    session := mongodb.Start()
    defer session.Close()

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("Error while listenning to port 8080: ", err)
    }
}