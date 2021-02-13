package commands

import (
	"fmt"
	"time"

	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"
	"github.com/bwmarrin/discordgo"
)

// LanguageSwitcher command, this command i keep in english and english only for understandable reasons
var LanguageSwitcher = &gophelper.Command{
	Name:    "👅 Language",
	Aliases: []string{"lang"},

	Category: gophelper.CATEGORY_CONFIG,

	NeededPermissions: discordgo.PermissionAdministrator,

	Description: "Change router language config on the fly (Admin only)",

	RateLimit: middleware.RateLimit{
		Limit:    1,
		Duration: time.Second * 30,
	},

	Handler: func(context *gophelper.CommandContext) {
		arguments := context.Arguments
		session := context.Session
		message := context.Event

		if len(arguments) == 0 {
			_, err := session.ChannelMessageSend(message.ChannelID, "You need to specify language config file")
			if err != nil {
				fmt.Println("Error on language command when sending message")
			}
			return
		}

		config := context.Router.Config

		err := config.LoadLanguage("configs/languages/" + arguments[0])
		context.Router.RefreshCommands()
		context.Router.RefreshCategories()

		if err != nil {
			_, err = session.ChannelMessageSend(message.ChannelID, "Something happened while loading this config file, are you sure it exists?")
			fmt.Println("Failed on language command when loading language file")
		} else {
			_, err = session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Successfully changed language config file to %s", arguments[0]))
		}

		if err != nil {
			fmt.Println("Error on language command when sending message")
		}
	},
}
