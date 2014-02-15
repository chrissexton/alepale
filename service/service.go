package service

// Interface for interacting with services
type Service interface {
	GetChan() (MessageChan, MessageChan)
}

// Simple pipe service, stores two chans and pipes between them
type ChanService struct {
	in  MessageChan
	out MessageChan
}

// Create simple pipe service
func NewChanService(in MessageChan, out MessageChan) *ChanService {
	service := &ChanService{
		in:  in,
		out: out,
	}
	go service.manageChan()
	return service
}

// Public interface method to get I/O for the service
func (s ChanService) GetChan() (MessageChan, MessageChan) {
	return s.in, s.out
}

// Just return whatever they sent
func (s *ChanService) handleMessage(val Message) Message {
	return val
}

// goroutine for background processing of the channel
func (s *ChanService) manageChan() {
	for {
		select {
		case val := <-s.in:
			s.out <- s.handleMessage(val)
		default:
		}
	}
}
