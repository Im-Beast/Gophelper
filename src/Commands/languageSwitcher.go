package commands

import (
	"fmt"
	"time"

	gophelper "../Gophelper"
)

// LanguageSwitcher command, this command i keep in english and english only for understandable reasons
var LanguageSwitcher = &gophelper.Command{
	Name:    "👅 Language",
	Aliases: []string{"lang"},

	Description: "Change router language config on the fly (Admin only)",

	RateLimit: gophelper.RateLimit{
		Limit:    1,
		Duration: time.Second * 1,
	},

	Handler: func(context *gophelper.CommandContext) {
		arguments := context.Arguments
		session := context.Session
		message := context.Event

		roles, _ := session.GuildRoles(message.GuildID)

		isAdmin := false

		for _, role := range message.Member.Roles {
			for _, drole := range roles {
				if drole.ID != role {
					continue
				}

				if drole.Permissions&0x8 != 0 {
					isAdmin = true
					break
				}
			}

			if isAdmin {
				break
			}
		}

		if isAdmin == false {
			session.ChannelMessageSend(message.ChannelID, "You need to be admin to use this command")
			return
		}

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
