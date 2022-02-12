package commands

import (
	"time"

	generics "github.com/Im-Beast/Gophelper/commands/generics"
	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"
)

var kisses = []string{
	"https://pa1.narvii.com/5823/f10cce909b5bfa6f05f0af496558a16ed4840c06_hq.gif",
	"https://media1.tenor.com/images/78095c007974aceb72b91aeb7ee54a71/tenor.gif",
	"https://media.giphy.com/media/G3va31oEEnIkM/giphy.gif",
	"https://media.giphy.com/media/y0H514IGMusQE/giphy.gif",
	"https://media1.tenor.com/images/d307db89f181813e0d05937b5feb4254/tenor.gif",
	"https://data.whicdn.com/images/166496706/original.gif",
}

// Kiss :*
var Kiss = &gophelper.Command{
	ID: "Kiss",

	Name:    "😘 Kiss",
	Aliases: []string{"kiss"},

	Category: gophelper.CATEGORY_FUN,

	Description: "Kiss someone or get kissed :*",

	RateLimit: middleware.RateLimit{
		Limit:    1,
		Duration: time.Second * 5,
	},

	Handler: func(context *gophelper.CommandContext) {
		generics.ImageResponseHandler(context, "Kiss", kisses)
	},
}
