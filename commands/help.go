package commands

import (
	"fmt"
	"strings"
	"time"

	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"
	"github.com/Im-Beast/Gophelper/utils"
	"github.com/bwmarrin/discordgo"
)

var cachedMainHelp *discordgo.MessageEmbed
var cachedCategoryPages map[*gophelper.Category][]*discordgo.MessageEmbed

var Help = &gophelper.Command{
	ID: "Help",

	Name:    "📜 Help",
	Aliases: []string{"help"},

	Category: gophelper.CATEGORY_MISC,

	Description: "Get some help",

	Usage: "help [_category [_page]/_command]",

	RateLimit: middleware.RateLimit{
		Limit:    1,
		Duration: time.Second * 1,
	},

	Handler: func(context *gophelper.CommandContext) {
		session := context.Session
		event := context.Event
		router := context.Router
		langCfg := context.Router.Config.Language
		arguments := context.Arguments
		page := 1

		switch args := len(arguments); args {
		case 0:
			msg, err := session.ChannelMessageSendEmbed(event.ChannelID, getMainHelp(router.Categories, &router.Config.Language))

			if err != nil {
				fmt.Printf("err: %s", err.Error())
				return
			}

			handleMainPageReactions(session, router, event, msg, &langCfg)
		case 1:
			embed, err := getStringCommandHelp(arguments[0], router, &langCfg)

			if err == nil {
				session.ChannelMessageSendEmbed(event.ChannelID, embed)
				return
			}
			fallthrough
		default:
			category, err := gophelper.StringToCategory(arguments[0])

			if err != nil {
				fmt.Println("invalid category")
				return
			}

			if args > 1 && utils.IsNumber(arguments[1]) {
				page = utils.StringToInt(arguments[1])
			}

			pages := getCategoryPages(router, &langCfg)
			embed := pages[category][page]

			session.ChannelMessageSendEmbed(event.ChannelID, embed)
		}
	},
}

func getCategoryPages(router *gophelper.Router, langCfg *gophelper.LanguageConfig) map[*gophelper.Category][]*discordgo.MessageEmbed {
	if cachedCategoryPages != nil {
		return cachedCategoryPages
	}

	embeds := make(map[*gophelper.Category][]*discordgo.MessageEmbed)
	helpLang := langCfg.Commands["Help"]

	for _, cmd := range router.Commands {
		category := cmd.Category

		page := len(embeds[category]) - 1

		if page < 0 {
			embeds[category] = append(embeds[category], &discordgo.MessageEmbed{})
			page = 0
		}

		fields := len(embeds[category][page].Fields)

		// page size
		if fields%9 == 0 {
			page++
			embeds[category] = append(embeds[category], &discordgo.MessageEmbed{
				Title:       fmt.Sprintf(helpLang.Embed.CategoryTitle, category.Name),
				Description: category.Description,
				Color:       0x007d9c,
				Footer: &discordgo.MessageEmbedFooter{
					Text: utils.RandomStringElement(langCfg.FunFacts),
				},
			})
		}

		embeds[category][page].Fields = append(embeds[category][page].Fields,
			([]*discordgo.MessageEmbedField{
				{
					Name:   helpLang.Embed.Name,
					Value:  fmt.Sprintf("```%s```", utils.GetStringVal(cmd.Name, helpLang.Embed.NoName)),
					Inline: true,
				},
				{
					Name:   helpLang.Embed.Description,
					Value:  fmt.Sprintf("```%s```", utils.GetStringVal(cmd.Description, helpLang.Embed.NoDescription)),
					Inline: true,
				},
				{
					Name:   helpLang.Embed.Aliases,
					Value:  fmt.Sprintf("```%s```", utils.GetStringVal(strings.Join(cmd.Aliases, ", "), helpLang.Embed.NoAliases)),
					Inline: true,
				},
			})...,
		)
	}

	cachedCategoryPages = embeds
	return embeds
}

func getMainHelp(categories []*gophelper.Category, lang *gophelper.LanguageConfig) *discordgo.MessageEmbed {
	if cachedMainHelp != nil {
		return cachedMainHelp
	}

	helpLanguage := lang.Commands["Help"]

	embed := &discordgo.MessageEmbed{
		Title:       helpLanguage.Embed.Main.Title,
		Description: helpLanguage.Embed.Main.Description,
		Color:       0x007d9c,
		Footer: &discordgo.MessageEmbedFooter{
			Text: utils.RandomStringElement(lang.FunFacts),
		},
	}

	for _, category := range categories {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("%s %s", category.ReactionEmoji, category.Name),
			Value:  fmt.Sprintf("*`%s`*", category.Description),
			Inline: true,
		})
	}

	cachedMainHelp = embed
	return embed
}

