package plugins

import (
	"bitbucket.org/phlyingpenguin/godeepintir/bot"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"math/rand"
	"strings"
	"time"
)

// This is a skeleton plugin to serve as an example and quick copy/paste for new plugins.

type userRemember struct {
	Nick    string
	Message string
	Date    time.Time
}

type RememberPlugin struct {
	Bot  *bot.Bot
	Coll *mgo.Collection
	Log  map[string][]bot.Message
}

// NewRememberPlugin creates a new RememberPlugin with the Plugin interface
func NewRememberPlugin(b *bot.Bot) *RememberPlugin {
	p := RememberPlugin{
		Bot: b,
		Log: make(map[string][]bot.Message),
	}
	p.LoadData()
	for _, channel := range b.Config.Channels {
		go p.quoteTimer(channel)
	}
	return &p
}

// Message responds to the bot hook on recieving messages.
// This function returns true if the plugin responds in a meaningful way to the users message.
// Otherwise, the function returns false and the bot continues execution of other plugins.
func (p *RememberPlugin) Message(message bot.Message) bool {
	// This bot does not reply to anything

	if message.Body == "quote" && message.Command {
		q := p.randQuote()
		p.Bot.SendMessage(message.Channel, q)

		// is it evil not to remember that the user said quote?
		return true
	}

	parts := strings.Fields(message.Body)
	if message.Command && len(parts) >= 3 && parts[0] == "remember" {
		// we have a remember!
		// look through the logs and find parts[1] as a user, if not, fuck this hoser
		nick := parts[1]
		snip := strings.Join(parts[2:], " ")

		if nick == message.User.Name {
			msg := fmt.Sprintf("Don't try to quote yourself, %s.", nick)
			p.Bot.SendMessage(message.Channel, msg)
			return true
		}

		for i := len(p.Log[message.Channel])-1; i >= 0; i-- {
			entry := p.Log[message.Channel][i]
			// find the entry we want
			fmt.Printf("Comparing '%s' to '%s'\n", entry.Raw, snip)
			if entry.User.Name == nick && strings.Contains(entry.Body, snip) {
				// insert new remember entry
				var msg string

				// check if it's an action
				if entry.Action {
					msg = fmt.Sprintf("*%s* %s", entry.User.Name, entry.Raw)
				} else {
					msg = fmt.Sprintf("<%s> %s", entry.User.Name, entry.Raw)
				}
				u := userRemember{
					Nick:    entry.User.Name,
					Message: msg,
					Date:    time.Now(),
				}
				p.Coll.Insert(u)

				// sorry, not creative with names so we're reusing msg
				msg = fmt.Sprintf("Okay, %s, remembering '%s'.",
					message.User.Name, msg)
				p.Bot.SendMessage(message.Channel, msg)
				p.Log[message.Channel] = append(p.Log[message.Channel], message)
				return true
			}
		}
		p.Bot.SendMessage(message.Channel, "Sorry, I don't know that phrase.")
	}
	p.Log[message.Channel] = append(p.Log[message.Channel], message)
	return false
}

// LoadData imports any configuration data into the plugin. This is not strictly necessary other
// than the fact that the Plugin interface demands it exist. This may be deprecated at a later
// date.
func (p *RememberPlugin) LoadData() {
	p.Coll = p.Bot.Db.C("remember")
	rand.Seed(time.Now().Unix())
}

// Help responds to help requests. Every plugin must implement a help function.
func (p *RememberPlugin) Help(channel string, parts []string) {
	msg := "!remember will let you quote your idiot friends. Just type !remember <nick>" +
		" <snippet> to remember what they said. Snippet can be any part of their message. " +
		"Later on, you can ask for a random !quote."
	p.Bot.SendMessage(channel, msg)
}

func (p *RememberPlugin) record(nick, msg string) {
	message := userRemember{
		Nick:    nick,
		Message: msg,
		Date:    time.Now(),
	}
	p.Coll.Insert(message)
}

// deliver a random quote out of the db.
// Note: this is the same cache for all channels joined. This plugin needs to be expanded
// to have this function execute a quote for a particular channel
func (p *RememberPlugin) randQuote() string {
	var quotes []userRemember
	iter := p.Coll.Find(bson.M{}).Iter()
	err := iter.All(&quotes)
	if err != nil {
		panic(iter.Err())
	}

	// rand quote idx
	nquotes := len(quotes)
	if nquotes == 0 {
		return "Sorry, I don't know any quotes."
	}
	quote := quotes[rand.Intn(nquotes)]
	return quote.Message
}

func (p *RememberPlugin) quoteTimer(channel string) {
	for {
		// this pisses me off: You can't multiply int * time.Duration so it
		// has to look ugly as shit.
		time.Sleep(time.Duration(p.Bot.Config.QuoteTime) * time.Minute)
		chance := 1.0 / p.Bot.Config.QuoteChance
		if rand.Intn(int(chance)) == 0 {
			msg := p.randQuote()
			fmt.Println("Delivering quote.")
			p.Bot.SendMessage(channel, msg)
		}
	}
}

// Empty event handler because this plugin does not do anything on event recv
func (p *RememberPlugin) Event(kind string, message bot.Message)  bool {
	return false
}
