package commands

import (
	"fmt"
	"time"

	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"
)

// Hentai cartoons
var Hentai = &gophelper.Command{
	ID: "Hentai",

	Name:    "🐙 Hentai",
	Aliases: []string{"hentai"},

	Category: gophelper.CATEGORY_FUN,

	Description: "Just japanese cartoons",

	RateLimit: middleware.RateLimit{
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
			_, err = session.ChannelMessageSend(message.ChannelID, "ur still horni :|")
		} else {
			_, err = session.ChannelMessageSend(message.ChannelID, "u horni >:( its nsfw only")
		}

		if err != nil {
			fmt.Println("Error on hentai command when sending message")
		}
	},
}
