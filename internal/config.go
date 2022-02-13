package gophelper

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type (
	// Structure for "stats" command
	StatsCommand struct {
		UserNotFound string `json:"userNotFound,omitempty"`
		Title        string `json:"title,omitempty"`
		CreationDate string `json:"creationDate,omitempty"`
		JoinDate     string `json:"joinDate,omitempty"`
	}

	// Structure for "8ball" command
	EightBallCommand struct {
		NoArgumentsMessage string   `json:"noArgumentsMessage,omitempty"`
		Answers            []string `json:"answers,omitempty"`
	}

	// Structure for "help" command
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
	}

	// Structure for give commands ("kitty", "doggy", "waifu", etc)
	ImageResponseCommand struct {
		Response struct {
			Mention    string `json:"mention"`
			NonMention string `json:"nonMention"`
		} `json:"response,omitempty"`
	}

	// Structure used for commands translation
	CommandConfig struct {
		Description string `json:"description,omitempty"`
		HelpCommand
		ImageResponseCommand
		EightBallCommand
		StatsCommand
	}

	// Strcuture that used for categories translation
	CategoryConfig struct {
		Description string `json:"description"`
	}

	// Translation config
	LanguageConfig struct {
		Errors struct {
			MessageSend       string `json:"messageSend"`
			TooFewPermissions string `json:"tooFewPermissions"`
			NSFWOnly          string `json:"nsfwOnly"`
			RateLimit         string `json:"rateLimit"`

			CommandNotFound struct {
				Title   string `json:"title"`
				Message string `json:"message"`
			} `json:"commandNotFound,omitempty"`
			CommandInvalidArguments struct {
				Title   string `json:"title"`
				Message string `json:"message"`
			} `json:"invalidCommandArguments,omitempty"`
			NoResultsFound struct {
				Title   string `json:"title"`
				Message string `json:"message"`
			} `json:"noResultsFound,omitempty"`
		} `json:"errors"`

		FunFacts   []string                  `json:"funFacts"`
		Commands   map[string]CommandConfig  `json:"commands"`
		Categories map[string]CategoryConfig `json:"categories"`
	}

	// Main router config structure
	Config struct {
		Prefix struct {
			Value         []string `json:"prefixes"`
			CaseSensitive bool     `json:"caseSensitive"`
		} `json:"prefix"`

		Commands struct {
			Help struct {
				Categories []string `json:"categories"`
			} `json:"help"`
		} `json:"commands"`

		Language LanguageConfig `json:"language,omitempty"`
	}
)

// Loads config from file
func LoadConfig(configPath string, languageConfigPath string) *Config {
	file, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Printf("Loading config errored: %s\n", err.Error())
		return nil
	}

	var config Config
	err = json.Unmarshal(file, &config)

	if err != nil {
		log.Printf("Loading config errored while unmarshaling json: %s\n", err.Error())
		return nil
	}

	err = config.LoadLanguage(languageConfigPath)

	if err != nil {
		log.Printf("Loading config errored while loading language: %s\n", err.Error())
		return nil
	}

	return &config
}

// Loads language from file to Config object
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
