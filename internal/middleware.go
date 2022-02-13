package gophelper

// Function called before running command handler
type Middleware func(*CommandContext) (bool, func(*CommandContext))
