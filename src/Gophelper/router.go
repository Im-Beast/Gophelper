package gophelper

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Router allows easier command handling by routing them from one place,
// it's even possible to use multiple routers e.g. for different prefixes/languages
type Router struct {
	Prefixes      []string
	CaseSensitive bool

	Middleware []func(*CommandContext) (bool, func(*CommandContext))

	Config *Config

	Commands []*Command
}

// AddCmd Adds command to router
func (router *Router) AddCmd(cmd *Command) {
	for _, langCmd := range router.Config.Language.Commands {
		if langCmd.ID == cmd.ID {
			cmd.LanguageSettings = langCmd
			cmd.Description = langCmd.Description
		}
	}

	router.Commands = append(router.Commands, cmd)
}

// RefreshCommands refreshes commands data
func (router *Router) RefreshCommands() {
	for _, cmd := range router.Commands {
		for _, langCmd := range router.Config.Language.Commands {
			if langCmd.ID == cmd.ID {
				cmd.LanguageSettings = langCmd
				cmd.Description = langCmd.Description
			}
		}
	}
}

// AddMiddleware Adds middleware to router
func (router *Router) AddMiddleware(middleware func(*CommandContext) (bool, func(*CommandContext))) {
	router.Middleware = append(router.Middleware, middleware)
}

// Init Initializes router
func (router *Router) Init(session *discordgo.Session) {
	session.AddHandler(router.handler)
}

// handler routes commands
func (router *Router) handler(session *discordgo.Session, event *discordgo.MessageCreate) {
	message := event.Message

	if message.Author.Bot {
		return
	}

	Prefix := ""

	for _, prefix := range router.Prefixes {
		if MatchesPrefix(message.Content, prefix, router.CaseSensitive) {
			Prefix = prefix
			break
		}
	}

	if Prefix == "" {
		return
	}

	arguments := strings.Fields(message.Content)
	arguments = arguments[len(strings.Fields(Prefix)):]

	argumentsString := strings.Join(arguments, " ")

	var command *Command

	for _, cmd := range router.Commands {
		for _, alias := range cmd.Aliases {
			argumentString := argumentsString[:ClampInt(len(alias)+1, 0, len(argumentsString))]

			if Matches(argumentString, alias, cmd.CaseSensitive) || Matches(argumentString, alias+" ", cmd.CaseSensitive) {
				arguments = arguments[len(strings.Fields(alias)):]
				command = cmd
				break
			}
		}

		if command != nil {
			context := CommandContext{
				Session: session,

				Command:   cmd,
				Arguments: arguments,

				Router: router,
				Event:  event,
			}

			command.run(&context)
			break
		}
	}
}
