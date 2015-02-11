package anaconda

type UrlEntity struct {
	display_url   string  `json:"display_url,omitempty"`
	Expanded_url  string  `json:"expanded_url,omitempty"`
	Indices       []int   `json:"indices,omitempty"`
	Url           string  `json:"url,omitempty"`
}

type Hashtags struct {
	Indices []int   `json:"indices,omitempty"`
	Text    string  `json:"text,omitempty"`
}

type UserMentionsEntity struct {
	Id          int64    `json:"id,omitempty"`
	Id_str      string   `json:"id_str,omitempty"`
	Indices     []int    `json:"indices,omitempty"`
	Name        string   `json:"name,omitempty"`
	Screen_name string   `json:"screen_name,omitempty"`
}

type MediaEntity struct {
	Display_url           string        `json:"display_url,omitempty"`
	Expanded_url          string        `json:"expanded_url,omitempty"`
	Id                    int64         `json:"id,omitempty"`
	Id_str                string        `json:"id_str,omitempty"`
	Indices               []int         `json:"indices,omitempty"`
	Media_url             string        `json:"media_url,omitempty"`
	Media_url_https       string        `json:"media_url_https,omitempty"`
	Sizes                 MediaSizes    `json:"sizes,omitempty"`
	Source_status_id      int64         `json:"source_status_id,omitempty"`
	Source_status_id_str  int64         `json:"source_status_id,omitempty"`
	Type                  string        `json:"type,omitempty"`
	Url                   string        `json:"url,omitempty"`
}

type MediaSizes struct {
	Thumb   MediaSize    `json:"thumb,omitempty"`
	Large   MediaSize    `json:"large,omitempty"`
	Medium  MediaSize    `json:"medium,omitempty"`
	Small   MediaSize    `json:"small,omitempty"`
}

type MediaSize struct {
	H        int      `json:"h,omitempty"`
	Resize   string   `json:"resize,omitempty"`
	W        int      `json:"w,omitempty"`
}

type Entities struct {
	Hashtags      []Hashtags             `json:"hashtags,omitempty"`
	Media         []MediaEntity          `json:"media,omitempty"`
	Urls          []UrlEntity            `json:"urls,omitempty"`
	User_mentions []UserMentionsEntity   `json:"user_mentions,omitempty"`
}
