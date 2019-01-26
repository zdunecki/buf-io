package utils

import (
	"log"
	"github.com/zdunecki/buf-io/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"
)

func GetBufIoConfig() config.BufIoConfig {
	data, err := ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "src/github.com/zdunecki/buf-io", "buf-io.yaml"))
	if err != nil {
		log.Fatal(err)
	}

	var c config.BufIoConfig
	err = yaml.Unmarshal([]byte(data), &c)

	return c
}

func ConfigNamespace(namespace string, model interface{}) string {
	var oldnew []string
	val := reflect.ValueOf(model).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		v := valueField.Interface()

		typeField := val.Type().Field(i)
		tag := typeField.Tag
		bufioField := strings.Split(tag.Get("bufio"), ",")[0]
		oldnew = append(oldnew, "{{" + bufioField + "}}", v.(string))
	}

	replacer := strings.NewReplacer(oldnew...)
	return replacer.Replace(namespace)
}