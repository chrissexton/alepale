// © 2014 the AlePale Authors under the WTFPL license. See AUTHORS for the list of authors.

package ircservice

import (
	"log"
	"testing"
)

// Online tests
func TestIrcService(t *testing.T) {
	server := "irc.freenode.net:6667"
	log.Println("test service about to be created")
	svc := NewIrcService(server, "alepaletest")

	if svc.server != server {
		t.Fail()
	}
}
