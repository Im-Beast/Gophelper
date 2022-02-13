package gophelper

import "fmt"

// Command category structure
type Category struct {
	Name          string
	Aliases       []string
	Description   string
	ReactionEmoji string
}

// Converts string to category
func StringToCategory(strCategory string) (*Category, error) {
	for _, category := range CATEGORIES {
		if category.Name == strCategory {
			return category, nil
		}

		for _, alias := range category.Aliases {
			if alias == strCategory {
				return category, nil
			}
		}
	}

	return nil, fmt.Errorf("couldn't parse %s to any category", strCategory)
}

var (
	CATEGORIES = []*Category{
		CATEGORY_MISC,
		CATEGORY_CONFIG,
		CATEGORY_FUN,
		CATEGORY_MOD,
		CATEGORY_GAMES,
	}

	//CATEGORY_MISC Miscellaneous commands
	CATEGORY_MISC = &Category{
		Name:          "Miscellaneous",
		Aliases:       []string{"misc", "miscellaneous"},
		Description:   "Other stuff",
		ReactionEmoji: "✨",
	}

	//CATEGORY_CONFIG Bot config commands
	CATEGORY_CONFIG = &Category{
		Name:          "Bot config",
		Aliases:       []string{"config", "botconfig", "bot_config"},
		Description:   "Bot config related commands",
		ReactionEmoji: "⚙️",
	}

	//CATEGORY_FUN Fun commands
	CATEGORY_FUN = &Category{
		Name:          "Fun",
		Aliases:       []string{"fun"},
		Description:   "Haha funni",
		ReactionEmoji: "😀",
	}

	//CATEGORY_MOD Moderation commands
	CATEGORY_MOD = &Category{
		Name:          "Moderation",
		Aliases:       []string{"moderation", "mod"},
		Description:   "Gang gang",
		ReactionEmoji: "🛡️",
	}

	//CATEGORY_GAMES Game commands
	CATEGORY_GAMES = &Category{
		Name:          "Games",
		Aliases:       []string{"games"},
		Description:   "Simple games vs bot",
		ReactionEmoji: "🎮",
	}
)
