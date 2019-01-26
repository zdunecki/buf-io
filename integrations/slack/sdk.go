package slack

import (
	"bytes"
	"encoding/json"
	"github.com/nlopes/slack"
	"io/ioutil"
	"net/http"
	"os"
)

func GetFileInfo(fileId string) (*FileInfo, error) {
	fileInfo := &FileInfo{}

	resp, err := RequestSlack("GET", "https://slack.com/api/files.info?file="+fileId)

	if err != nil {
		return nil, err
	}

	fileInfoBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(fileInfoBody, &fileInfo); err != nil {
		return nil, err
	}

	return fileInfo, nil
}

func GetChannelInfo(channelId string) (*slack.Channel, error) {
	token := os.Getenv("SLACK_TOKEN")

	api := slack.New(token)

	return api.GetChannelInfo(channelId)
}

func Upload(fileName string, content []byte) {
	token := os.Getenv("SLACK_TOKEN")

	api := slack.New(token)

	newFile := slack.FileUploadParameters{
		Filename: fileName,
		Reader:   bytes.NewReader(content),
	}

	if _, err := api.UploadFile(newFile); err != nil {
		panic(err)
	}
}

func PostEphemeral(channelID, userID string, options... slack.MsgOption) error {
	token := os.Getenv("SLACK_TOKEN")

	api := slack.New(token)

	_, err := api.PostEphemeral(
		channelID,
		userID,
		options...,
	)

	return err
}

func RequestSlack(method, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("SLACK_TOKEN"))
	client := &http.Client{}

	return client.Do(req)
}
