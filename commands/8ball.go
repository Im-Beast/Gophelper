package commands

import (
	"log"
	"math/rand"
	"time"

	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"
)

// Command which answers your questions 🔮
var EightBall = &gophelper.Command{
	ID: "8Ball",

	Name:    "🎱 8Ball",
	Aliases: []string{"8ball"},

	Category: gophelper.CATEGORY_FUN,

	Description: "I'll tell you truth",

	RateLimit: middleware.RateLimit{
		Limit:    1,
		Duration: time.Second * 5,
	},

	Handler: func(context *gophelper.CommandContext) {
		args := context.Arguments
		session := context.Session
		event := context.Event
		cmdLang := context.Router.Config.Language.Commands[context.Command.ID]
		answers := cmdLang.Answers

		var err error

		if len(args) == 0 {
			_, err = session.ChannelMessageSend(event.ChannelID, cmdLang.NoArgumentsMessage)
		} else {
			index := rand.Intn(len(answers))
			randomAnswer := answers[index]
			_, err = session.ChannelMessageSend(event.ChannelID, randomAnswer)
		}

		if err != nil {
			log.Printf("Command 8Ball errored while sending a message: %s", err.Error())
		}
	},
}