func getStringCommandHelp(strCmd string, router *gophelper.Router, langCfg *gophelper.LanguageConfig) (*discordgo.MessageEmbed, error) {
	var searchedCmd *gophelper.Command

	for _, cmd := range router.Commands {
		if strCmd == cmd.Name {
			searchedCmd = cmd
			break
		}

		for _, alias := range cmd.Aliases {
			if strCmd == alias {
				searchedCmd = cmd
				break
			}
		}
	}

	if searchedCmd == nil {
		return nil, fmt.Errorf("command not found")
	}

	return getCommandHelp(searchedCmd, langCfg), nil
}

func getCommandHelp(cmd *gophelper.Command, langCfg *gophelper.LanguageConfig) *discordgo.MessageEmbed {
	helpLang := langCfg.Commands["Help"]

	embed := &discordgo.MessageEmbed{
		Title:       utils.GetStringVal(cmd.Name, helpLang.Embed.NoName),
		Description: utils.GetStringVal(cmd.Description, helpLang.Embed.NoName),
		Color:       0x007d9c,
		Footer: &discordgo.MessageEmbedFooter{
			Text: utils.RandomStringElement(langCfg.FunFacts),
		},
	}

	if cmd.Usage != "" {
		embed.Fields = []*discordgo.MessageEmbedField{
			{
				Name:  langCfg.Commands["Help"].Embed.Usage,
				Value: cmd.Usage,
			},
		}
	}

	return embed
}

func handleCategoryPageReactions(session *discordgo.Session, router *gophelper.Router, origEvent *discordgo.MessageCreate, msg *discordgo.Message, langCfg *gophelper.LanguageConfig, category *gophelper.Category) {
	go func() {
		err := session.MessageReactionsRemoveAll(msg.ChannelID, msg.ID)

		if err != nil {
			fmt.Printf("Failed to remove all emojis")
		}

		emojis := [3]string{"⬅️", "➡️", "⬇️"}
		for _, emoji := range emojis {
			err := session.MessageReactionAdd(msg.ChannelID, msg.ID, emoji)
			if err != nil {
				fmt.Printf("Failed to add emoji %s", emoji)
			}
		}
	}()

	pages := getCategoryPages(router, langCfg)[category]
	length := len(pages)
	page := 1

	refreshPage := func() {
		embed := pages[page]
		_, err := session.ChannelMessageEditEmbed(msg.ChannelID, msg.ID, embed)

		if err != nil {
			fmt.Printf("Failed to edit embed message %s\n", err.Error())
		}
	}

	refreshPage()

	var closeHandler func()
	handler := func(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
		if event.UserID != session.State.User.ID {
			err := session.MessageReactionRemove(event.ChannelID, event.MessageID, event.Emoji.Name, event.UserID)
			if err != nil {
				fmt.Printf("Failed to remove reaction %s", event.Emoji.Name)
			}
		}

		if event.MessageID != msg.ID || event.UserID != origEvent.Author.ID {
			return
		}

		switch event.Emoji.Name {
		case "⬅️":
			page = utils.ClampInt(page-1, 1, length-1)
			refreshPage()
		case "➡️":
			page = utils.ClampInt(page+1, 1, length-1)
			refreshPage()
		case "⬇️":
			closeHandler()
			handleMainPageReactions(session, router, origEvent, msg, langCfg)
		}
	}
	closeHandler = session.AddHandler(handler)
}

func handleMainPageReactions(session *discordgo.Session, router *gophelper.Router, origEvent *discordgo.MessageCreate, msg *discordgo.Message, langCfg *gophelper.LanguageConfig) {
	_, err := session.ChannelMessageEditEmbed(msg.ChannelID, msg.ID, getMainHelp(router.Categories, &router.Config.Language))

	if err != nil {
		fmt.Printf("err: %s", err.Error())
		return
	}

	go func() {
		err = session.MessageReactionsRemoveAll(msg.ChannelID, msg.ID)

		if err != nil {
			fmt.Printf("Failed to remove all emojis")
		}

		for _, category := range router.Categories {
			err := session.MessageReactionAdd(msg.ChannelID, msg.ID, category.ReactionEmoji)
			if err != nil {
				fmt.Printf("Couldn't add emoji for category %s", category.Name)
			}
		}
	}()

	var closeHandler func()
	handler := func(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
		if event.UserID != session.State.User.ID {
			err = session.MessageReactionRemove(event.ChannelID, event.MessageID, event.Emoji.Name, event.UserID)
			if err != nil {
				fmt.Printf("Failed to remove reaction %s", event.Emoji.Name)
			}
		}

		if event.MessageID != msg.ID || event.UserID != origEvent.Author.ID {
			return
		}

		for _, category := range router.Categories {
			if category.ReactionEmoji != event.Emoji.Name {
				continue
			}

			closeHandler()
			handleCategoryPageReactions(session, router, origEvent, msg, langCfg, category)
			break
		}
	}

	closeHandler = session.AddHandler(handler)
}
