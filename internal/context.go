package gophelper

import "github.com/bwmarrin/discordgo"

// CommandContext contains necessary information about command and from where it comes from
type CommandContext struct {
	Session *discordgo.Session
	Router  *Router

	Command   *Command
	Arguments []string

	Event *discordgo.MessageCreate
}
