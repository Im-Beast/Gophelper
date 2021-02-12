package commands

import (
	"fmt"
	"math/rand"
	"time"

	gophelper "../Gophelper"
	middleware "../Middleware"
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
		arguments := context.Arguments
		session := context.Session
		message := context.Event

		language := context.Command.LanguageSettings

		answers := language.Answers

		var err error

		if len(arguments) == 0 {
			_, err = session.ChannelMessageSend(message.ChannelID, language.NoArgumentsMessage)
		} else {
			index := rand.Intn(len(answers))
			randomAnswer := answers[index]
			_, err = session.ChannelMessageSend(message.ChannelID, randomAnswer)
		}

		if err != nil {
			fmt.Println("Error on 8ball command when sending message")
		}
	},
}
