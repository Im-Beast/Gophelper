package generics

import (
	"fmt"
	"log"
	"strings"
	"time"

	gophelper "github.com/Im-Beast/Gophelper/internal"
	"github.com/Im-Beast/Gophelper/utils"
	"github.com/bwmarrin/discordgo"
)

func ImageResponseHandler(context *gophelper.CommandContext, cmdName string, responses []string) {
	args := context.Arguments
	event := context.Event
	session := context.Session
	lang := context.Router.Config.Language
	language := context.Command.LanguageSettings

	title := language.Response.NonMention

	var err error

	if len(args) > 0 {
		userID := args[0]

		if utils.IsMention(userID) {
			userID = utils.MentionToID(userID)
		}

		member, err := session.GuildMember(event.GuildID, userID)

		nick := event.Author.Username

		if err == nil {
			nick = member.User.Username
		}

		title = fmt.Sprintf(language.Response.Mention, nick)
	}

	url := utils.RandomStringElement(responses)

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

	msg, err := session.ChannelMessageSendEmbed(event.ChannelID, embed)

	if err != nil {
		log.Printf("Command %s errored while sending embed: %s", cmdName, err.Error())
		_, err = session.ChannelMessageSend(event.ChannelID, lang.Errors.MessageSend)
		log.Printf("Command %s errored while sending error message: %s", cmdName, err.Error())
	} else {
		err = session.MessageReactionAdd(msg.ChannelID, msg.ID, "🐕")
		if err != nil {
			log.Printf("Command %s errored while reacting to a message: %s", cmdName, err.Error())
		}
	}
}
