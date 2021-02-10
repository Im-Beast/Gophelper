package commands

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	gophelper "../Gophelper"
	"github.com/bwmarrin/discordgo"
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

// Pet pet
var Pet = &gophelper.Command{
	ID: "Pet",

	Name:    "✋ Pet",
	Aliases: []string{"pet"},

	Description: "Pet someone or get pet",

	RateLimit: gophelper.RateLimit{
		Limit:    1,
		Duration: time.Second * 5,
	},

	Handler: func(context *gophelper.CommandContext) {
		arguments := context.Arguments
		message := context.Event.Message
		session := context.Session

		routerLanguage := context.Router.Config.Language
		language := context.Command.LanguageSettings

		isTag := false
		userID := message.Author.ID

		if len(arguments) > 0 {
			isTag = true
			userID = arguments[0]

			if gophelper.IsMention(userID) {
				userID = gophelper.MentionToID(userID)
			}
		}

		title := language.Response.NonMention

		if isTag {
			member, err := session.GuildMember(message.GuildID, userID)

			nick := message.Author.Username

			if err == nil {
				nick = member.User.Username
			}

			title = fmt.Sprintf(language.Response.Mention, nick)
		}

		index := rand.Intn(len(pets))

		url := pets[index]

		if strings.Contains(url, "?") {
			url = fmt.Sprintf("%s&nocache=%d", url, time.Now().UnixNano())
		} else {
			url = fmt.Sprintf("%s?nocache=%d", url, time.Now().UnixNano())
		}

		embed := &discordgo.MessageEmbed{
			Title: title,
			Color: 0xFFbbbb,
			Provider: &discordgo.MessageEmbedProvider{
				URL: url,
			},
			Image: &discordgo.MessageEmbedImage{
				URL: url,
			},
		}

		message, err := session.ChannelMessageSendEmbed(message.ChannelID, embed)

		if err != nil {
			session.ChannelMessageSend(message.ChannelID, routerLanguage.MessageSendError)
		} else {
			session.MessageReactionAdd(message.ChannelID, message.ID, "✋")
		}
	},
}
