package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/zdunecki/buf-io/config"
	"github.com/zdunecki/buf-io/integrations/slack"
	"testing"
)

func TestBufIoConfig(t *testing.T) {
	assert := assert.New(t)

	testNamespace := "{{CHANNEL_NAME}}({{CHANNEL_ID}})"
	model := &slack.InteractiveMessageValueCallback{
		ChannelName: "test-name",
		ChannelId: "test-id",
	}

	conf := &config.BufIoConfig{}

	assert.Equal(conf.NamespaceToPath(testNamespace, model), "test-name(test-id)")

	conf2 := &config.BufIoConfig{}

	testNamespace2 := "{{CHANNEL_NAME}} - {{CHANNEL_ID}}"
	model2 := &slack.InteractiveMessageValueCallback{
		ChannelName: "test-name",
		ChannelId: "test-id",
	}
	assert.Equal(conf2.NamespaceToPath(testNamespace2, model2), "test-name - test-id")

}
