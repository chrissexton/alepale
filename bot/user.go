package bot

import "github.com/chrissexton/alepale/service"

type User struct {
	Name     string
	Aliases  []string
	Services []*service.Service
}

func NewUser() *User {
	return &User{}
}
