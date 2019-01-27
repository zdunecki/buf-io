package slack

import (
	"encoding/json"
	"github.com/zdunecki/buf-io/config"
	"log"
)

type InteractiveMessageType string

type InteractiveMessageValueCallback struct {
	ChannelName string `json:"channel_name" bufio:"CHANNEL_NAME"`
	ChannelId   string `json:"channel_id" bufio:"CHANNEL_ID"`
	FileName    string `json:"file_name" bufio:"FILE_NAME"`
	DownloadURL string `json:"download_url"`
	Val         string `json:"val"`
}

func (t *InteractiveMessageType) callback(msg InteractiveComponent, integration config.Integration, storage []string) {
	for _, action := range msg.Actions {
		var val InteractiveMessageValueCallback
		if err := json.Unmarshal([]byte(action.Value), &val); err != nil {
			log.Fatal(err)
		}

		if val.Val != "1" {
			continue
		}
		go StorageUpload(val, integration, storage)
	}
}
