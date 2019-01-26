package slack

type Event struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	TeamId   string `json:"team_id"`
	ApiAppId string `json:"api_app_id"`
	Event    struct {
		Type    string `json:"type"`
		UserId  string `json:"user_id"`
		EventId string `json:"event_id"`
	} `json:"event"`
	Challenge string `json:"challenge"`
}

var EventType = struct {
	FileCreated FileCreatedType
}{
	"file_created",
}
