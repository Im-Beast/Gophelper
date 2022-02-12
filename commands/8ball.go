package commands

import (
	"log"
	"math/rand"
	"time"

	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"
)

// EightBall doggies
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
		lang := context.Command.LanguageSettings
		answers := lang.Answers

		var err error

		if len(args) == 0 {
			_, err = session.ChannelMessageSend(event.ChannelID, lang.NoArgumentsMessage)
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
