// © 2014 the AlePale Authors under the WTFPL license. See AUTHORS for the list of authors.

package service

import "testing"

func TestNewChannelUsrs(t *testing.T) {
	users := make([]*User, 0)
	ch := NewChannelUsers("woo", users)

	if ch.Name != "woo" {
		t.Fail()
	}

	if len(ch.Users) != len(users) {
		t.Fail()
	}

	for i, user := range ch.Users {
		if user != users[i] {
			t.Fail()
		}
	}

	if len(ch.Log) != 0 {
		t.Fail()
	}
}

func TestNewChannel(t *testing.T) {
	users := make([]*User, 0)
	ch := NewChannel("woo")

	if ch.Name != "woo" {
		t.Fail()
	}

	if len(ch.Users) != len(users) {
		t.Fail()
	}

	for i, user := range ch.Users {
		if user != users[i] {
			t.Fail()
		}
	}

	if len(ch.Log) != 0 {
		t.Fail()
	}
}

func TestAddUser(t *testing.T) {
	ch := NewChannel("woo")

	user := NewUser("biff")

	ch.AddUser(user)

	if len(ch.Users) != 1 {
		t.Fail()
	}

	if ch.Users[0] != user {
		t.Fail()
	}
}
