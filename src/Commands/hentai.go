package commands

import (
	"time"

	gophelper "../Gophelper"
)

// Hentai cartoons
var Hentai = &gophelper.Command{
	ID: "Hentai",

	Name:    "🐙 Hentai",
	Aliases: []string{"hentai"},

	Description: "Just japanese cartoons",

	RateLimit: gophelper.RateLimit{
		Limit:    1,
		Duration: time.Second * 5,
	},

	Handler: func(context *gophelper.CommandContext) {
		session := context.Session
		message := context.Event

		channel, err := session.Channel(message.ChannelID)

		if err != nil {
			return
		}

		if channel.NSFW {
			session.ChannelMessageSend(message.ChannelID, "ur still horni :|")
		} else {
			session.ChannelMessageSend(message.ChannelID, "u horni >:( its nsfw only")
		}
	},
}
