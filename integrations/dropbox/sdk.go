package dropbox

import (
	"bytes"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"os"
)

func Upload(fileName string, content []byte) error {
	config := dropbox.Config{
		Token: os.Getenv("DROPBOX_TOKEN"),
	}
	dbx := files.New(config)

	info := &files.CommitInfo{
		Path: "/" + fileName,
		Mode: &files.WriteMode{
			Tagged: dropbox.Tagged{
				Tag: "overwrite",
			},
		},
	}

	_, err := dbx.Upload(info, bytes.NewReader(content))

	return err
}
