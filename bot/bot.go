package bot

import (
	"log"

	s "github.com/chrissexton/alepale/service"
)

// sketch of setting up a new service
func setupService() {
	in, out := make(s.MessageChan, 10), make(s.MessageChan, 10)
	log.Println("Making service")
	s.NewChanService(in, out)
}
