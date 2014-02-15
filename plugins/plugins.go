package plugins

import (
	"fmt"

	"github.com/chrissexton/alepale/bot"
)

type Plugin interface {
	Process([]bot.Message) []bot.Message
}

type TestPlugin struct {
	word string
}

func NewTestPlugin(word string) *TestPlugin {
	return &TestPlugin{
		word: word,
	}
}

func (t *TestPlugin) Process(messages []bot.Message) []bot.Message {
	for i, message := range messages {
		messages[i].Text = fmt.Sprintf("%s %s", t.word, message.Text)
	}
	return messages
}
