package plugins

import "bitbucket.org/phlyingpenguin/godeepintir/bot"

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strings"
)

// This is a counter plugin to count arbitrary things.

type CounterPlugin struct {
	Bot  *bot.Bot
	Coll *mgo.Collection
}

type Item struct {
	Nick  string
	Item  string
	Count int
}

// NewCounterPlugin creates a new CounterPlugin with the Plugin interface
func NewCounterPlugin(bot *bot.Bot) *CounterPlugin {
	return &CounterPlugin{
		Bot:  bot,
		Coll: bot.Db.C("counter"),
	}
}

// Message responds to the bot hook on recieving messages.
// This function returns true if the plugin responds in a meaningful way to the users message.
// Otherwise, the function returns false and the bot continues execution of other plugins.
func (p *CounterPlugin) Message(message bot.Message) bool {
	// This bot does not reply to anything
	nick := message.User.Name
	channel := message.Channel
	parts := strings.Fields(message.Body)

	if len(parts) == 0 {
		return false
	}

	if message.Command && parts[0] == "count" {
		var subject string
		var itemName string

		if len(parts) == 3 {
			// report count for parts[1]
			subject = strings.ToLower(parts[1])
			itemName = strings.ToLower(parts[2])
		} else if len(parts) == 2 {
			subject = strings.ToLower(nick)
			itemName = strings.ToLower(parts[1])
		} else {
			return false
		}

		var item Item
		err := p.Coll.Find(bson.M{"nick": subject, "item": itemName}).One(&item)
		if err != nil {
			p.Bot.SendMessage(channel, fmt.Sprintf("I don't think %s has any %s.",
				subject, itemName))
			return true
		}

		p.Bot.SendMessage(channel, fmt.Sprintf("%s has %d %s.", subject, item.Count, itemName))

		return true
	} else if len(parts) == 1 {
		subject := strings.ToLower(nick)
		itemName := strings.ToLower(parts[0])[:len(parts[0])-2]

		if strings.HasSuffix(parts[0], "++") {
			// ++ those fuckers
			item := p.update(subject, itemName, 1)
			p.Bot.SendMessage(channel, fmt.Sprintf("You have %d %s, %s.", item.Count, item.Item, nick))
			return true
		} else if strings.HasSuffix(parts[0], "--") {
			// -- those fuckers
			item := p.update(subject, itemName, -1)
			p.Bot.SendMessage(channel, fmt.Sprintf("You have %d %s, %s.", item.Count, item.Item, nick))
			return true
		}
	}

	return false
}

func (p *CounterPlugin) update(subject, itemName string, delta int) Item {
	var item Item
	err := p.Coll.Find(bson.M{"nick": subject, "item": itemName}).One(&item)
	if err != nil {
		// insert it
		item = Item{
			Nick:  subject,
			Item:  itemName,
			Count: 1,
		}
		p.Coll.Insert(item)
	} else {
		// update it
		item.Count += delta
		p.Coll.Update(bson.M{"nick": subject, "item": itemName}, item)
	}
	return item
}

// LoadData imports any configuration data into the plugin. This is not strictly necessary other
// than the fact that the Plugin interface demands it exist. This may be deprecated at a later
// date.
func (p *CounterPlugin) LoadData() {
	// This bot has no data to load
}

// Help responds to help requests. Every plugin must implement a help function.
func (p *CounterPlugin) Help(channel string, parts []string) {
	p.Bot.SendMessage(channel, "Sorry, I don't really know what counter does.")
}

// Empty event handler because this plugin does not do anything on event recv
func (p *CounterPlugin) Event(kind string, message bot.Message) bool {
	return false
}
