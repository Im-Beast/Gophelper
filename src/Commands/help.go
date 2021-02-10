package commands

import (
	"fmt"
	"strings"
	"time"

	gophelper "../Gophelper"
	"github.com/bwmarrin/discordgo"
)

// Code of this command is pretty shit
// I'll recode it soon so it supports categories and code won't be so shit

// Help shows help
var Help = &gophelper.Command{
	ID: "Help",

	Name:    "📜 Help",
	Aliases: []string{"help"},

	Description: "Get some help",

	Usage: "help [_page/_command]",

	RateLimit: gophelper.RateLimit{
		Limit:    5,
		Duration: time.Second * 30,
	},

	Handler: func(context *gophelper.CommandContext) {
		session := context.Session
		message := context.Event

		routerLanguage := context.Router.Config.Language
		language := context.Command.LanguageSettings

		embed := &discordgo.MessageEmbed{
			Color: 0x007d9c,
		}

		arguments := context.Arguments

		helpForCommand := len(arguments) > 0 && !gophelper.IsNumber(arguments[0])

		if helpForCommand {
			commandName := arguments[0]

			for _, command := range context.Router.Commands {
				for _, alias := range command.Aliases {
					if !gophelper.MatchesPrefix(alias, commandName, false) {
						continue
					}

					name := command.Name
					if name == "" {
						name = language.Embed.NoName
					}

					embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
						Name:   language.Embed.Name,
						Value:  fmt.Sprintf("```%s```", name),
						Inline: true,
					})

					description := command.Description
					if description == "" {
						description = language.Embed.NoDescription
					}

					embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
						Name:   language.Embed.Description,
						Value:  fmt.Sprintf("```%s```", description),
						Inline: true,
					})

					usage := command.Usage
					if usage != "" {
						embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
							Name:   language.Embed.Usage,
							Value:  fmt.Sprintf("```%s```", usage),
							Inline: true,
						})
					}
				}
			}

			if len(embed.Fields) == 0 {
				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
					Name:   language.CommandNotFound.Title,
					Value:  fmt.Sprintf("```%s```", language.CommandNotFound.Message),
					Inline: false,
				})
			}

			embed.Footer = &discordgo.MessageEmbedFooter{
				Text: gophelper.RandomStringElement(routerLanguage.FunFacts),
			}

			session.ChannelMessageSendEmbed(message.ChannelID, embed)
			return
		}

		pages := make(map[int][]*gophelper.Command)

		refreshPage := func(page int, embed *discordgo.MessageEmbed) {
			for _, command := range pages[page] {
				name := command.Name
				if name == "" {
					name = language.Embed.NoName
				}

				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
					Name:   language.Embed.Name,
					Value:  fmt.Sprintf("```%s```", name),
					Inline: true,
				})

				description := command.Description
				if description == "" {
					description = language.Embed.NoDescription
				}

				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
					Name:   language.Embed.Description,
					Value:  fmt.Sprintf("```%s```", description),
					Inline: true,
				})

				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
					Name:   language.Embed.Aliases,
					Value:  fmt.Sprintf("```%s```", strings.Join(command.Aliases, ", ")),
					Inline: true,
				})
			}

			embed.Footer = &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("%s %d/%d", language.Page, page, len(pages)),
			}
		}

		pageCount := 0
		for i, command := range context.Router.Commands {
			if i%4 == 0 {
				pageCount++
			}

			pages[pageCount] = append(pages[pageCount], command)
		}

		currentPage := 1
		if len(arguments) > 0 {
			currentPage = gophelper.StringToInt(arguments[0])
			if currentPage > pageCount {
				currentPage = pageCount
			} else if currentPage <= 0 {
				currentPage = 0
			}
		}

		refreshPage(currentPage, embed)

		reactMessage, err := session.ChannelMessageSendEmbed(message.ChannelID, embed)
		if err != nil {
			fmt.Println("Error while sending help embed", err)
			return
		}

		session.MessageReactionAdd(reactMessage.ChannelID, reactMessage.ID, "⬅️")
		session.MessageReactionAdd(reactMessage.ChannelID, reactMessage.ID, "⏹️")
		session.MessageReactionAdd(reactMessage.ChannelID, reactMessage.ID, "➡️")

		var cancelHelp func()
		isCancelled := false

		closeHandler := session.AddHandler(func(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
			if event.UserID != session.State.User.ID {
				session.MessageReactionRemove(event.ChannelID, event.MessageID, event.Emoji.Name, event.UserID)
			}

			if event.MessageID != reactMessage.ID || event.UserID != message.Author.ID {
				return
			}

			embed.Fields = []*discordgo.MessageEmbedField{}

			switch event.Emoji.Name {
			case "⬅️":
				currentPage--
				break
			case "➡️":
				currentPage++
				break
			case "⏹️":
				cancelHelp()
				return
			default:
				return
			}

			if currentPage > pageCount {
				currentPage = pageCount
			} else if currentPage <= 0 {
				currentPage = 1
			}

			refreshPage(currentPage, embed)
			session.ChannelMessageEditEmbed(reactMessage.ChannelID, reactMessage.ID, embed)
		})

		cancelHelp = func() {
			isCancelled = true
			session.MessageReactionsRemoveAll(reactMessage.ChannelID, reactMessage.ID)
			closeHandler()
			session.MessageReactionAdd(reactMessage.ChannelID, reactMessage.ID, "⛔")
		}

		go func() {
			time.Sleep(time.Second * 30)
			if !isCancelled {
				cancelHelp()
			}
		}()
	},
}
