package gophelper

import (
	"time"
)

// Basic structure of a command
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

	Handler func(*CommandContext)
}

// Runs command with given context
func (cmd *Command) run(context *CommandContext) {
	for _, middleware := range context.Router.CmdMiddleware {
		ok, handler := middleware(context)

		if !ok && handler != nil {
			go handler(context)
			return
		}
	}

	cmd.Handler(context)
}
