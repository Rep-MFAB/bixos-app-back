package config

import (
    "encoding/json"
    "os"
    "log"
)

type Configuration struct {
    Mongodb     struct {
        Host        string
        Database    string
    }
}

func Get() (Configuration){
    file, _ := os.Open("./config/dev.json")
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        log.Fatal("Json Decode: ", err)
    }
    return configuration
}

var Config = Get()