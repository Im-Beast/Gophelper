package middleware

import (
	"fmt"
	"strings"

	gophelper "../Gophelper"
	"github.com/bwmarrin/discordgo"
)

var (
	// MainHelpEmbed is main help page that shows categories
	MainHelpEmbed = &discordgo.MessageEmbed{}
	// HelpCategoryEmbeds contain embeds used to show per-category help page
	HelpCategoryEmbeds = make(map[*gophelper.Category][]*discordgo.MessageEmbed)
	// HelpEmbeds contain help embeds about all commands (stored via all aliases)
	HelpEmbeds = make(map[string]*discordgo.MessageEmbed)
	// HelpStringCategories contains categories sorted by their aliases
	HelpStringCategories = make(map[string]*gophelper.Category)
)

// HelpInitMiddleware prrt
func HelpInitMiddlware(context *gophelper.CommandContext) (bool, func(*gophelper.CommandContext)) {
	addHelpForCommand(context.Command, &context.Router.Config.Language)
	generateMainEmbed(&context.Router.Config.Language)
	return true, nil
}

func generateMainEmbed(language *gophelper.LanguageConfig) {
	helpLanguage := language.Commands["Help"]

	embed := &discordgo.MessageEmbed{
		Title:       helpLanguage.Embed.Main.Title,
		Description: helpLanguage.Embed.Main.Description,
		Color:       0x007d9c,
		Footer: &discordgo.MessageEmbedFooter{
			Text: gophelper.RandomStringElement(language.FunFacts),
		},
	}

	for category := range HelpCategoryEmbeds {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("%s %s", category.ReactionEmoji, category.Name),
			Value:  fmt.Sprintf("*`%s`*", category.Description),
			Inline: true,
		})
	}

	MainHelpEmbed = embed
}

func addHelpForCommand(command *gophelper.Command, language *gophelper.LanguageConfig) {
	helpLanguage := language.Commands["Help"]

	category := command.Category
	embeds := HelpCategoryEmbeds[category]

	go func() {
		for _, alias := range category.Aliases {
			HelpStringCategories[strings.ToLower(alias)] = category
		}
	}()

	var embed *discordgo.MessageEmbed
	var index int = 0

	if len(embeds) > 0 {
		index = len(embeds) - 1
		embed = embeds[index]
	} else {
		HelpCategoryEmbeds[category] = append(HelpCategoryEmbeds[category], getBasicEmbed(category, language))
		embed = HelpCategoryEmbeds[category][0]
	}

	if len(embed.Fields) >= 4*3 { //auto page
		HelpCategoryEmbeds[category] = append(HelpCategoryEmbeds[category], getBasicEmbed(category, language))
		addHelpForCommand(command, language)
		return
	}

	for _, alias := range command.Aliases {
		HelpEmbeds[strings.ToLower(alias)] = getCommandEmbed(command, &helpLanguage)
	}

	embed.Fields = append(embed.Fields, getCommandEmbedFields(command, &helpLanguage)...)

	HelpCategoryEmbeds[category][index].Fields = embed.Fields
}

func getBasicEmbed(category *gophelper.Category, language *gophelper.LanguageConfig) *discordgo.MessageEmbed {
	helpLanguage := language.Commands["Help"]
	return &discordgo.MessageEmbed{
		Title:       fmt.Sprintf(helpLanguage.Embed.CategoryTitle, category.Name),
		Description: category.Description,
		Color:       0x007d9c,
		Footer: &discordgo.MessageEmbedFooter{
			Text: gophelper.RandomStringElement(language.FunFacts),
		},
	}
}
func getCommandEmbed(command *gophelper.Command, language *gophelper.CommandConfig) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{}

	embed.Title = gophelper.GetStringVal(command.Name, language.Embed.NoName)
	embed.Description = gophelper.GetStringVal(command.Description, language.Embed.NoName)

	if command.Usage != "" {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  language.Embed.Usage,
			Value: command.Usage,
		})
	}

	return embed
}

func getCommandEmbedFields(command *gophelper.Command, language *gophelper.CommandConfig) []*discordgo.MessageEmbedField {
	return []*discordgo.MessageEmbedField{
		{
			Name:   language.Embed.Name,
			Value:  fmt.Sprintf("```%s```", gophelper.GetStringVal(command.Name, language.Embed.NoName)),
			Inline: true,
		},
		{
			Name:   language.Embed.Description,
			Value:  fmt.Sprintf("```%s```", gophelper.GetStringVal(command.Description, language.Embed.NoDescription)),
			Inline: true,
		},
		{
			Name:   language.Embed.Aliases,
			Value:  fmt.Sprintf("```%s```", gophelper.GetStringVal(strings.Join(command.Aliases, ", "), language.Embed.NoAliases)),
			Inline: true,
		},
	}
}
