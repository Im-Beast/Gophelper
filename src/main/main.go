package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	commands "../Commands"
	gophelper "../Gophelper"

	"github.com/bwmarrin/discordgo"
)

var (
	router = gophelper.Router{
		Prefixes:      []string{"go"},
		CaseSensitive: false,
		Config:        gophelper.LoadConfig("config.json", "../Languages/english.json"),
	}
)

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	if err != nil {
		fmt.Println("Some error happened while launching bot:", err)
		return
	}

	err = discord.Open()

	registerMiddleware()
	registerCommands()

	router.Init(discord)

	if err != nil {
		fmt.Println("Failed opening connection", err)
		return
	}

	fmt.Printf("Bot %s is up and running\n", discord.State.User.Username)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func registerMiddleware() {
	router.AddMiddleware(gophelper.RateLimiterMiddleware)
}

func registerCommands() {
	router.AddCmd(commands.Ping)
	router.AddCmd(commands.Kitty)
	router.AddCmd(commands.Doggie)
	router.AddCmd(commands.Kiss)
	router.AddCmd(commands.Hug)
	router.AddCmd(commands.Help)
	router.AddCmd(commands.Waifu)
	router.AddCmd(commands.Hentai)
	router.AddCmd(commands.EightBall)
	router.AddCmd(commands.Pet)
	router.AddCmd(commands.PingPong)
	router.AddCmd(commands.Stats)
	router.AddCmd(commands.LanguageSwitcher)
}
