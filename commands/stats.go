package commands

import (
	"fmt"
	"log"
	"time"

	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"
	"github.com/Im-Beast/Gophelper/utils"
	"github.com/bwmarrin/discordgo"
)

// Command which replies with information about specific user
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
		msg := context.Event
		args := context.Arguments
		lang := context.Router.Config.Language
		cmdLang := lang.Commands[context.Command.ID]

		memberID := msg.Author.ID

		if len(args) > 0 {
			memberID = args[0]

			if utils.IsMention(memberID) {
				memberID = utils.MentionToID(memberID)
			}
		}

		member, err := session.GuildMember(msg.GuildID, memberID)

		if err != nil {
			_, err := session.ChannelMessageSend(msg.ChannelID, cmdLang.UserNotFound)
			if err != nil {
				log.Printf("Command Stats errored while sending error message: %s\n", err.Error())
			}
			return
		}

		creationDate, _ := discordgo.SnowflakeTimestamp(memberID)
		joinDate, _ := member.JoinedAt.Parse()

		embed := &discordgo.MessageEmbed{
			Color: 0x007d9c,
			Title: fmt.Sprintf(cmdLang.Title, member.User.Username, member.User.Discriminator),
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: member.User.AvatarURL("512"),
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   cmdLang.CreationDate,
					Value:  creationDate.Format("2 January 2006 **15:04**"),
					Inline: false,
				},
				{
					Name:   cmdLang.JoinDate,
					Value:  joinDate.Format("2 January 2006 **15:04**"),
					Inline: false,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: utils.RandomStringElement(lang.FunFacts),
			},
		}

		_, err = session.ChannelMessageSendEmbed(msg.ChannelID, embed)

		if err != nil {
			log.Printf("Command Stats errored while sending embed message: %s\n", err.Error())
		}
	},
}
