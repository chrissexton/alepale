package bot

import "testing"

func TestNewUser(t *testing.T) {
	u := NewUser()
	u.Name = "flyngngn"

	if u == nil {
		t.Fail()
	}
}
