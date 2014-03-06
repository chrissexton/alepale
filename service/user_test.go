// © 2014 the AlePale Authors under the WTFPL license. See AUTHORS for the list of authors.

package service

import "testing"

func TestNewUser(t *testing.T) {
	u := NewUser("flyngpngn")

	if u == nil {
		t.Fail()
	}

	if u.Name != "flyngpngn" {
		t.Fail()
	}
}
