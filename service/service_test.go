package service

import (
	"log"
	"testing"
)

// Just test to see that it don't blow up
func TestChanService(t *testing.T) {
	in, out := make(chan string, 10), make(chan string, 10)
	log.Println("Making service")
	NewChanService(in, out)
	log.Println("Sending a message to chan")
	in <- "testing"
	log.Println("Waiting for message back")
	val := <-out
	log.Println("Got message back")
	if val != "testing" {
		t.Errorf("Did not recieve good value: %s", val)
	}
	log.Println("Got value:", val)
}

func TestGetChan(t *testing.T) {
	in, out := make(chan string, 10), make(chan string, 10)
	s := NewChanService(in, out)
	actualIn, actualOut := s.GetChan()

	if in != actualIn {
		t.Fail()
	}

	if out != actualOut {
		t.Fail()
	}

	actualIn <- "testing"
	val := <-actualOut
	if val != "testing" {
		t.Fail()
	}
}
