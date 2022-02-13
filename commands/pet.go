package commands

import (
	"time"

	generics "github.com/Im-Beast/Gophelper/commands/generics"
	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"
)

var pets = []string{
	"https://images-ext-2.discordapp.net/external/j9sTeJILNLFQAaY-pD_bSiym-qnnTudFfS7FLUKZY8g/https/gifimage.net/wp-content/uploads/2018/11/pet-gif-anim%25C3%25A9-1.gif",
	"https://images-ext-1.discordapp.net/external/XLglLO52mhgjtw0THbcqs2WCg_NwLlgqA3cx2x27nZs/https/cdn.pixilart.com/photos/orginal/c85fbace7380095.gif",
	"https://images-ext-2.discordapp.net/external/LezCAlUBLBvvW4tsL756v-FjQNVhgcHy0WSJXYrumHA/http/cdn.lowgif.com/medium/b5d3a6ed359bd6e1-.gif",
	"https://images-ext-2.discordapp.net/external/S68QdCr9DmFuWUBcUUuyeaBCpMZwMSAjvAuDSQsS6hc/%3Fitemid%3D11118254/https/media1.tenor.com/images/1a8e560e8873ce2a48b3dfbbaa7805ec/tenor.gif",
	"https://imgur.com/3rn2JHt.gif",
	"https://images-ext-1.discordapp.net/external/9_QHBJturF7UqxiFyN7QhlOpk5oW1IKJT8ii_dLZTNM/https/i.chzbgr.com/full/8155571968/hEE6AB87F/petting-the-kitty",
	"https://images-ext-1.discordapp.net/external/enWfU6I6iJG2ZBfjVnnwoX0CYypRyYQ7smjB75Od9NY/%3Fitemid%3D13236885/https/media1.tenor.com/images/9bf3e710f33cae1eed1962e7520f9cf3/tenor.gif",
	"https://images-ext-2.discordapp.net/external/YGfES-ulmqGYpRHtQ2CPkNNb8v-czYIl6BCysdhfXyQ/%3Fitemid%3D5359308/https/media1.tenor.com/images/a4a2b1eaa47fd0d8d0951433bc59ab9a/tenor.gif",
}

// Command which replies to you with images of petting ✋
var Pet = &gophelper.Command{
	ID: "Pet",

	Category: gophelper.CATEGORY_FUN,

	Name:    "✋ Pet",
	Aliases: []string{"pet"},

	Description: "Pet someone or get pet",

	RateLimit: middleware.RateLimit{
		Limit:    1,
		Duration: time.Second * 5,
	},

	Handler: func(context *gophelper.CommandContext) {
		generics.ImageResponseHandler(context, "Pet", "✋", pets)
	},
}
