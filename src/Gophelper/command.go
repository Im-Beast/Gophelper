package gophelper

// Command basic structure of a command
type Command struct {
	ID string

	Name          string
	Aliases       []string
	CaseSensitive bool

	Description  string
	Usage        string
	UsageOnError bool

	RateLimit RateLimit

	LanguageSettings CommandConfig

	Handler func(*CommandContext)
}

// Function that is used to start command
func (cmd *Command) run(context *CommandContext) {
	for _, middleware := range context.Router.Middleware {
		ok, handler := middleware(context)

		if !ok {
			if handler != nil {
				handler(context)
			}
			return
		}
	}

	cmd.Handler(context)
}
