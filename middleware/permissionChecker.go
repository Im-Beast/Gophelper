package middleware

import (
	"fmt"

	gophelper "github.com/Im-Beast/Gophelper/internal"
	utils "github.com/Im-Beast/Gophelper/utils"
)

// PermissionCheckMiddleware is middleware that confirms whether user that executed command actually has enough permissions to use it
func PermissionCheckMiddleware(context *gophelper.CommandContext) (bool, func(*gophelper.CommandContext)) {
	session := context.Session
	message := context.Event
	command := context.Command

	routerLanguage := context.Router.Config.Language

	if command.NSFWOnly {
		if utils.IsNSFW(session, message.ChannelID) {
			return false, func(*gophelper.CommandContext) {
				_, err := session.ChannelMessageSend(message.ChannelID, routerLanguage.Errors.NSFWOnly)
				if err != nil {
					fmt.Println("Error while sending too few permissions message")
				}
			}
		}
	}

	if command.NeededPermissions == 0 {
		return true, nil
	}

	permissions, err := session.UserChannelPermissions(message.Author.ID, message.ChannelID)
	if err != nil {
		fmt.Println("Failed getting perms")
		permissions = 0
	}

	if permissions&command.NeededPermissions == 0 {
		return false, func(*gophelper.CommandContext) {
			_, err := session.ChannelMessageSend(message.ChannelID, routerLanguage.Errors.TooFewPermissions)
			if err != nil {
				fmt.Println("Error while sending too few permissions message")
			}
		}
	}

	return true, nil
}
