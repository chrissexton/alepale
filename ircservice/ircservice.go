// © 2014 the AlePale Authors under the WTFPL license. See AUTHORS for the list of authors.

package ircservice

import (
	"time"

	"code.google.com/p/velour/irc"
	svc "github.com/chrissexton/alepale/service"
)

const (
	// DefaultPort is the port used to connect to
	// the server if one is not specified.
	defaultPort = "6667"

	// InitialTimeout is the initial amount of time
	// to delay before reconnecting.  Each failed
	// reconnection doubles the timout until
	// a connection is made successfully.
	initialTimeout = 2 * time.Second

	// PingTime is the amount of inactive time
	// to wait before sending a ping to the server.
	pingTime = 120 * time.Second
)

// Implements alepale.Service
type IrcService struct {
	in  svc.MessageChan
	out svc.MessageChan

	server   string
	nick     string
	fullName string
	password string
	channels []string

	client *irc.Client

	Quit chan bool
}

func NewIrcService(server string, nick string) *IrcService {
	in, out := svc.NewMessageChan(), svc.NewMessageChan()

	// Apparently you can't specify keys properly when using anonymous fields!?
	ircSvc := IrcService{
		in:       in,
		out:      out,
		channels: make([]string, 0),
		Quit:     make(chan bool),
	}
	ircSvc.nick = nick
	ircSvc.server = server
	ircSvc.channels = make([]string, 0)

	return &ircSvc
}
