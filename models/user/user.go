package datamodels

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// TODO incorporate this into a dedicated util.
type user struct {
	Name  string        `json:"name" bson:"name"`
	Photo string        `json:"photo" bson:"photo"`
	RA    string        `json:"ra" bson:"ra"`
	UID   string        `json:"uid" bson:"uid"`
	_id   bson.ObjectId `json:"-" bson:"_id,omitempty"`

	Account struct {
		Email    string `json:"email" bson:"email"`
		Password []byte `json:"-" bson:"pwd"`
		Salt     []byte `json:"-" bson:"salt"`
	}

	Personal struct {
		Birth        time.Time `json:"birth" bson:"birth"`
		Phone        string    `json:"phone" bson:"phone"`
		Drinks       int       `json:"drinks" bson:"drinks"`
		Relationship struct {
			Sex    string `json:"sex" bson:"sex"`
			Gender string `json:"gender" bson:"gender"`
			Status string `json:"status" bson:"status"`
		}
		Origin struct {
			Country string `json:"country" bson:"country"`
			State   string `json:"state" bson:"state"`
			City    string `json:"city" bson:"city"`
			School  string `json:"school" bson:"school"`
		}
	}

	University struct {
		Grad   string `json:"grad" bson:"grad"`
		Reason struct {
			Unicamp string `json:"unicamp" bson:"unicamp"`
			Grad    string `json:"grad" bson:"grad"`
		}
	}

	Work struct {
		Job     string `json:"job" bson:"job"`
		Company string `json:"company" bson:"company"`
	}

	Social map[string]string `json:"social" bson:"social"`

	Hobbies struct {
		Languages []string `json:"languages" bson:"languages"`

		Games struct {
			Favorite map[string]string `json:"favorite" bson:"favorite"`
			Consoles map[string]string `json:"consoles" bson:"consoles"`
			PC       bool              `json:"pc" bson:"pc"`
			Extra    string            `json:"extra" bson:"extra"`
			Accounts map[string]string `json:"accounts" bson:"accounts"`
		}

		Sports struct {
			Plays []string `json:"plays" bson:"plays"`
			Likes []string `json:"likes" bson:"likes"`
			Teams []string `json:"teams" bson:"teams"`
			Extra string   `json:"extra" bson:"extra"`
		}

		TV struct {
			Genres   []string          `json:"genres" bson:"genres"`
			Movies   []string          `json:"movies" bson:"movies"`
			Series   []string          `json:"series" bson:"series"`
			Cartoons []string          `json:"cartoons" bson:"cartoons"`
			Animes   []string          `json:"animes" bson:"animes"`
			Accounts map[string]string `json:"accounts" bson:"accounts"`
		}

		Programming struct {
			Languages []string `json:"languages" bson:"languages"`
			OS        []string `json:"languages" bson:"languages"`
		}

		Music struct {
			Genres []string `json:"genres" bson:"genres"`
			Bands  []string `json:"bands" bson:"bands"`
			Albuns []string `json:"albuns" bson:"albuns"`
			Songs  []string `json:"songs" bson:"songs"`
		}
	}

	Friends []struct {
		RA  string `json:"ra" bson:"ra"`
		UID string `json:"uid" bson:"uid"`
	}
}
