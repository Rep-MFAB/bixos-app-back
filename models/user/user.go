package datamodels

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// TODO incorporate this into a dedicated util.
type User struct {
	Name  string        `json:"name" bson:"name"`
	Photo string        `json:"photo" bson:"photo"`
	RA    string        `json:"ra" bson:"ra"`
	UID   string        `json:"uid" bson:"uid"`
	_id   bson.ObjectId `json:"-" bson:"_id,omitempty"`

	Account struct {
		Email    string `json:"email" bson:"email"`
		Password []byte `json:"-" bson:"pwd"`
		Salt     []byte `json:"-" bson:"salt"`
	} `json:"account" bson:"account"`

	Personal struct {
		Birth        time.Time `json:"birth" bson:"birth"`
		Phone        string    `json:"phone,omitempty" bson:"phone,omitempty"`
		Drinks       int       `json:"drinks,omitempty" bson:"drinks,omitempty"`
		Relationship struct {
			Sex    string `json:"sex" bson:"sex"`
			Gender string `json:"gender,omitempty" bson:"gender,omitempty"`
			Status string `json:"status,omitempty" bson:"status,omitempty"`
		} `json:"relationship" bson:"relationship"`
		Origin struct {
			Country string `json:"country" bson:"country"`
			State   string `json:"state,omitempty" bson:"state,omitempty"`
			City    string `json:"city,omitempty" bson:"city,omitempty"`
			School  string `json:"school,omitempty" bson:"school,omitempty"`
		} `json:"origin" bson:"origin"`
	} `json:"personal" bson:"personal"`

	University struct {
		Grad   string `json:"grad" bson:"grad"`
		Reason struct {
			Unicamp string `json:"unicamp,omitempty" bson:"unicamp,omitempty"`
			Grad    string `json:"grad,omitempty" bson:"grad,omitempty"`
		} `json:"reason" bson:"reason"`
	} `json:"university" bson:"university"`

	Work struct {
		Job     string `json:"job" bson:"job"`
		Company string `json:"company" bson:"company"`
	} `json:"work,omitempty" bson:"work,omitempty"`

	Social map[string]string `json:"social,omitempty" bson:"social,omitempty"`

	Hobbies struct {
		Languages []string `json:"languages,omitempty" bson:"languages,omitempty"`

		Games struct {
			Favorite map[string]string `json:"favorite,omitempty" bson:"favorite,omitempty"`
			Consoles map[string]string `json:"consoles,omitempty" bson:"consoles,omitempty"`
			PC       bool              `json:"pc,omitempty" bson:"pc,omitempty"`
			Extra    string            `json:"extra,omitempty" bson:"extra,omitempty"`
			Accounts map[string]string `json:"accounts,omitempty" bson:"accounts,omitempty"`
		} `json:"games,omitempty" bson:"games,omitempty"`

		Sports struct {
			Plays []string `json:"plays,omitempty" bson:"plays,omitempty"`
			Likes []string `json:"likes,omitempty" bson:"likes,omitempty"`
			Teams []string `json:"teams,omitempty" bson:"teams,omitempty"`
			Extra string   `json:"extra,omitempty" bson:"extra,omitempty"`
		} `json:"sports,omitempty" bson:"sports,omitempty"`

		TV struct {
			Genres   []string          `json:"genres,omitempty" bson:"genres,omitempty"`
			Movies   []string          `json:"movies,omitempty" bson:"movies,omitempty"`
			Series   []string          `json:"series,omitempty" bson:"series,omitempty"`
			Cartoons []string          `json:"cartoons,omitempty" bson:"cartoons,omitempty"`
			Animes   []string          `json:"animes,omitempty" bson:"animes,omitempty"`
			Accounts map[string]string `json:"accounts,omitempty" bson:"accounts,omitempty"`
		} `json:"tv,omitempty" bson:"tv,omitempty"`

		Programming struct {
			Languages []string `json:"languages,omitempty" bson:"languages,omitempty"`
			OS        []string `json:"languages,omitempty" bson:"languages,omitempty"`
		} `json:"programming,omitempty" bson:"programming,omitempty"`

		Music struct {
			Genres []string `json:"genres,omitempty" bson:"genres,omitempty"`
			Bands  []string `json:"bands,omitempty" bson:"bands,omitempty"`
			Albuns []string `json:"albuns,omitempty" bson:"albuns,omitempty"`
			Songs  []string `json:"songs,omitempty" bson:"songs,omitempty"`
		} `json:"music,omitempty" bson:"music,omitempty"`
	} `json:"hobby,omitempty" bson:"hobby,omitempty"`

	Friends []struct {
		RA  string `json:"ra" bson:"ra"`
		UID string `json:"uid" bson:"uid"`
	} `json:"friends" bson:"friends"`
}
