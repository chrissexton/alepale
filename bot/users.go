package bot

import (
	// "labix.org/v2/mgo"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

// User type stores user history. This is a vehicle that will follow the user for the active
// session
type User struct {
	// Current nickname known
	Name string

	// LastSeen	DateTime

	// Alternative nicknames seen
	Alts   []string
	Parent string

	Admin bool

	//bot *Bot
}

func (b *Bot) NewUser(nick string) *User {
	return &User{
		Name:  nick,
		Admin: b.checkAdmin(nick),
	}
}

func (b *Bot) GetUser(nick string) *User {
	coll := b.Db.C("users")
	query := coll.Find(bson.M{"name": nick})
	var user *User

	count, err := query.Count()
	log.Printf("Searching for %s. Found %d entries.\n", nick, count)

	if err != nil {
		log.Printf("Error fetching user, %s: %s\n", nick, err)
		user = b.NewUser(nick)
		coll.Insert(*user)
	} else if count == 1 {
		err = query.One(&user)
		if err != nil {
			panic(err)
		}
		log.Printf("Found user: %s\n", nick)
	} else if count == 0 {
		// create the user
		log.Printf("Creating new user: %s\n", nick)
		user = b.NewUser(nick)
		coll.Insert(user)
	} else {
		log.Printf("Error: %s appears to have more than one user?\n", nick)
		query.One(user)
	}

	// grab linked user, if any
	if user.Parent != "" {
		query := coll.Find(bson.M{"Name": user.Parent})
		if count, err := query.Count(); err != nil && count == 1 {
			query.One(user)
		} else {
			log.Printf("Error: bad linkage on %s -> %s.\n",
				user.Name,
				user.Parent)
		}
	}

	found := false
	for _, u := range b.Users {
		if u.Name == user.Name {
			found = true
		}
	}
	if !found {
		b.Users = append(b.Users, *user)
	}

	return user
}

// Modify user entry to be a link to other, return other
func (u *User) LinkUser(coll *mgo.Collection, other *User) *User {
	user := u

	other.Alts = append(other.Alts, user.Alts...)
	user.Alts = []string{}
	user.Parent = other.Name

	err := coll.Update(bson.M{"name": u.Name}, u)
	if err != nil {
		log.Printf("Error updating user: %s\n", u.Name)
	}

	err = coll.Update(bson.M{"name": other.Name}, other)
	if err != nil {
		log.Printf("Error updating other user: %s\n", other.Name)
	}

	return other
}

func (b *Bot) checkAdmin(nick string) bool {
	for _, u := range b.Config.Admins {
		if nick == u {
			return true
		}
	}
	return false
}
