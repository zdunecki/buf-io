package slack

type File struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	TeamId   string `json:"team_id"`
	ApiAppId string `json:"api_app_id"`
	Event    struct {
		Type string `json:"type"`
		File struct {
			Id string `json:"id"`
		} `json:"file"`
		FileId  string `json:"file_id"`
		UserId  string `json:"user_id"`
		EventId string `json:"event_id"`
	} `json:"event"`
	Challenge string `json:"challenge"`
}

type ChannelId string

type SharesContent struct {
	Ts          string `json:"ts"`
	ChannelName string `json:"channel_name"`
	TeamId      string `json:"team_id"`
}

type FileInfo struct {
	File struct {
		Name               string `json:"name"`
		IsPublic           bool   `json:"is_public"`
		UrlPrivateDownload string `json:"url_private_download"`
		Shares             struct {
			Public  map[ChannelId][]SharesContent `json:"public"`
			Private map[ChannelId][]SharesContent `json:"private"`
		} `json:"shares"`
		Channels []string `json:"channels"`
	} `json:"file"`
}
