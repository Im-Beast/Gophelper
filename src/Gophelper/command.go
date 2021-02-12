package gophelper

import (
	"time"
)

// Command basic structure of a command
type Command struct {
	ID string

	Name string

	Category *Category

	Aliases       []string
	CaseSensitive bool

	NeededPermissions int64

	NSFWOnly bool

	Description  string
	Usage        string
	UsageOnError bool

	RateLimit struct {
		Limit    int
		Duration time.Duration
	}

	LanguageSettings CommandConfig

	Handler func(*CommandContext)
}

// Function that is used to start command
func (cmd *Command) run(context *CommandContext) {
	for _, middleware := range context.Router.Middleware {
		ok, handler := middleware(context)

		if !ok && handler != nil {
			go handler(context)
			return
		}
	}

	cmd.Handler(context)
}
