package plugins

import (
	"fmt"
	"testing"

	"github.com/chrissexton/alepale/service"
)

func TestTestPlugin(t *testing.T) {
	word := "test"
	expected := fmt.Sprintf("%s %s", word, word)
	tp := NewTestPlugin(word)
	msg := service.NewMessage(nil, nil, nil, word)
	actual := tp.Process([]service.Message{msg})[0].Text

	if actual != "test test" {
		t.Errorf("Expected %s, got %s.", expected, actual)
	}
}
