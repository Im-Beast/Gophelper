package commands

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	gophelper "../Gophelper"
	"github.com/bwmarrin/discordgo"
)

var hugs = []string{
	"https://media1.tenor.com/images/074d69c5afcc89f3f879ca473e003af2/tenor.gif?itemid=4898650",
	"https://media1.tenor.com/images/108c2257683620292f4687262f26e872/tenor.gif?itemid=17258498",
	"https://media1.tenor.com/images/8ac5ada8524d767b77d3d54239773e48/tenor.gif?itemid=16334628",
	"https://i.imgur.com/rioNdmc.gif",
	"https://i.gifer.com/2QEa.gif",
	"https://acegif.com/wp-content/uploads/anime-hug.gif",
	"https://media.tenor.com/images/b6d0903e0d54e05bb993f2eb78b39778/tenor.gif",
	"https://data.whicdn.com/images/219995514/original.gif",
	"https://images-ext-2.discordapp.net/external/rxstxw_1DcDfXP2ZTHcq5Fk4um5Q57mXsF7Klwyz6Q4/https/data.whicdn.com/images/334957046/original.gif",
	"https://media.tenor.co/videos/161dd2416944be9249bcd0a6d69e0463/mp4",
}

// Hug hug
var Hug = &gophelper.Command{
	ID: "Hug",

	Name:    "🤗 Hug",
	Aliases: []string{"hug"},

	Description: "Hug someone or get hugged	",

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

		index := rand.Intn(len(hugs))

		url := hugs[index]

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
			session.MessageReactionAdd(message.ChannelID, message.ID, "🤗")
		}
	},
}
