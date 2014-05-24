// © 2014 the AlePale Authors under the WTFPL license. See AUTHORS for the list of authors.

package service

import "testing"

// func NewMessage(u *User, s *service.Service, ch *Channel, txt string, data ...interface{}) Message {

func TestNewMessage(t *testing.T) {
	m := NewMessage(nil, nil, nil, "test")
	if m.Processed {
		t.Error("Message should not be processed if not specified")
	}
}

func TestProcessedMessage(t *testing.T) {
	m := NewMessage(nil, nil, nil, "test", false)
	if !m.Processed {
		t.Error("Message should be processed")
	}
	if len(m.Data) > 0 {
		t.Error("Bool should have been taken out of data")
	}
	m = NewMessage(nil, nil, nil, "test", false, "a", "b")
	if len(m.Data) != 2 || m.Data[0] != "a" || m.Data[1] != "b" {
		t.Error("Incorrect data segment on message")
	}
	m = NewMessage(nil, nil, nil, "test", "a", "b")
	if len(m.Data) != 2 || m.Data[0] != "a" || m.Data[1] != "b" {
		t.Error("Incorrect data segment on message:", m.Data)
	}
}

func TestNewMessageChan(t *testing.T) {
	c := NewMessageChan()
	if BUFF_SZ != cap(c) {
		t.Errorf("Channel returned is of incorrect capacity: %d", cap(c))
	}
}

func TestSend(t *testing.T) {
	in, out := make(MessageChan, 10), make(MessageChan, 10)
	s := NewChanService(in, out)
	m := NewMessage(nil, s, nil, "test")
	m.Send()
}
