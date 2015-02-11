package anaconda

import (
	"fmt"
	"time"
)

type Tweet struct {
	Contributors         []Contributor     `json:"contributors,omitempty"` // Not yet generally available to all, so hard to test
	Coordinates          *Coordinates      `json:"coordinates,omitempty"`
	CreatedAt            string            `json:"created_at,omitempty"`
	Entities             Entities          `json:"entities,omitempty"`
	FavoriteCount        int               `json:"favorite_count,omitempty"`
	Favorited            bool              `json:"favorited,omitempty"`
	FilterLevel          string            `json:"filter_level,omitempty"`
	Id                   int64             `json:"id,omitempty"`
	IdStr                string            `json:"id_str,omitempty"`
	InReplyToScreenName  string            `json:"in_reply_to_screen_name,omitempty"`
	InReplyToStatusID    int64             `json:"in_reply_to_status_id,omitempty"`
	InReplyToStatusIdStr string            `json:"in_reply_to_status_id_str,omitempty"`
	InReplyToUserID      int64             `json:"in_reply_to_user_id,omitempty"`
	InReplyToUserIdStr   string            `json:"in_reply_to_user_id_str,omitempty"`
	Lang                 string            `json:"lang,omitempty"`
	Place                Place             `json:"place,omitempty"`
	PossiblySensitive    bool              `json:"possibly_sensitive,omitempty"`
	RetweetCount         int               `json:"retweet_count,omitempty"`
	Retweeted            bool              `json:"retweeted,omitempty"`
	RetweetedStatus      *Tweet            `json:"retweeted_status,omitempty"`
	Source               string            `json:"source,omitempty"`
	Scopes               map[string]string `json:"scopes,omitempty"`
	Text                 string            `json:"text,omitempty"`
	Truncated            bool              `json:"truncated,omitempty"`
	User                 User              `json:"user,omitempty"`
	WithheldCopyright    bool              `json:"withheld_copyright,omitempty"`
	WithheldInCountries  []string          `json:"withheld_in_countries,omitempty"`
	WithheldScope        string            `json:"withheld_scope,omitempty"`

	//Geo is deprecated
	//Geo                  interface{} `json:"geo"`
}

// CreatedAtTime is a convenience wrapper that returns the Created_at time, parsed as a time.Time struct
func (t Tweet) CreatedAtTime() (time.Time, error) {
	return time.Parse(time.RubyDate, t.CreatedAt)
}

// It may be worth placing these in an additional source file(s)

// Could also use User, since the fields match, but only these fields are possible in Contributor
type Contributor struct {
	Id         int64  `json:"id"`
	IdStr      string `json:"id_str"`
	ScreenName string `json:"screen_name"`
}

type Coordinates struct {
	Coordinates [2]float64 `json:"coordinates"` // Coordinate always has to have exactly 2 values
	Type        string     `json:"type"`
}

// HasCoordinates is a helper function to easily determine if a Tweet has coordinates associated with it
func (t Tweet) HasCoordinates() bool {
	if t.Coordinates != nil {
		if t.Coordinates.Type == "Point" {
			return true
		}
	}
	return false
}

// The following provide convenience and eliviate confusion about the order of coordinates in the Tweet

// Latitude is a convenience wrapper that returns the latitude easily
func (t Tweet) Latitude() (float64, error) {
	if t.HasCoordinates() {
		return t.Coordinates.Coordinates[1], nil
	}
	return 0, fmt.Errorf("No Coordinates in this Tweet")
}

// Longitude is a convenience wrapper that returns the longitude easily
func (t Tweet) Longitude() (float64, error) {
	if t.HasCoordinates() {
		return t.Coordinates.Coordinates[0], nil
	}
	return 0, fmt.Errorf("No Coordinates in this Tweet")
}

// X is a concenience wrapper which returns the X (Longitude) coordinate easily
func (t Tweet) X() (float64, error) {
	return t.Longitude()
}

// Y is a convenience wrapper which return the Y (Lattitude) corrdinate easily
func (t Tweet) Y() (float64, error) {
	return t.Latitude()
}
