package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/zdunecki/buf-io/integrations/slack"
	"github.com/zdunecki/buf-io/utils"
	"testing"
)

func TestBufIoConfig(t *testing.T) {
	assert := assert.New(t)

	testNamespace := "{{CHANNEL_NAME}}({{CHANNEL_ID}})"
	model := &slack.InteractiveMessageValueCallback{
		ChannelName: "test-name",
		ChannelId: "test-id",
	}

	assert.Equal(utils.ConfigNamespace(testNamespace, model), "test-name(test-id)")

	testNamespace2 := "{{CHANNEL_NAME}} - {{CHANNEL_ID}}"
	model2 := &slack.InteractiveMessageValueCallback{
		ChannelName: "test-name",
		ChannelId: "test-id",
	}
	assert.Equal(utils.ConfigNamespace(testNamespace2, model2), "test-name - test-id")

}
