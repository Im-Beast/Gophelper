package middleware

import (
	"log"

	gophelper "github.com/Im-Beast/Gophelper/internal"
	utils "github.com/Im-Beast/Gophelper/utils"
)

// Confirms whether user that wanted to execute command has enough permissions to use it
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
					log.Printf("Middleware PermissionChecker errored while sending NSFWOnly error message: %s\n", err.Error())
				}
			}
		}
	}

	if command.NeededPermissions == 0 {
		return true, nil
	}

	permissions, err := session.UserChannelPermissions(message.Author.ID, message.ChannelID)
	if err != nil {
		log.Printf("Middleware PermissionChecker errored while getting permissions: %s\n", err.Error())
		permissions = 0
	}

	if permissions&command.NeededPermissions == 0 {
		return false, func(*gophelper.CommandContext) {
			_, err := session.ChannelMessageSend(message.ChannelID, routerLanguage.Errors.TooFewPermissions)
			if err != nil {
				log.Printf("Middleware PermissionChecker errored while sending TooFewPermissions error message: %s\n", err.Error())
			}
		}
	}

	return true, nil
}
