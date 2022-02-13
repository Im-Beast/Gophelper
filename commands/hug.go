package commands

import (
	"time"

	generics "github.com/Im-Beast/Gophelper/commands/generics"
	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"
)

var hugs = []string{
	"https://media1.tenor.com/images/074d69c5afcc89f3f879ca473e003af2/tenor.gif?itemid=4898650",
	"https://media1.tenor.com/images/108c2257683620292f4687262f26e872/tenor.gif?itemid=17258498",
	"https://media1.tenor.com/images/8ac5ada8524d767b77d3d54239773e48/tenor.gif?itemid=16334628",
	"https://i.imgur.com/rioNdmc.gif",
	"https://i.gifer.com/2QEa.gif",
	"https://acegif.com/wp-content/uploads/anime-hug.gif",
	"https://media.tenor.com/images/b6d0903e0d54e05bb993f2eb78b39778/tenor.gif",
	"https://data.whicdn.com/images/219995514/original.gif",
	"https://images-ext-2.discordapp.net/external/rxstxw_1DcDfXP2ZTHcq5Fk4um5Q57mXsF7Klwyz6Q4/https/data.whicdn.com/images/334957046/original.gif",
	"https://media.tenor.co/videos/161dd2416944be9249bcd0a6d69e0463/mp4",
}

// Command which replies to you with hugs 🤗 and love!
var Hug = &gophelper.Command{
	ID: "Hug",

	Name:    "🤗 Hug",
	Aliases: []string{"hug"},

	Category: gophelper.CATEGORY_FUN,

	Description: "Hug someone or get hugged	",

	RateLimit: middleware.RateLimit{
		Limit:    1,
		Duration: time.Second * 5,
	},

	Handler: func(context *gophelper.CommandContext) {
		generics.ImageResponseHandler(context, "Hug", "🤗", hugs)
	},
}
