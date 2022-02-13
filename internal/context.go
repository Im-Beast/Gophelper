package gophelper

import "github.com/bwmarrin/discordgo"

// Contains necessary information about command
type CommandContext struct {
	Session *discordgo.Session
	Router  *Router

	Command   *Command
	Arguments []string

	Event *discordgo.MessageCreate
}
