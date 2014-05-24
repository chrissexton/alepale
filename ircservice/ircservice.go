// © 2014 the AlePale Authors under the WTFPL license. See AUTHORS for the list of authors.

package ircservice

import (
	"io"
	"log"
	"os"
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

	go ircSvc.manageChan()
	go ircSvc.start()

	return &ircSvc
}

func (s *IrcService) start() {
	client, err := irc.DialServer(s.server,
		s.nick,
		"test",
		s.password)

	if err != nil {
		log.Fatalf("Fatal: %s", err)
	}

	s.client = client

	log.Println("passing off the connection to the handlers")
	go s.handleConnection()
}

func (s *IrcService) handleConnection() {
	log.Println("#handleConnection starting")
	t := time.NewTimer(pingTime)

	defer func() {
		t.Stop()
		close(s.client.Out)
		for err := range s.client.Errors {
			if err != io.EOF {
				log.Println(err)
				s.Quit <- true // let anybody listening know we died
			}
		}
	}()

	for {
		select {
		case msg, ok := <-s.client.In:
			log.Println("#handleConnection got message:", msg, ok)
			if !ok { // disconnect
				return
			}
			t.Stop()
			t = time.NewTimer(pingTime)
			s.handleIncomming(msg)

		case <-t.C:
			log.Println("#handleConnection sending ping")
			s.client.Out <- irc.Msg{Cmd: irc.PING, Args: []string{s.client.Server}}
			t = time.NewTimer(pingTime)

		case err, ok := <-s.client.Errors:
			log.Println("#handleConnection got error:", err, ok)
			if ok && err != io.EOF {
				log.Println(err)
				return
			}
		}
	}
}

// Translate irc.Msg to svc.Message real quick
func msgToMessage(msg irc.Msg) svc.Message {

	// To complete this, we need a method of looking up users and channels by string
	return svc.Message{}
}

// Figure out what to do with incomming messages
func (s *IrcService) handleIncomming(msg irc.Msg) {
	switch msg.Cmd {
	case irc.ERROR:
		log.Println(1, "Received error: "+msg.Raw)

	case irc.PING:
		s.client.Out <- irc.Msg{Cmd: irc.PONG}

	case irc.PONG:
		// OK, ignore

	case irc.ERR_NOSUCHNICK:
		s.in <- msgToMessage(msg)

	case irc.ERR_NOSUCHCHANNEL:
		s.in <- msgToMessage(msg)

	case irc.RPL_MOTD:
		s.in <- msgToMessage(msg)

	case irc.RPL_NAMREPLY:
		s.in <- msgToMessage(msg)

	case irc.RPL_TOPIC:
		s.in <- msgToMessage(msg)

	case irc.KICK:
		s.in <- msgToMessage(msg)

	case irc.TOPIC:
		s.in <- msgToMessage(msg)

	case irc.MODE:
		s.in <- msgToMessage(msg)

	case irc.JOIN:
		s.in <- msgToMessage(msg)

	case irc.PART:
		s.in <- msgToMessage(msg)

	case irc.QUIT:
		os.Exit(1)

	case irc.NOTICE:
		s.in <- msgToMessage(msg)

	case irc.PRIVMSG:
		s.in <- msgToMessage(msg)

	case irc.NICK:
		s.in <- msgToMessage(msg)

	case irc.RPL_WHOREPLY:
		s.in <- msgToMessage(msg)

	case irc.RPL_ENDOFWHO:
		s.in <- msgToMessage(msg)

	default:
		cmd := irc.CmdNames[msg.Cmd]
		log.Println("(" + cmd + ") " + msg.Raw)
	}
}

// Translate svc.Message to irc.Msg real quick
func (s *IrcService) handleOutgoing(msg svc.Message) {
}

// goroutine for background processing of messages
func (s *IrcService) manageChan() {
	for {
		select {
		case val := <-s.in:
			s.handleOutgoing(val)
		default:
		}
	}
}
