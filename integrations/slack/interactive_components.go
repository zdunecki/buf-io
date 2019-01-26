package slack

type InteractiveComponent struct {
	Type    string `json:"type"`
	Actions []struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"actions"`
	CallbackId string `json:"callback_id"`
	Team       struct {
		Id     string `json:"id"`
		Domain string `json:"domain"`
	} `json:"team"`
	Channel struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"channel"`
}

var InteractiveComponentType = struct {
	InteractiveMessage InteractiveMessageType
}{
	"interactive_message",
}
