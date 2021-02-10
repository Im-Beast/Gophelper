package commands

import (
	"fmt"
	"time"

	gophelper "../Gophelper"
)

// Ping pong
var Ping = &gophelper.Command{
	ID: "Ping",

	Name:    "🏓 Ping",
	Aliases: []string{"ping"},

	Description: "Literally pong",

	RateLimit: gophelper.RateLimit{
		Limit:    1,
		Duration: time.Second * 5,
	},

	Handler: func(context *gophelper.CommandContext) {
		session := context.Session
		message := context.Event

		messageTime, _ := message.Timestamp.Parse()
		timeDiff := (time.Now().UnixNano() - messageTime.UnixNano()) / 1000000

		msg, _ := session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("🏓 Pong\n Discord: `%d ms`", timeDiff))

		messageTime2, _ := msg.Timestamp.Parse()
		timeDiff2 := (messageTime2.UnixNano() - messageTime.UnixNano()) / 1000000

		session.ChannelMessageEdit(msg.ChannelID, msg.ID, fmt.Sprintf("🏓 Pong\n Discord: `%d ms`\n Took: `%d ms`", timeDiff, timeDiff2))
	},
}
