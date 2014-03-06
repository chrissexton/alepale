// © 2014 the AlePale Authors under the WTFPL license. See AUTHORS for the list of authors.

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
