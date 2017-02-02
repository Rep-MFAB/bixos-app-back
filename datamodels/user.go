package datamodels

import "time"

// TODO incorporate this into a dedicated util.
type mongodbUser struct {
	Name  string
	Photo string
	RA    string

	Account struct {
		Email    string
		Password []byte
		Salt     []byte
		UID      string
	}

	Personal struct {
		Birth        time.Time
		Phone        string
		Drinks       int
		Relationship struct {
			Sex    string
			Gender string
			Status string
		}
		Origin struct {
			Country string
			State   string
			City    string
			School  string
		}
	}

	University struct {
		Grad      string
		Languages string
		Reason    struct {
			Unicamp string
			Grad    string
		}
	}

	Work struct {
		Job     string
		Company string
	}

	Social map[string]string

	Hobbies struct {
		Games struct {
			Favorite map[string]string
			Consoles map[string]string
			PC       bool
			Extra    string
			Accounts map[string]string
		}

		Sports struct {
			Plays []string
			Likes []string
			Teams []string
			Extra string
		}

		TV struct {
			Genres   []string
			Movies   []string
			Series   []string
			Cartoons []string
			Disney   []string
			Animes   []string
			Accounts map[string]string
		}

		Programming struct {
			Languages []string
			OS        []string
		}

		Music struct {
			Genres []string
			Bands  []string
			Albuns []string
			Songs  []string
		}
	}

	Friends []struct {
		RA  string
		UID string
	}
}
