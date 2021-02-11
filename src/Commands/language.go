package commands

import (
	"fmt"
	"time"

	gophelper "../Gophelper"
	"github.com/bwmarrin/discordgo"
)

// LanguageSwitcher command, this command i keep in english and english only for understandable reasons
var LanguageSwitcher = &gophelper.Command{
	Name:    "👅 Language",
	Aliases: []string{"lang"},

	NeededPermissions: discordgo.PermissionAdministrator,

	Description: "Change router language config on the fly (Admin only)",

	RateLimit: gophelper.RateLimit{
		Limit:    1,
		Duration: time.Second * 30,
	},

	Handler: func(context *gophelper.CommandContext) {
		arguments := context.Arguments
		session := context.Session
		message := context.Event

		if len(arguments) == 0 {
			session.ChannelMessageSend(message.ChannelID, "You need to specify language config file")
			return
		}

		config := context.Router.Config

		err := config.LoadLanguage("../Languages/" + arguments[0])
		context.Router.RefreshCommands()

		if err != nil {
			session.ChannelMessageSend(message.ChannelID, "Something happened while loading this config file, are you sure it exists?")
		} else {
			session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Successfully changed language config file to %s", arguments[0]))
		}
	},
}
