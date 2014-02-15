package ircservice

import "testing"

func TestIrcService(t *testing.T) {
	server := "irc.freenode.net"
	svc := NewIrcService(server)

	if svc.Server != server {
		t.Fail()
	}
}
