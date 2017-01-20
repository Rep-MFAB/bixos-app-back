package models

import (
    "labix.org/v2/mgo/bson"
    "time"
)

type Configuration struct {
    ID      bson.ObjectId
    Name    string
    RA      string
    Account struct {
        Email   string
        Pass    string
        UID     string
    }
    Personal struct {
        Birth   time.Time
        Phone   string
        Drinks  int
        Relationship    struct {
            Sex     string
            Gender  string
            Status  string
        }
        Origin  struct {
            Country string
            State   string
            City    string
            School  string
        }
    }
    School  struct {
        Grad        string
        Languages   []string
        Reason      struct {
            Unicamp     string
            Grad        string
        }
    }
    Work    struct {
        Job     string
        Company string
    }
    Social  struct {
        facebook    string
        twitter     string
        instagram   string
        snapchat    string
        google      string
        others      []string
    }
    Hobbies struct {
        Games   struct {
            Favorites   []string
            Consoles    []string
            PC          bool
            Accounts    struct {
                Steam       string
                Origin      string
                PSN         string
                Live        string
                Battlenet   string
                Uplay       string
                others      []string
            }
            Extra       string
        }
        Sports  struct {
            Plays       []string
            Likes       []string
            Teams       []string
            Extra       string
        }
        TV      struct {
            Genres      []string
            Movies      []string
            Series      []string
            Cartoons    []string
            Disney      []string
            Animes      []string
            Accounts    struct {
                IMDB        string
                Tvshowtime  string
                Anidb       string
                Myanimelist string
                Anilist     string
                Others      []string  
            }
            Extra       string
        }
        Programming struct {
            Languages   []string
            OS          []string
            Extra       string
        }
        Music   struct {
            Genres      []string
            Bands       []string
            Albuns      []string
            Songs       []string
            Instruments []string
            Extra       string
        }
        Extra   string
    }
    Friends []struct {
        RA  string
        UID string
    }
}