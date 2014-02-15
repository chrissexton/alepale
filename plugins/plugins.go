package plugins

import (
	"fmt"

	"github.com/chrissexton/alepale/service"
)

type Plugin interface {
	Process([]service.Message) []service.Message
}

// Example plugin that adds a string to input
type TestPlugin struct {
	word string
}

// Create a new test plugin with a given string to add
func NewTestPlugin(word string) *TestPlugin {
	return &TestPlugin{
		word: word,
	}
}

// Translate the message into its final form
func (t *TestPlugin) Process(messages []service.Message) []service.Message {
	for i, message := range messages {
		messages[i].Text = fmt.Sprintf("%s %s", t.word, message.Text)
	}
	return messages
}
