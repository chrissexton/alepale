// © 2014 the AlePale Authors under the WTFPL license. See AUTHORS for the list of authors.

package service

const BUFF_SZ = 10

type MessageChan chan Message

func NewMessageChan() MessageChan {
	return make(MessageChan, BUFF_SZ)
}

// Representation of a single event from a service to the bot
type Message struct {
	// Who initiated the event
	User *User

	// Which service the user was operating through
	Service *Service

	// Which channel the user was operating on
	Channel *Channel

	// The body of the event
	Text string

	// Flag indicating the process status
	Processed bool

	// Context store
	Data []interface{}
}

// Create a new message for procesing
func NewMessage(u *User, s *Service, ch *Channel, txt string, data ...interface{}) Message {
	processed := false
	if len(data) > 0 {
		switch data[0].(type) {
		case bool:
			processed = true
			data = data[1:]
			break
		}
	}
	return Message{
		User:      u,
		Service:   s,
		Channel:   ch,
		Text:      txt,
		Processed: processed,
		Data:      data,
	}
}

// Array of messages
type Log []Message

// Distributes the message to the services by which it should be sent
func (m *Message) Send() {
	srv := *m.Service
	_, out := srv.GetChan()
	out <- *m
}
