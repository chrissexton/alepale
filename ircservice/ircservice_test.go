package ircservice

import (
	"log"
	"testing"
)

func TestIrcService(t *testing.T) {
	server := "irc.freenode.net:6667"
	svc := NewIrcService(server)

	log.Println("test")

	if svc.server != server {
		t.Fail()
	}
}
