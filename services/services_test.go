package services

import (
	"os"
	"strings"
	"testing"
)

// Just test to see that it don't blow up
func TestRecieve(t *testing.T) {
	buf := strings.NewReader("testing")
	s := NewFileService(buf, os.Stdout)
	ch := s.GetChan()
	val := <-ch
	if val != "testing" {
		t.Errorf("Did not recieve good value: %s", val)
	}
	t.Logf("Got value: %s", val)
}
