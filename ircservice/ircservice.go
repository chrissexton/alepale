// © 2014 the AlePale Authors under the WTFPL license. See AUTHORS for the list of authors.

package ircservice

import (
	"fmt"
	"log"
	"time"

	svc "github.com/chrissexton/alepale/service"
	irc "github.com/fluffle/goirc/client"

	"github.com/chrissexton/alepale/service"
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

	nick     string
	channels []string

	// Signal anybody else that we've exited
	Quit chan bool

	client *irc.Conn

	*irc.Config
}

func NewIrcService(server string, nick string) *IrcService {
	in, out := svc.NewMessageChan(), svc.NewMessageChan()

	// Apparently you can't specify keys properly when using anonymous fields!?
	ircSvc := IrcService{
		in:       in,
		out:      out,
		nick:     nick,
		channels: make([]string, 0),
		Quit:     make(chan bool),
		Config:   irc.NewConfig(nick),
	}

	ircSvc.Server = server

	return &ircSvc
}

// Start the IRC connection
func (i *IrcService) Start() {
	i.client = irc.SimpleClient(i.nick)

	i.client.HandleFunc(irc.DISCONNECTED, func(conn *irc.Conn, line *irc.Line) {
		log.Println("IrcService#Start disconnecting")
		i.Quit <- true
	})

	i.client.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {
		i.out <- i.convertLine(line)
	})

	log.Println("IrcService#Start connecting")
	if err := i.client.ConnectTo(i.Server); err != nil {
		fmt.Printf("Connection error: %s\n", err)
	}
}

func (i *IrcService) Disconnect(message string) {
	i.client.Quit(message)
}

// Public interface method to get I/O for the service
func (i *IrcService) GetChan() (service.MessageChan, service.MessageChan) {
	return i.in, i.out
}

func (i *IrcService) convertLine(line *irc.Line) service.Message {
	ch := service.NewChannel(line.Target())
	u := ch.Users[line.Nick]
	return service.NewMessage(u, i, ch, line.Text(), line)
}
