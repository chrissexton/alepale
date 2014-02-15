package bot

import "testing"

func TestNewChannel(t *testing.T) {
	users := make([]*User, 0)
	ch := NewChannel(users)

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
	users := make([]*User, 0)
	ch := NewChannel(users)

	user := NewUser()

	ch.AddUser(user)

	if len(ch.Users) != 1 {
		t.Fail()
	}

	if ch.Users[0] != user {
		t.Fail()
	}
}
