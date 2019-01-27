package slack

import (
	"github.com/zdunecki/buf-io/config"
	"github.com/zdunecki/buf-io/integrations/dropbox"
	"io/ioutil"
	"log"
	"path"
)

func StorageUpload(val InteractiveMessageValueCallback, integration config.Integration, storage []string) {
	resp, err := RequestSlack("GET", val.DownloadURL)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	conf, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	uploadPath := conf.NamespaceToPath(integration.Namespace, &val)

	for _, s := range storage {
		go func() {
			switch s {
			case config.DropBox:
				{
					//TODO: to scale up we need upload via session
					err = dropbox.Upload(path.Join(uploadPath, val.FileName), content)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}()
	}
}
