package gophelper

// Category command category
type Category struct {
	Name          string
	Aliases       []string
	Description   string
	ReactionEmoji string
}

var (
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
