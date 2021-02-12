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

	Middleware []Middleware

	Config *Config

	Commands []*Command

	Categories []*Category
}

var routerMiddleware = make(map[string][]Middleware)

// AddCmd Adds command to router
func (router *Router) AddCommand(command *Command) {
	router.RefreshCommand(command)

	if len(routerMiddleware["AddCommand"]) > 0 {
		for _, middleware := range routerMiddleware["AddCommand"] {
			middleware(&CommandContext{Command: command, Router: router})
		}
	}

	router.Commands = append(router.Commands, command)
}

// RemoveCmd returns function that removes command from maps
func (router *Router) RemoveCommand(command *Command) {
	if len(routerMiddleware["RemoveCommand"]) > 0 {
		for _, middleware := range routerMiddleware["RemoveCommand"] {
			middleware(&CommandContext{Command: command, Router: router})
		}
	}

	for i, cmd := range router.Commands {
		if cmd == command {
			router.Commands = append(router.Commands[:i], router.Commands[i+1:]...)
			return
		}
	}
}

// RefreshCategories refreshes all categories
func (router *Router) RefreshCategories() {
	for _, category := range router.Categories {
		category.Description = router.Config.Language.Categories[category.Name].Description
	}
}

// RefreshCommand refreshes given command
func (router *Router) RefreshCommand(command *Command) {
	if len(routerMiddleware["RefreshCommand"]) > 0 {
		for _, middleware := range routerMiddleware["RefreshCommand"] {
			middleware(&CommandContext{Command: command, Router: router})
		}
	}

	commandLanguageSettings := router.Config.Language.Commands[command.ID]

	command.LanguageSettings = commandLanguageSettings
	command.Description = commandLanguageSettings.Description
}

// RefreshCommands does RefreshCommand to all commands in router
func (router *Router) RefreshCommands() {
	for _, command := range router.Commands {
		router.RefreshCommand(command)
	}
}

// AddMiddleware Adds middleware when executing command to router
func (router *Router) AddMiddleware(middleware Middleware) {
	router.Middleware = append(router.Middleware, middleware)
}

// AddRouterMiddleware Adds middleware to router on specified event
func (router *Router) AddRouterMiddleware(event string, middleware Middleware) {
	routerMiddleware[event] = append(routerMiddleware[event], middleware)
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

			go command.run(&context)
			break
		}
	}
}
