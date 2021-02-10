package commands

import (
	"fmt"
	"time"

	gophelper "../Gophelper"
	"github.com/bwmarrin/discordgo"
)

//PingPongCommand Config structure
type PingPongCommand struct {
	TooManyMatchesMessage string `json:"tooManyMatches,omitempty"`

	Win struct {
		BotTrophyMessage  string `json:"botTrophyMessage"`
		UserTrophyMessage string `json:"userTrophyMessage"`
		Message           string `json:"message"`
	} `json:"win,omitempty"`

	ScoreboardMessage string `json:"scoreboardMessage,omitempty"`
}

//PingPong simple game
var PingPong = &gophelper.Command{
	ID: "PingPong",

	Name:    "🏓 Ping Pong",
	Aliases: []string{"game pingpong"},

	Description: "ping pong king kong",

	Usage: "game pingpong [_number of points to win] [_reaction ms]",

	RateLimit: gophelper.RateLimit{
		Limit:    1,
		Duration: time.Second * 5,
	},

	Handler: func(context *gophelper.CommandContext) {
		session := context.Session
		message := context.Event
		arguments := context.Arguments

		language := context.Command.LanguageSettings

		userScore := 0
		gopherScore := 0

		winPoints := 5
		pointLimit := 15
		delay := time.Millisecond * 1000

		if len(arguments) > 0 && gophelper.IsNumber(arguments[0]) {
			num := gophelper.StringToInt(arguments[0])
			if num <= pointLimit {
				winPoints = num
			} else {
				session.ChannelMessageSend(message.ChannelID, fmt.Sprintf(language.TooManyMatchesMessage, pointLimit))
				return
			}
		}

		if len(arguments) > 1 && gophelper.IsNumber(arguments[1]) {
			delay = time.Millisecond * time.Duration(gophelper.StringToInt(arguments[1]))
		}

		userName := message.Member.Nick
		if userName == "" {
			userName = message.Author.Username
		}

		getScore := func(msg string) string {
			return fmt.Sprintf("**🏓 %s** \n\t**· %s**: %d\n\t**· %s**: %d\n%s", language.ScoreboardMessage, userName, userScore, session.State.User.Username, gopherScore, msg)
		}

		getWinMessage := func() string {
			winner := "Gopher"
			trophyMessage := language.Win.BotTrophyMessage

			if userScore > gopherScore {
				winner = userName
				trophyMessage = language.Win.UserTrophyMessage
			}

			return fmt.Sprintf("**🏓 %s %s**\n\t%s\n\n%s", winner, language.Win.Message, trophyMessage, getScore(""))

		}

		gameMessage, _ := session.ChannelMessageSend(message.ChannelID, getScore("Ping!"))
		session.MessageReactionAdd(gameMessage.ChannelID, gameMessage.ID, "🏓")

		closeReactionHandler := session.AddHandler(func(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
			if event.UserID != session.State.User.ID {
				session.MessageReactionRemove(event.ChannelID, event.MessageID, event.Emoji.ID, event.UserID)
			}

			if event.MessageID != gameMessage.ID || event.UserID != message.Author.ID || event.Emoji.Name != "🏓" {
				return
			}

			userScore++
		})

		ticker := time.NewTicker(time.Millisecond * 250)

		for range ticker.C {
			oldUserScore := userScore

			if userScore >= winPoints || gopherScore >= winPoints {
				closeReactionHandler()
				session.ChannelMessageEdit(gameMessage.ChannelID, gameMessage.ID, getWinMessage())
				ticker.Stop()
				break
			}

			time.Sleep(delay)

			if userScore == oldUserScore {
				gopherScore++
			}

			session.MessageReactionRemove(gameMessage.ChannelID, gameMessage.ID, "🏓", message.Author.ID)
			session.ChannelMessageEdit(gameMessage.ChannelID, gameMessage.ID, getScore(""))
		}
	},
}
