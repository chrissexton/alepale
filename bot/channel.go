package bot

type Channel struct {
	Users []*User
	Log   Log
}

func NewChannel(users []*User) *Channel {
	log := make(Log, 0)
	return &Channel{
		Users: users,
		Log:   log,
	}
}

func (ch *Channel) AddUser(user *User) {
	ch.Users = append(ch.Users, user)
}
