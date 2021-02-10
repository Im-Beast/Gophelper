package commands

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	gophelper "../Gophelper"
	"github.com/bwmarrin/discordgo"
)

var kisses = []string{
	"https://pa1.narvii.com/5823/f10cce909b5bfa6f05f0af496558a16ed4840c06_hq.gif",
	"https://media1.tenor.com/images/78095c007974aceb72b91aeb7ee54a71/tenor.gif",
	"https://media.giphy.com/media/G3va31oEEnIkM/giphy.gif",
	"https://media.giphy.com/media/y0H514IGMusQE/giphy.gif",
	"https://media1.tenor.com/images/d307db89f181813e0d05937b5feb4254/tenor.gif",
	"https://data.whicdn.com/images/166496706/original.gif",
}

// Kiss :*
var Kiss = &gophelper.Command{
	ID: "Kiss",

	Name:    "😘 Kiss",
	Aliases: []string{"kiss"},

	Description: "Kiss someone or get kissed :*",

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

		index := rand.Intn(len(kisses))

		url := kisses[index]

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
			session.ChannelMessageSend(message.ChannelID, routerLanguage.MessageSendError)
		} else {
			session.MessageReactionAdd(message.ChannelID, message.ID, "😘")
		}
	},
}
