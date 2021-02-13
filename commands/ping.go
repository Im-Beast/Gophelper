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

		messageTime, _ := message.Timestamp.Parse()
		timeDiff := (time.Now().UnixNano() - messageTime.UnixNano()) / 1000000

		msg, err := session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("🏓 Pong\n Discord: `%d ms`", timeDiff))

		if err != nil {
			fmt.Println("Error on ping command when sending message")
		}

		messageTime2, _ := msg.Timestamp.Parse()
		timeDiff2 := (messageTime2.UnixNano() - messageTime.UnixNano()) / 1000000

		_, err = session.ChannelMessageEdit(msg.ChannelID, msg.ID, fmt.Sprintf("🏓 Pong\n Discord: `%d ms`\n Took: `%d ms`", timeDiff, timeDiff2))

		if err != nil {
			fmt.Println("Error on ping command when editing message")
		}
	},
}
