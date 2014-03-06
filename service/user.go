// © 2014 the AlePale Authors under the WTFPL license. See AUTHORS for the list of authors.

package service

type User struct {
	// The primary identifier for this user
	Name string

	// Keybag for any extra info
	//
	// This could include:
	// - hostname
	// - server connected
	// - Admin status
	ExtraInfo map[string]interface{}

	// Any alternate names we've ever seen for the user
	Aliases []string

	// Link the user to multiple services
	Services []ServiceAlias

	// Track the users' current state
	Online bool
}

// Set up a new user object
func NewUser(name string) *User {
	return &User{
		Name:      name,
		ExtraInfo: make(map[string]interface{}),
		Aliases:   make([]string, 0),
		Services:  make([]ServiceAlias, 0),
	}
}
