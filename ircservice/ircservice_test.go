// © 2014 the AlePale Authors under the WTFPL license. See AUTHORS for the list of authors.

package ircservice

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/chrissexton/alepale/service"
)

func TestIrcService(t *testing.T) {
	server := "irc.freenode.net:6667"
	log.Println("test service about to be created")
	svc := NewIrcService(server, "alepaletest")

	if svc.Server != server {
		t.Fail()
	}
}

func TestConnection(t *testing.T) {
	if os.Getenv("TEST_CONN") == "" {
		t.Skipf("skipping connection test in short mode.")
	}
	server := "irc.freenode.net:6667"
	svc := NewIrcService(server, "alepaletest")

	svc.Start()
	time.Sleep(3 * time.Second)
	if !svc.client.Connected() {
		t.Fail()
	}
	svc.Disconnect("quit")
	<-svc.Quit
}

// Check to be sure that we are actually implementing a Service
func TestIrcServiceInterface(t *testing.T) {
	v := IrcService{}
	var i interface{} = v
	_, ok := i.(service.Service)
	if ok {
		t.Fail()
	}

	var p interface{} = &v
	_, ok = p.(service.Service)
	if !ok {
		t.Fail()
	}
}
