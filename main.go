package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Im-Beast/Gophelper/commands"

	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"

	"github.com/bwmarrin/discordgo"
)

var router = gophelper.Router{
	Prefixes:      []string{"go"},
	CaseSensitive: false,
	Config:        gophelper.LoadConfig("configs/bot.json", "configs/languages/english.json"),
}

func main() {
	bot, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Failed launching bot: %s", err.Error())
	}

	bot.StateEnabled = true

	err = bot.Open()
	if err != nil {
		log.Fatalf("Failed opening bot connection: %s", err.Error())
	}

	log.Printf("Bot %s is up and running\n", bot.State.User.Username)

	registerMiddleware()
	registerCategories()
	registerCommands()

	router.Init(bot)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	log.Println("Closing bot")
	err = bot.Close()
	if err != nil {
		log.Printf("Failed closing bot: %s", err.Error())
	}
}

func registerCategories() {
	var sorted []*gophelper.Category

	for _, strCategory := range router.Config.Commands.Help.Categories {
		for _, category := range gophelper.CATEGORIES {
			if category.Name == strCategory {
				sorted = append(sorted, category)
			}
		}
	}

	router.Categories = sorted

	router.RefreshCategories()
}

func registerMiddleware() {
	router.AddMiddleware(middleware.RateLimiterMiddleware)
	router.AddMiddleware(middleware.PermissionCheckMiddleware)
}

func registerCommands() {
	router.AddCommand(commands.Pet)
	router.AddCommand(commands.Hug)
	router.AddCommand(commands.Ping)
	router.AddCommand(commands.Kiss)
	router.AddCommand(commands.Help)
	router.AddCommand(commands.Kitty)
	router.AddCommand(commands.Stats)
	router.AddCommand(commands.Waifu)
	router.AddCommand(commands.Doggie)
	router.AddCommand(commands.Pinterest)
	router.AddCommand(commands.EightBall)
	router.AddCommand(commands.LanguageSwitcher)
}
