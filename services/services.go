package services

import (
	"bufio"
	"io"
	"log"
)

type Service interface {
	Send(message string)
	GetChan() chan string
}

type FileService struct {
	in  *bufio.Scanner
	out io.Writer
	ch  chan string
}

func NewFileService(in io.Reader, out io.Writer) *FileService {
	service := &FileService{
		in:  bufio.NewScanner(in),
		out: out,
		ch:  make(chan string),
	}
	go service.Recieve()
	return service
}

func (s *FileService) Send(message string) {
	_, err := s.out.Write([]byte(message))
	if err != nil {
		log.Println("Send failed:", err)
	}
}

func (s *FileService) GetChan() chan string {
	return s.ch
}

func (s *FileService) Recieve() {
	for s.in.Scan() {
		s.ch <- s.in.Text()
	}
	if err := s.in.Err(); err != nil {
		log.Println("reading standard input:", err)
	}
}
