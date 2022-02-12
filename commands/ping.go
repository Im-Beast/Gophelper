package commands

import (
	"fmt"
	"log"
	"time"

	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"
)

// Ping pong
var Ping = &gophelper.Command{
	ID: "Ping",

	Name:    "🏓 Ping",
	Aliases: []string{"ping"},

	Category: gophelper.CATEGORY_MISC,

	Description: "Literally pong",

	RateLimit: middleware.RateLimit{
		Limit:    1,
		Duration: time.Second * 5,
	},

	Handler: func(context *gophelper.CommandContext) {
		session := context.Session
		message := context.Event

		msgTime, err := message.Timestamp.Parse()
		if err != nil {
			log.Printf("Command Ping errored while parsing message timestamp: %s\n", err.Error())
			return
		}

		timeDiff := time.Since(msgTime).Milliseconds()

		msg, err := session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("🏓 Pong\n Discord: `%d ms`", timeDiff))
		if err != nil {
			log.Printf("Command Ping errored while sending message: %s\n", err.Error())
			return
		}

		msgTime2, err := msg.Timestamp.Parse()
		if err != nil {
			log.Printf("Command Ping errored while parsing message timestamp: %s\n", err.Error())
			return
		}

		timeDiff2 := msgTime2.Sub(msgTime).Milliseconds()

		_, err = session.ChannelMessageEdit(msg.ChannelID, msg.ID, fmt.Sprintf("🏓 Pong\n Discord: `%d ms`\n Took: `%d ms`", timeDiff, timeDiff2))
		if err != nil {
			log.Printf("Command Ping errored while editing message: %s\n", err.Error())
		}
	},
}
