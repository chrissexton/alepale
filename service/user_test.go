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
