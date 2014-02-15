package bot

import (
	"log"

	"github.com/chrissexton/alepale/service"
)

// sketch of setting up a new service
func setupService() {
	in, out := make(chan string, 10), make(chan string, 10)
	log.Println("Making service")
	service.NewChanService(in, out)
	log.Println("Sending a message to chan")
	in <- "testing"
	log.Println("Waiting for message back")
	val := <-out
	log.Println("Got message back")
	log.Println("Got value:", val)
}
