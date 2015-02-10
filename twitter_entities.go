package anaconda

type UrlEntity struct {
	display_url		string	`json:"display_url"`
	Expanded_url	string	`json:"expanded_url"`
	Indices				[]int		`json:"indices"`
	Url						string	`json:"url"`
}

type Hashtags struct {
	Indices	[]int		`json:"indices"`
	Text		string	`json:"text"`
}

type UserMentionsEntity struct {
	Id					int64		`json:"id"`
	Id_str			string	`json:"id_str"`
	Indices			[]int		`json:"indices"`
	Name				string	`json:"name"`
	Screen_name string	`json:"screen_name"`
}

type MediaEntity struct {
	Display_url						string				`json:"display_url"`
	Expanded_url					string				`json:"expanded_url"`
	Id										int64					`json:"id"`
	Id_str								string				`json:"id_str"`
	Indices								[]int					`json:"indices"`
	Media_url							string				`json:"media_url"`
	Media_url_https				string				`json:"media_url_https"`
	Sizes									MediaSizes		`json:"sizes"`
	Source_status_id			int64					`json:"source_status_id"`
	Source_status_id_str	int64					`json:"source_status_id"`
	Type									string				`json:"type"`
	Url										string				`json:"url"`
}

type MediaSizes struct {
	Thumb	MediaSize		`json:"thumb"`
	Large	MediaSize		`json:"large"`
	Medium MediaSize	`json:"medium"`
	Small	MediaSize		`json:"small"`
}

type MediaSize struct {
	H				int			`json:"h"`
	Resize	string	`json:"resize"`
	W				int			`json:"w"`
}

type Entities struct {
	Hashtags			[]Hashtags						`json:"hashtags"`
	Media					[]MediaEntity					`json:"media"`
	Urls					[]UrlEntity						`json:"urls"`
	User_mentions	[]UserMentionsEntity	`json:"user_mentions"`
}
