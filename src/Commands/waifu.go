package commands

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	gophelper "../Gophelper"
	middleware "../Middleware"
	"github.com/bwmarrin/discordgo"
)

var waifus = []string{
	"https://images-ext-2.discordapp.net/external/3xtOnwJCdo4KBurSga7nkBku9JABMrI8Ft_0zrT7dZw/https/img.devilchan.com/original/8e/48/elaina-3047-2000x3430.jpeg",
	"https://i.pinimg.com/originals/59/e4/aa/59e4aa928860dd9cdb93ab6670484027.png",
	"https://i.pinimg.com/originals/63/1d/4d/631d4d0fe45927af4fcca400e756d603.jpg",
	"https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/0eb1ac2f-1a1a-467f-b2c7-69f9eaca9700/dbic6fb-757be321-dc98-46b7-8a5c-cecaf9fc5098.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOiIsImlzcyI6InVybjphcHA6Iiwib2JqIjpbW3sicGF0aCI6IlwvZlwvMGViMWFjMmYtMWExYS00NjdmLWIyYzctNjlmOWVhY2E5NzAwXC9kYmljNmZiLTc1N2JlMzIxLWRjOTgtNDZiNy04YTVjLWNlY2FmOWZjNTA5OC5wbmcifV1dLCJhdWQiOlsidXJuOnNlcnZpY2U6ZmlsZS5kb3dubG9hZCJdfQ.baPAJO1cIznJpG0rcfCfVDBV1dQc-CZumwk9I7MS9tk",
	"https://i.pinimg.com/originals/31/ff/fe/31fffe03a58148ab6122da4ac90a45a8.jpg",
	"https://data.whicdn.com/images/328607597/original.png",
	"https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/a6ce62a7-bf87-485f-a844-cc45c137d88b/dc9ard0-f5d17c05-b030-4aed-b075-b6c0e604f1a9.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOiIsImlzcyI6InVybjphcHA6Iiwib2JqIjpbW3sicGF0aCI6IlwvZlwvYTZjZTYyYTctYmY4Ny00ODVmLWE4NDQtY2M0NWMxMzdkODhiXC9kYzlhcmQwLWY1ZDE3YzA1LWIwMzAtNGFlZC1iMDc1LWI2YzBlNjA0ZjFhOS5wbmcifV1dLCJhdWQiOlsidXJuOnNlcnZpY2U6ZmlsZS5kb3dubG9hZCJdfQ.QyenHuPwniDIdmb8I1OvWe6bohWgvPqdGBxyICTt4Gs",
	"https://i.pinimg.com/originals/43/ca/5a/43ca5a4a5b896deb2ddfb05472f5e88a.jpg",
	"https://i.pinimg.com/originals/41/65/e4/4165e4667584f379246d20e151c3dc39.png",
	"https://i.pinimg.com/originals/0e/90/19/0e901958b069789a586c51016237bbfc.jpg",
	"https://files.wallpaperpass.com/2019/10/violet%20evergarden%20wallpaper%20165%20-%201440x2560-768x1365.jpg",
	"https://i.pinimg.com/originals/b7/25/12/b725125aaebafbcbf2fb3886a55d2d6f.jpg",
	"https://i.pinimg.com/originals/43/f6/e8/43f6e880af7709390ad2408ebf7d0aaf.jpg",
	"https://static.wikia.nocookie.net/akamegakill/images/9/93/Chelsea_.png/revision/latest/scale-to-width-down/340?cb=20140926185030",
	"https://i.pinimg.com/originals/05/fa/66/05fa6636db4f896bf2984460adfd6eda.jpg",
	"https://i.pinimg.com/originals/89/28/ce/8928ce0d8699ca29878977787c7227b1.jpg",
	"https://lh4.googleusercontent.com/proxy/r4KuZtZ2iL-KQDcI5-N1XeYP6BGbQBIpzDpLl2VbyRxLBxUr1OrhyEgI-CJMTrUPF2BuSZBt4rzvRE3KebuvP2WlVA7Emy2wvzqW0AGCvJZ914qDjv3Rfhoo6KsrQvLM6K53BhTuxzQdsFHCELNYA=s0-d",
	"https://i.pinimg.com/originals/c6/56/2f/c6562f261f9cf80cddbfe8ad19220664.jpg",
	"https://i.pinimg.com/originals/37/48/c1/3748c19400b766946b2b9cb2cba7b827.jpg",
	"https://web.johnson-local.com/wp-content/uploads/2015/02/4575109_20140320093429.jpg",
	"https://themepack.me/i/c/749x468/media/g/632/no-game-no-life-theme-go14.jpg",
	"https://c.wallhere.com/photos/e8/a4/anime_girls_Majo_no_Tabitabi_Elaina_Majo_no_Tabitabi_Azuuru_witch-1947219.jpg",
	"https://i.pinimg.com/originals/64/c7/4e/64c74e3bec506de7e8c5b159f3dfff0a.jpg",
	"https://i.pinimg.com/736x/e1/81/de/e181de626c17fe9f6cc9f081e489f582.jpg",
	"https://i.pinimg.com/originals/d1/02/f4/d102f497adc351ad6a49caf02e8aa4f1.jpg",
	"https://c.wallhere.com/photos/d8/93/anime_girls_Re_Zero_Kara_Hajimeru_Isekai_Seikatsu_Echidna_Re_Zero_Kara_Hajimeru_Isekai_Seikatsu_Hanakanzarashi-1946225.jpg",
	"https://c4.wallpaperflare.com/wallpaper/772/24/887/rokudenashi-majutsu-koushi-to-akashic-records-rumia-anime-girls-miniskirt-wallpaper-preview.jpg",
}

// Waifu waifu
var Waifu = &gophelper.Command{
	ID: "Waifu",

	Name:    "🌸 Waifu",
	Aliases: []string{"waifu"},

	Category: gophelper.CATEGORY_FUN,

	Description: "Some weird weeb shit",

	RateLimit: middleware.RateLimit{
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

		index := rand.Intn(len(waifus))

		url := waifus[index]

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
			_, err = session.ChannelMessageSend(message.ChannelID, routerLanguage.Errors.MessageSend)
		} else {
			err = session.MessageReactionAdd(message.ChannelID, message.ID, "🌸")
		}

		if err != nil {
			fmt.Println("Error on waifu command when reacting/sending message")
		}
	},
}
