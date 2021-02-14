package commands

import (
	"fmt"
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

		var timeDiff, timeDiff2 int64
		var msgTime, msgTime2 time.Time

		msgTime, _ = message.Timestamp.Parse()
		timeDiff = time.Since(msgTime).Milliseconds()

		msg, err := session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("🏓 Pong\n Discord: `%d ms`", timeDiff))

		if err != nil {
			fmt.Println("Error on ping command when sending message")
		}

		msgTime2, _ = msg.Timestamp.Parse()
		timeDiff2 = msgTime2.Sub(msgTime).Milliseconds()

		_, err = session.ChannelMessageEdit(msg.ChannelID, msg.ID, fmt.Sprintf("🏓 Pong\n Discord: `%d ms`\n Took: `%d ms`", timeDiff, timeDiff2))

		if err != nil {
			fmt.Println("Error on ping command when editing message")
		}
	},
}
