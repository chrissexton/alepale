// © 2014 the AlePale Authors under the WTFPL license. See AUTHORS for the list of authors.

package service

import "testing"

func makeMessage(s Service, text string) Message {
	u := NewUser("biff")
	ch := NewChannel("woo")
	return NewMessage(u, &s, ch, text)
}

// Just test to see that it don't blow up
func TestChanService(t *testing.T) {
	in, out := make(MessageChan, 10), make(MessageChan, 10)
	t.Log("Making service")
	s := *NewChanService(in, out)
	t.Log("Sending a message to chan")
	in <- makeMessage(s, "testing")
	t.Log("Waiting for message back")
	val := <-out
	t.Log("Got message back")
	if val.Text != "testing" {
		t.Errorf("Did not recieve good value: %s", val.Text)
	}
	t.Log("Got value:", val.Text)
}

func TestGetChan(t *testing.T) {
	in, out := make(MessageChan, 10), make(MessageChan, 10)
	s := NewChanService(in, out)
	actualIn, actualOut := s.GetChan()

	if in != actualIn {
		t.Fail()
	}

	if out != actualOut {
		t.Fail()
	}

	actualIn <- makeMessage(s, "testing")
	val := <-actualOut
	if val.Text != "testing" {
		t.Fail()
	}
}
