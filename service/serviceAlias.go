package service

// Type to track a user and service connection
type ServiceAlias struct {
	// The user on the service specified
	// This could be different than the user that we are looking at the from
	User *User

	// The service this user resides on
	Service *Service
}
