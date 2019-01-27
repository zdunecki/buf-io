package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"
)

const DropBox = "dropbox"

type Integration struct {
	Namespace string   `json:"namespace"`
	NoAck     []string `json:"noack"`
}

type BufIoConfig struct {
	Config struct {
		Storage      []string `json:"storage"`
		Integrations struct {
			Slack Integration `json:"slack"`
		} `json:"integrations"`
	} `json:"config"`
}

func (c *BufIoConfig) NamespaceToPath(namespace string, model interface{}) string {
	var oldnew []string
	val := reflect.ValueOf(model).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		v := valueField.Interface()

		typeField := val.Type().Field(i)
		tag := typeField.Tag
		bufioField := strings.Split(tag.Get("bufio"), ",")[0]
		oldnew = append(oldnew, "{{"+bufioField+"}}", v.(string))
	}

	replacer := strings.NewReplacer(oldnew...)
	return replacer.Replace(namespace)
}

func Get() (*BufIoConfig, error) {
	data, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "src/github.com/zdunecki/buf-io", "buf-io.yaml"))
	if err != nil {
		return nil, err
	}

	var c BufIoConfig
	err = yaml.Unmarshal([]byte(data), &c)

	return &c, nil
}
