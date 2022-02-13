package gophelper

import (
	"strings"

	"github.com/Im-Beast/Gophelper/utils"
	"github.com/bwmarrin/discordgo"
)

// Router allows easier command handling by routing them from one place,
// it's even possible to use multiple routers e.g. for different prefixes/languages
type Router struct {
	Prefixes      []string
	CaseSensitive bool

	Config *Config

	Middleware    map[string][]Middleware
	CmdMiddleware []Middleware

	Commands   []*Command
	Categories []*Category
}

// Adds command to router
func (router *Router) AddCommand(command *Command) {
	router.RefreshCommand(command)

	if len(router.Middleware["AddCommand"]) > 0 {
		for _, middleware := range router.Middleware["AddCommand"] {
			middleware(&CommandContext{Command: command, Router: router})
		}
	}

	router.Commands = append(router.Commands, command)
}

// Returns function that removes command from maps (Commands and Middleware)
func (router *Router) RemoveCommand(command *Command) {
	if len(router.Middleware["RemoveCommand"]) > 0 {
		for _, middleware := range router.Middleware["RemoveCommand"] {
			middleware(&CommandContext{Command: command, Router: router})
		}
	}

	for i, cmd := range router.Commands {
		if cmd != command {
			continue
		}

		router.Commands = append(router.Commands[:i], router.Commands[i+1:]...)
		break
	}
}

// Refreshes all categories
func (router *Router) RefreshCategories() {
	for _, category := range router.Categories {
		category.Description = router.Config.Language.Categories[category.Name].Description
	}
}

// Refreshes given command
func (router *Router) RefreshCommand(command *Command) {
	if len(router.Middleware["RefreshCommand"]) > 0 {
		for _, middleware := range router.Middleware["RefreshCommand"] {
			middleware(&CommandContext{Command: command, Router: router})
		}
	}

	cmdLanguage := router.Config.Language.Commands[command.ID]
	command.Description = utils.GetStringVal(cmdLanguage.Description, command.Description)
}

// Calls RefreshCommand on all commands in router
func (router *Router) RefreshCommands() {
	for _, command := range router.Commands {
		router.RefreshCommand(command)
	}
}

// Adds middleware which fires when command executes in a router
func (router *Router) AddMiddleware(middleware Middleware) {
	router.CmdMiddleware = append(router.CmdMiddleware, middleware)
}

// Adds middleware to router which fires on specified event
func (router *Router) AddRouterMiddleware(event string, middleware Middleware) {
	router.Middleware[event] = append(router.Middleware[event], middleware)
}

// Initializes router
func (router *Router) Init(session *discordgo.Session) {
	session.AddHandler(router.handler)
}

// Routes commands
func (router *Router) handler(session *discordgo.Session, event *discordgo.MessageCreate) {
	msg := event.Message

	if msg.Author.Bot {
		return
	}

	foundPrefix := ""

	for _, prefix := range router.Prefixes {
		if utils.MatchesPrefix(msg.Content, prefix, router.CaseSensitive) {
			foundPrefix = prefix
			break
		}
	}

	if foundPrefix == "" {
		return
	}

	args := strings.Fields(msg.Content)
	args = args[len(strings.Fields(foundPrefix)):]

	strArgs := strings.Join(args, " ")

	var command *Command

	for _, cmd := range router.Commands {
		for _, alias := range cmd.Aliases {
			argumentString := strArgs[:utils.ClampInt(len(alias)+1, 0, len(strArgs))]

			if utils.StringMatches(argumentString, alias, cmd.CaseSensitive) || utils.StringMatches(argumentString, alias+" ", cmd.CaseSensitive) {
				args = args[len(strings.Fields(alias)):]
				command = cmd
				break
			}
		}

		if command != nil {
			context := CommandContext{
				Session: session,

				Command:   cmd,
				Arguments: args,

				Router: router,
				Event:  event,
			}

			go command.run(&context)
			break
		}
	}
}
