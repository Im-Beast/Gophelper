package gophelper

// Middleware handler used before running command
type Middleware func(*CommandContext) (bool, func(*CommandContext))
