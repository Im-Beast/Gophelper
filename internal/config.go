package gophelper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type (
	// StatsCommand structure for stats command
	StatsCommand struct {
		UserNotFound string `json:"userNotFound,omitempty"`
		Title        string `json:"title,omitempty"`
		CreationDate string `json:"creationDate,omitempty"`
		JoinDate     string `json:"joinDate,omitempty"`
	}

	// PingPongCommand structure for game pingpong command
	PingPongCommand struct {
		TooManyMatchesMessage string `json:"tooManyMatches,omitempty"`

		Win struct {
			BotTrophyMessage  string `json:"botTrophyMessage"`
			UserTrophyMessage string `json:"userTrophyMessage"`
			Message           string `json:"message"`
		} `json:"win,omitempty"`

		ScoreboardMessage string `json:"scoreboardMessage,omitempty"`
	}

	// EightBallCommand structure for 8ball command
	EightBallCommand struct {
		NoArgumentsMessage string   `json:"noArgumensMessage,omitempty"`
		Answers            []string `json:"answers,omitempty"`
	}

	// HelpCommand structure for help command
	HelpCommand struct {
		Embed struct {
			Main struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			} `json:"main"`

			CategoryTitle string `json:"categoryTitle"`

			Name   string `json:"name"`
			NoName string `json:"noName"`

			Description   string `json:"description"`
			NoDescription string `json:"noDescription"`

			Usage string `json:"usage"`

			Aliases   string `json:"aliases"`
			NoAliases string `json:"noAliases"`
		} `json:"embed,omitempty"`

		Page string `json:"page,omitempty"`

		CommandNotFound struct {
			Title   string `json:"title"`
			Message string `json:"message"`
		} `json:"commandNotFound,omitempty"`
	}

	// GiveCommand structure for kitty,doggy etc commands
	GiveCommand struct {
		Response struct {
			Mention    string `json:"mention"`
			NonMention string `json:"nonMention"`
		} `json:"response,omitempty"`
	}

	// CommandConfig contains structure that commands used for their translation
	CommandConfig struct {
		Description string `json:"description,omitempty"`
		HelpCommand
		GiveCommand
		EightBallCommand
		PingPongCommand
		StatsCommand
	}

	// CategoryConfig contains structure that categories used for their translation
	CategoryConfig struct {
		Description string `json:"description"`
	}

	// LanguageConfig contains config information that is mainly about translation
	LanguageConfig struct {
		Errors struct {
			MessageSend       string `json:"messageSend"`
			TooFewPermissions string `json:"tooFewPermissions"`
			NSFWOnly          string `json:"nsfwOnly"`
			RateLimit         string `json:"rateLimit"`
		} `json:"errors"`

		FunFacts   []string                  `json:"funFacts"`
		Commands   map[string]CommandConfig  `json:"commands"`
		Categories map[string]CategoryConfig `json:"categories"`
	}

	// Config main router config structure
	Config struct {
		Prefix struct {
			Value         []string `json:"prefixes"`
			CaseSensitive bool     `json:"caseSensitive"`
		} `json:"prefix"`

		Language LanguageConfig `json:"language,omitempty"`
	}
)

// LoadConfig loads config from file
func LoadConfig(configPath string, languageConfigPath string) *Config {
	file, err := ioutil.ReadFile(configPath)

	if err != nil {
		fmt.Println("Failed loading config ", err)
		return nil
	}

	var config Config
	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println("Failed unmarshaling config file", err)
		return nil
	}

	err = config.LoadLanguage(languageConfigPath)

	if err != nil {
		fmt.Println("Failed loading language config ", err)
		return nil
	}

	return &config
}

// LoadLanguage loads language from file to Config object
func (config *Config) LoadLanguage(languageConfigPath string) error {
	file, err := ioutil.ReadFile(languageConfigPath)

	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &config.Language)

	if err != nil {
		return err
	}

	return nil
}
