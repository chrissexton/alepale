// © 2014 the AlePale Authors under the WTFPL license. See AUTHORS for the list of authors.

package service

// Type to track a user and service connection
type ServiceAlias struct {
	// The user on the service specified
	// This could be different than the user that we are looking at the from
	User *User

	// The service this user resides on
	Service *Service
}
