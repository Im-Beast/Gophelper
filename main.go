package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Im-Beast/Gophelper/commands"

	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"

	"github.com/bwmarrin/discordgo"
)

var (
	router = gophelper.Router{
		Prefixes:      []string{"go"},
		CaseSensitive: false,
		Config:        gophelper.LoadConfig("configs/bot.json", "configs/languages/english.json"),
	}
)

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
		<-sc

		fmt.Println("\rDisabling bot")

		discord.Close()
	}()

	if err != nil {
		fmt.Println("Some error happened while launching bot:", err)
		return
	}

	err = discord.Open()

	if err != nil {
		fmt.Println("Failed opening connection", err)
		return
	}

	fmt.Printf("Bot %s is up and running\n", discord.State.User.Username)

	registerCategories()
	registerMiddleware()
	registerCommands()

	router.Init(discord)
}

func registerCategories() {
	router.Categories = []*gophelper.Category{
		gophelper.CATEGORY_FUN,
		gophelper.CATEGORY_MISC,
		gophelper.CATEGORY_MOD,
		gophelper.CATEGORY_GAMES,
		gophelper.CATEGORY_CONFIG,
	}

	router.RefreshCategories()
}

func registerMiddleware() {
	router.AddRouterMiddleware("RefreshCommand", middleware.HelpInitMiddlware)

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
	router.AddCommand(commands.Hentai)
	router.AddCommand(commands.Doggie)
	router.AddCommand(commands.PingPong)
	router.AddCommand(commands.EightBall)
	router.AddCommand(commands.LanguageSwitcher)
}
