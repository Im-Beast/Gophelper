package commands

import (
	"fmt"
	"strings"
	"time"

	gophelper "../Gophelper"
	middleware "../Middleware"
	"github.com/bwmarrin/discordgo"
)

// Help shows help
var Help = &gophelper.Command{
	ID: "Help",

	Name:    "📜 Help",
	Aliases: []string{"help"},

	Category: gophelper.CATEGORY_MISC,

	Description: "Get some help",

	Usage: "help [_category [_page]/_command]",

	RateLimit: middleware.RateLimit{
		Limit:    1,
		Duration: time.Second * 30,
	},

	Handler: func(context *gophelper.CommandContext) {
		session := context.Session
		message := context.Event
		router := context.Router
		arguments := context.Arguments
		language := context.Command.LanguageSettings

		var (
			embed           *discordgo.MessageEmbed
			reactMessage    *discordgo.Message
			closeHandler    func()
			page            int
			pages           int
			categoryAlias   string
			err             error
			categoryHandler func(session *discordgo.Session, event *discordgo.MessageReactionAdd)
			pageHandler     func(session *discordgo.Session, event *discordgo.MessageReactionAdd)

			expireCooldown int = 15
			expireTimer    int = expireCooldown
		)

		pageHandler = func(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
			if event.UserID != session.State.User.ID {
				err := session.MessageReactionRemove(event.ChannelID, event.MessageID, event.Emoji.Name, event.UserID)
				if err != nil {
					fmt.Printf("Failed to remove reaction %s", event.Emoji.Name)
				}
			}

			if event.MessageID != reactMessage.ID || event.UserID != message.Author.ID {
				return
			}

			switch event.Emoji.Name {
			case "⬅️":
				page = gophelper.ClampInt(page-1, 0, pages-1)
			case "➡️":
				page = gophelper.ClampInt(page+1, 0, pages-1)
			case "⬇️":
				_, err = session.ChannelMessageEditEmbed(reactMessage.ChannelID, reactMessage.ID, middleware.MainHelpEmbed)
				if err != nil {
					fmt.Println("Failed to edit embed message")
				}

				go addCategoryReactions(context, reactMessage.ChannelID, reactMessage.ID)

				closeHandler()
				closeHandler = session.AddHandler(categoryHandler)
				return
			default:
				return
			}

			expireTimer = expireCooldown
			embed, pages = getCategoryEmbed(categoryAlias, page, &language)
			_, err = session.ChannelMessageEditEmbed(reactMessage.ChannelID, reactMessage.ID, embed)
			if err != nil {
				fmt.Println("Error on editing embed", err)
			}
		}

		categoryHandler = func(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
			if event.UserID != session.State.User.ID {
				err = session.MessageReactionRemove(event.ChannelID, event.MessageID, event.Emoji.Name, event.UserID)
				if err != nil {
					fmt.Printf("Failed to remove reaction %s", event.Emoji.Name)
				}
			}

			if event.MessageID != reactMessage.ID || event.UserID != message.Author.ID {
				return
			}

			for _, category := range router.Categories {
				if category.ReactionEmoji == event.Emoji.Name {
					categoryAlias = category.Aliases[0]
					embed, pages = getCategoryEmbed(categoryAlias, 0, &language)

					expireTimer = expireCooldown

					_, err = session.ChannelMessageEditEmbed(reactMessage.ChannelID, reactMessage.ID, embed)

					if err != nil {
						fmt.Println("Failed to edit embed message")
					}

					go addPageReactions(session, reactMessage.ChannelID, reactMessage.ID)
					closeHandler()
					closeHandler = session.AddHandler(pageHandler)
					return
				}
			}
		}

		if len(arguments) == 0 {
			reactMessage, err = session.ChannelMessageSendEmbed(message.ChannelID, middleware.MainHelpEmbed)

			if err != nil {
				fmt.Println("Error while sending embed:", err)
				return
			}

			go addCategoryReactions(context, reactMessage.ChannelID, reactMessage.ID)

			closeHandler = session.AddHandler(categoryHandler)
		} else {
			name := arguments[0]

			if len(arguments) > 1 && gophelper.IsNumber(arguments[1]) {
				page = gophelper.StringToInt(arguments[1]) - 1 // count from 1 for user conveniency
			}

			embed, _ = getCategoryEmbed(name, page, &language)
			if embed == nil {
				embed = getCommandEmbed(name)
				if embed == nil {
					return
				}
			}

			reactMessage, err = session.ChannelMessageSendEmbed(message.ChannelID, embed)
			go addPageReactions(session, reactMessage.ChannelID, reactMessage.ID)
			closeHandler = session.AddHandler(pageHandler)
		}

		ticker := time.NewTicker(time.Second)

		for range ticker.C {
			if expireTimer <= 0 {
				err = session.MessageReactionsRemoveAll(reactMessage.ChannelID, reactMessage.ID)
				err = session.MessageReactionAdd(reactMessage.ChannelID, reactMessage.ID, "⛔")

				closeHandler()
				ticker.Stop()
			}

			expireTimer--
		}

	},
}

func addCategoryReactions(context *gophelper.CommandContext, ChannelID string, MessageID string) {
	session := context.Session
	router := context.Router

	err := session.MessageReactionsRemoveAll(ChannelID, MessageID)

	if err != nil {
		fmt.Printf("Failed to remove all emojis")
	}

	for _, category := range router.Categories {
		err := session.MessageReactionAdd(ChannelID, MessageID, category.ReactionEmoji)
		if err != nil {
			fmt.Printf("Couldn't add emoji for category %s", category.Name)
		}
	}
}

func addPageReactions(session *discordgo.Session, ChannelID string, MessageID string) {
	err := session.MessageReactionsRemoveAll(ChannelID, MessageID)

	if err != nil {
		fmt.Printf("Failed to remove all emojis")
	}

	emojis := [3]string{"⬅️", "➡️", "⬇️"}

	for _, emoji := range emojis {
		err := session.MessageReactionAdd(ChannelID, MessageID, emoji)
		if err != nil {
			fmt.Printf("Failed to add emoji %s", emoji)
		}
	}
}

func getCategoryEmbed(name string, page int, language *gophelper.CommandConfig) (*discordgo.MessageEmbed, int) {
	category := middleware.HelpStringCategories[strings.ToLower(name)]
	categoryEmbeds := middleware.HelpCategoryEmbeds[category]

	if len(categoryEmbeds) > 0 && page >= 0 && page <= len(categoryEmbeds)-1 {
		embed := *categoryEmbeds[page]
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   language.Page,
			Value:  fmt.Sprintf("%d/%d", page+1, len(categoryEmbeds)),
			Inline: false,
		})

		return &embed, len(categoryEmbeds)
	} else {
		return nil, len(categoryEmbeds)
	}
}

func getCommandEmbed(name string) *discordgo.MessageEmbed {
	return middleware.HelpEmbeds[strings.ToLower(name)]
}
