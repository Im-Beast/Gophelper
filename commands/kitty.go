package commands

import (
	"time"

	generics "github.com/Im-Beast/Gophelper/commands/generics"
	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"
)

var kitties = []string{"http://www.randomkittengenerator.com/cats/rotator.php"}

// Kitty kitties
var Kitty = &gophelper.Command{
	ID: "Kitty",

	Name:    "🐈 Kitty",
	Aliases: []string{"kitty", "kittie", "cat"},

	Category: gophelper.CATEGORY_FUN,

	Description: "random pics of cute kitties",

	RateLimit: middleware.RateLimit{
		Limit:    1,
		Duration: time.Second * 5,
	},

	Handler: func(context *gophelper.CommandContext) {
		generics.ImageResponseHandler(context, "Kitty", kitties)
	},
}
