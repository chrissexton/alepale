package service

type Service interface {
	GetChan() (chan string, chan string)
}

type ChanService struct {
	in  chan string
	out chan string
}

func NewChanService(in chan string, out chan string) *ChanService {
	service := &ChanService{
		in:  in,
		out: out,
	}
	go service.manageChan()
	return service
}

func (s *ChanService) GetChan() (chan string, chan string) {
	return s.in, s.out
}

// Just return whatever they sent
func (s *ChanService) handleMessage(val string) string {
	return val
}

func (s *ChanService) manageChan() {
	for {
		select {
		case val := <-s.in:
			s.out <- s.handleMessage(val)
		default:
		}
	}
}
