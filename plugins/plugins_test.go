package plugins

import (
	"fmt"
	"testing"

	"github.com/chrissexton/alepale/bot"
)

func TestTestPlugin(t *testing.T) {
	word := "test"
	expected := fmt.Sprintf("%s %s", word, word)
	tp := NewTestPlugin(word)
	msg := bot.NewMessage(nil, nil, nil, word)
	actual := tp.Process([]bot.Message{msg})[0].Text

	if actual != "test test" {
		t.Errorf("Expected %s, got %s.", expected, actual)
	}
}
