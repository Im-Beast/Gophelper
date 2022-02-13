package commands

import (
	"time"

	generics "github.com/Im-Beast/Gophelper/commands/generics"
	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"
)

var doggies = []string{"https://www.randomdoggiegenerator.com/randomdoggie.php"}

// Command which replies to you with cute doggie images 🐕
var Doggie = &gophelper.Command{
	ID: "Doggie",

	Name:    "🐕 Doggie",
	Aliases: []string{"doggie", "doggy", "dog"},

	Category: gophelper.CATEGORY_FUN,

	Description: "Random pics of cute doggies",

	RateLimit: middleware.RateLimit{
		Limit:    1,
		Duration: time.Second * 5,
	},

	Handler: func(context *gophelper.CommandContext) {
		generics.ImageResponseHandler(context, "Doggie", "🐕", doggies)
	},
}
