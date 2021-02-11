package commands

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	gophelper "../Gophelper"
	"github.com/bwmarrin/discordgo"
)

var doggies = []string{
	"https://www.randomdoggiegenerator.com/randomdoggie.php",
}

// Doggie doggies
var Doggie = &gophelper.Command{
	ID: "Doggie",

	Name:    "🐕 Doggie",
	Aliases: []string{"doggie", "doggy", "dog"},

	Description: "Random pics of cute doggies",

	RateLimit: gophelper.RateLimit{
		Limit:    1,
		Duration: time.Second * 5,
	},

	Handler: func(context *gophelper.CommandContext) {
		arguments := context.Arguments
		message := context.Event.Message
		session := context.Session

		routerLanguage := context.Router.Config.Language
		language := context.Command.LanguageSettings

		isTag := false
		userID := message.Author.ID

		if len(arguments) > 0 {
			isTag = true
			userID = arguments[0]

			if gophelper.IsMention(userID) {
				userID = gophelper.MentionToID(userID)
			}
		}

		title := language.Response.NonMention

		if isTag {
			member, err := session.GuildMember(message.GuildID, userID)

			nick := message.Author.Username

			if err == nil {
				nick = member.User.Username
			}

			title = fmt.Sprintf(language.Response.Mention, nick)
		}

		index := rand.Intn(len(doggies))

		url := doggies[index]

		if strings.Contains(url, "?") {
			url = fmt.Sprintf("%s&nocache=%d", url, time.Now().UnixNano())
		} else {
			url = fmt.Sprintf("%s?nocache=%d", url, time.Now().UnixNano())
		}

		embed := &discordgo.MessageEmbed{
			Title: title,
			Color: 0xFFbbbb,
			Provider: &discordgo.MessageEmbedProvider{
				URL: url,
			},
			Image: &discordgo.MessageEmbedImage{
				URL: url,
			},
		}

		message, err := session.ChannelMessageSendEmbed(message.ChannelID, embed)

		if err != nil {
			session.ChannelMessageSend(message.ChannelID, routerLanguage.Errors.MessageSend)
		} else {
			session.MessageReactionAdd(message.ChannelID, message.ID, "❤️")
			session.MessageReactionAdd(message.ChannelID, message.ID, "🐕")
		}
	},
}
