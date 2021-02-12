package commands

import (
	"fmt"
	"time"

	gophelper "../Gophelper"
	middleware "../Middleware"
	"github.com/bwmarrin/discordgo"
)

//Stats statistics about user
var Stats = &gophelper.Command{
	ID: "Stats",

	Name:    "📑 Stats",
	Aliases: []string{"stats"},

	Category: gophelper.CATEGORY_MOD,

	Description: "Gives you information about user",

	Usage: "stats [_user{mention/id}]",

	RateLimit: middleware.RateLimit{
		Limit:    1,
		Duration: time.Second * 5,
	},

	Handler: func(context *gophelper.CommandContext) {
		session := context.Session
		message := context.Event
		arguments := context.Arguments

		language := context.Command.LanguageSettings
		routerLanguage := context.Router.Config.Language

		memberID := message.Author.ID

		if len(arguments) > 0 {
			memberID = arguments[0]

			if gophelper.IsMention(memberID) {
				memberID = gophelper.MentionToID(memberID)
			}
		}

		member, err := session.GuildMember(message.GuildID, memberID)

		if err != nil {
			_, err := session.ChannelMessageSend(message.ChannelID, language.UserNotFound)
			if err != nil {
				fmt.Println("Error on stats command when sending message")
			}
			return
		}

		creationDate, _ := discordgo.SnowflakeTimestamp(memberID)
		joinDate, _ := member.JoinedAt.Parse()

		embed := &discordgo.MessageEmbed{
			Color: 0x007d9c,
			Title: fmt.Sprintf(language.Title, member.User.Username, member.User.Discriminator),
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: member.User.AvatarURL("512"),
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   language.CreationDate,
					Value:  creationDate.Format("2 January 2006 **15:04**"),
					Inline: false,
				},
				{
					Name:   language.JoinDate,
					Value:  joinDate.Format("2 January 2006 **15:04**"),
					Inline: false,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: gophelper.RandomStringElement(routerLanguage.FunFacts),
			},
		}

		_, err = session.ChannelMessageSendEmbed(message.ChannelID, embed)

		if err != nil {
			fmt.Println("Error on stats command when sending message")
		}
	},
}
