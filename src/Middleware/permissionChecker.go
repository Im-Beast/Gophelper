package middleware

import (
	gophelper "../Gophelper"
)

// PermissionCheckMiddleware is middleware that confirms whether user that executed command actually has enough permissions to use it
func PermissionCheckMiddleware(context *gophelper.CommandContext) (bool, func(*gophelper.CommandContext)) {
	session := context.Session
	member := context.Event.Member
	message := context.Event
	command := context.Command

	routerLanguage := context.Router.Config.Language

	if command.NSFWOnly {
		if gophelper.IsNSFW(session, message.ChannelID) {
			return false, func(*gophelper.CommandContext) {
				session.ChannelMessageSend(message.ChannelID, routerLanguage.Errors.NSFWOnly)
			}
		}

	}

	if command.NeededPermissions == 0 {
		return true, nil
	}

	roles, _ := session.GuildRoles(message.GuildID)
	memberPermissions := gophelper.GetMemberPermissions(roles, member)

	if memberPermissions&command.NeededPermissions == 0 {
		return false, func(*gophelper.CommandContext) {
			session.ChannelMessageSend(message.ChannelID, routerLanguage.Errors.TooFewPermissions)
		}
	}

	return true, nil
}
