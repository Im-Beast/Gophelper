package commands

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"
	"github.com/Im-Beast/Gophelper/utils"
	"github.com/bwmarrin/discordgo"
)

var kitties = []string{
	"http://www.randomkittengenerator.com/cats/rotator.php",
}

// Kitty kitties
var Kitty = &gophelper.Command{
	ID: "Kitty",

	Name:    "🐈 Kitty",
	Aliases: []string{"kitty", "kittie", "cat"},

	Category: gophelper.CATEGORY_FUN,

	Description: "random pics of cute kitties",

	RateLimit: middleware.RateLimit{
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

			if utils.IsMention(userID) {
				userID = utils.MentionToID(userID)
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

		index := rand.Intn(len(kitties))

		url := kitties[index]

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
			_, err = session.ChannelMessageSend(message.ChannelID, routerLanguage.Errors.MessageSend)
			fmt.Println("Failed to send message")
		} else {
			err = session.MessageReactionAdd(message.ChannelID, message.ID, "🐈")
		}

		if err != nil {
			fmt.Println("Error on kitty command when reacting/sending message")
		}
	},
}
