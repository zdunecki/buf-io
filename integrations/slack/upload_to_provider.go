package slack

import (
	"github.com/zdunecki/buf-io/integrations/dropbox"
	"github.com/zdunecki/buf-io/utils"
	"io/ioutil"
	"log"
	"path"
)

func UploadToProvider(val InteractiveMessageValueCallback) {
	resp, err := RequestSlack("GET", val.DownloadURL)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	namespace := utils.GetBufIoConfig().Config.Integrations.Slack.Namespace
	uploadPath := utils.ConfigNamespace(namespace, &val)

	//TODO: to scale up we need upload via session
	err = dropbox.Upload(path.Join(uploadPath, val.FileName), content)

	if err != nil {
		log.Fatal(err)
	}
}