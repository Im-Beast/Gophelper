package gophelper

import (
	"github.com/bwmarrin/discordgo"
)

// PermissionCheckMiddleware is middleware that confirms whether user that executed command actually has enough permissions to use it
func PermissionCheckMiddleware(context *CommandContext) (bool, func(*CommandContext)) {
	session := context.Session
	member := context.Event.Member
	message := context.Event
	command := context.Command

	routerLanguage := context.Router.Config.Language

	if command.NSFWOnly && isNSFW(session, message.ChannelID) {
		return false, func(*CommandContext) {
			session.ChannelMessageSend(message.ChannelID, routerLanguage.Errors.NSFWOnly)
		}
	}

	roles, _ := session.GuildRoles(message.GuildID)
	memberPermissions := GetMemberPermissions(roles, member)

	if command.NeededPermissions != 0 && memberPermissions&command.NeededPermissions == 0 {
		return false, func(*CommandContext) {
			session.ChannelMessageSend(message.ChannelID, routerLanguage.Errors.TooFewPermissions)
		}
	}

	return true, nil
}

// GetMemberPermissions Gets Member Permission
func GetMemberPermissions(roles []*discordgo.Role, member *discordgo.Member) int64 {
	var permissions int64

	for _, mRole := range member.Roles {
		for _, gRole := range roles {
			if gRole.ID != mRole {
				continue
			}

			permissions |= gRole.Permissions
		}
	}

	return permissions
}

func isNSFW(session *discordgo.Session, channelID string) bool {
	channel, err := session.Channel(channelID)
	if err != nil {
		return false
	}
	return channel.NSFW
}
