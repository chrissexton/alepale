// © 2014 the AlePale Authors under the WTFPL license. See AUTHORS for the list of authors.

package service

var channels map[string]*Channel

// A channel defines a set of users and history for a chat group
type Channel struct {
	Name  string
	Users map[string]*User
	Log   Log
}

// Make a new channel
//
// This could be a user on a service (like IRC), or a venue for multiple users
func NewChannelUsers(name string, users map[string]*User) *Channel {
	if _, ok := channels[name]; ok {
		return channels[name]
	}
	log := make(Log, 0)
	ch := &Channel{
		Name:  name,
		Users: users,
		Log:   log,
	}
	channels[name] = ch
	return ch
}

// Return a new named channel
func NewChannel(name string) *Channel {
	return NewChannelUsers(name, make(map[string]*User))
}

// Add a new User to a channel.
// Generally done on login and never removed
// Set the users' offline bit (and online bit) after they exist
func (ch *Channel) AddUser(user *User) {
	ch.Users[user.Name] = user
}

func init() {
	channels = make(map[string]*Channel)
}
