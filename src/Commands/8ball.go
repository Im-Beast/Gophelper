package commands

import (
	"math/rand"
	"time"

	gophelper "../Gophelper"
)

// EightBall doggies
var EightBall = &gophelper.Command{
	ID: "8Ball",

	Name:    "🎱 8Ball",
	Aliases: []string{"8ball"},

	Description: "I'll tell you truth",

	RateLimit: gophelper.RateLimit{
		Limit:    1,
		Duration: time.Second * 5,
	},

	Handler: func(context *gophelper.CommandContext) {
		arguments := context.Arguments
		session := context.Session
		message := context.Event

		language := context.Command.LanguageSettings

		answers := language.Answers

		if len(arguments) == 0 {
			session.ChannelMessageSend(message.ChannelID, language.NoArgumentsMessage)
		} else {
			index := rand.Intn(len(answers))
			randomAnswer := answers[index]
			session.ChannelMessageSend(message.ChannelID, randomAnswer)
		}
	},
}
