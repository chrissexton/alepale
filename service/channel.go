package service

// A channel defines a set of users and history for a chat group
type Channel struct {
	Name  string
	Users []*User
	Log   Log
}

// Make a new channel
//
// This could be a user on a service (like IRC), or a venue for multiple users
func NewChannelUsers(name string, users []*User) *Channel {
	log := make(Log, 0)
	return &Channel{
		Name:  name,
		Users: users,
		Log:   log,
	}
}

// Return a new named channel
func NewChannel(name string) *Channel {
	return NewChannelUsers(name, make([]*User, 0))
}

// Add a new User to a channel.
// Generally done on login and never removed
// Set the users' offline bit (and online bit) after they exist
func (ch *Channel) AddUser(user *User) {
	ch.Users = append(ch.Users, user)
}
