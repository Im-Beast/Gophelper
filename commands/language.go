package commands

import (
	"fmt"
	"log"
	"time"

	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"
	"github.com/bwmarrin/discordgo"
)

// LanguageSwitcher command, this command i keep in english and english only for understandable reasons
var LanguageSwitcher = &gophelper.Command{
	ID: "Language",

	Name:    "👅 Language",
	Aliases: []string{"lang", "language"},

	Category: gophelper.CATEGORY_CONFIG,

	NeededPermissions: discordgo.PermissionAdministrator,

	Description: "Change router language config on the fly (Admin only)",

	RateLimit: middleware.RateLimit{
		Limit:    1,
		Duration: time.Second * 30,
	},

	Handler: func(context *gophelper.CommandContext) {
		args := context.Arguments
		event := context.Event
		session := context.Session

		if len(args) == 0 {
			_, err := session.ChannelMessageSend(event.ChannelID, "You need to specify language config file")
			if err != nil {
				log.Printf("Command Language errored while sending error message: %s", err.Error())
			}
			return
		}

		config := context.Router.Config

		lang := args[0]
		err := config.LoadLanguage("configs/languages/" + lang)

		context.Router.RefreshCommands()
		context.Router.RefreshCategories()

		if err != nil {
			log.Printf("Failed loading %s language from configs/languages/%s", lang, lang)
			_, err = session.ChannelMessageSend(event.ChannelID, "Something happened while loading this config file, are you sure it exists?")
		} else {
			_, err = session.ChannelMessageSend(event.ChannelID, fmt.Sprintf("Successfully changed language config file to %s", lang))
		}

		if err != nil {
			log.Printf("Command Language errored while sending error message: %s", err.Error())
		}
	},
}
