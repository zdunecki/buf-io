package slack

import (
	"encoding/json"
	"github.com/nlopes/slack"
	"github.com/satori/go.uuid"
	"github.com/zdunecki/buf-io/config"
	"github.com/zdunecki/buf-io/utils"
)

type FileCreatedType string

func (t *FileCreatedType) createAttachments(valueYes, valueNo InteractiveMessageValueCallback) slack.MsgOption {
	valueYesJSON, err := json.Marshal(valueYes)
	if err != nil {
		panic(err)
	}

	valueNoJSON, err := json.Marshal(valueNo)
	if err != nil {
		panic(err)
	}

	return slack.MsgOptionAttachments(slack.Attachment{
		Text:       "Wanna upload?",
		Fallback:   "You are unable to upload a file",
		CallbackID: uuid.NewV1().String(),
		Color:      "#3AA3E3",
		Actions: []slack.AttachmentAction{
			{
				Name:  "yes",
				Text:  "Yes",
				Type:  "button",
				Value: string(valueYesJSON),
			},
			{
				Name:  "no",
				Text:  "No",
				Type:  "button",
				Value: string(valueNoJSON),
			},
		},
	})
}

func (t *FileCreatedType) callback(fileEvent []byte, event *Event) error {
	file := &File{}
	err := json.Unmarshal(fileEvent, &file)

	if err != nil {
		return err
	}

	fileInfo, err := GetFileInfo(file.Event.FileId)
	if err != nil {
		return err
	}

	var sharesContent map[ChannelId][]SharesContent

	if fileInfo.File.IsPublic {
		sharesContent = fileInfo.File.Shares.Public
	} else {
		sharesContent = fileInfo.File.Shares.Private
	}

	//TODO: what if we share file between channels?
	for channelId := range sharesContent {
		channelName, err := GetChannelInfo(string(channelId))
		if err != nil {
			return err
		}

		valueYes := InteractiveMessageValueCallback{
			ChannelName: channelName.Name,
			ChannelId:   string(channelId),
			FileName:    fileInfo.File.Name,
			DownloadURL: fileInfo.File.UrlPrivateDownload,
			Val:         "1",
		}

		conf, err := config.Get()
		if err != nil {
			return err
		}

		noAckRoom := utils.Contains(conf.Config.Integrations.Slack.NoAck, string(channelId))

		if noAckRoom {
			go StorageUpload(valueYes, conf.Config.Integrations.Slack, conf.Config.Storage)
			continue
		}


		valueNo := valueYes
		valueNo.Val = "0"

		attachments := t.createAttachments(valueYes, valueNo)

		if err := PostEphemeral(string(channelId), event.Event.UserId, attachments); err != nil {
			return err
		}
	}

	return nil
}
