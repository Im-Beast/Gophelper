package utils

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// IsMention returns whether given string is mention <@...>/<@!...>
func IsMention(str string) bool {
	matches, _ := regexp.MatchString("<@!?.*>", str)
	return matches
}

// IsNumber returns whether given string is number
func IsNumber(str string) bool {
	matches, _ := regexp.MatchString("^\\d+$", str)
	return matches
}

// StringToInt converts string to int, if its invalid returns -1
func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}
	return num
}

// MentionToID converts mention <@...> thing to user ID
func MentionToID(mention string) string {
	return strings.Replace(strings.Replace(strings.Replace(mention, "<@", "", 1), ">", "", 1), "!", "", 1)
}

// Matches returns whether string matches second string
func Matches(str string, str2 string, caseSensitive bool) bool {
	if !caseSensitive && strings.ToLower(str) == strings.ToLower(str2) {
		return true
	} else if str == str2 {
		return true
	}

	return false
}

// MatchesPrefix returns whether string has prefix
func MatchesPrefix(str string, prefix string, caseSensitive bool) bool {
	if !caseSensitive && strings.HasPrefix(strings.ToLower(str), strings.ToLower(prefix)) {
		return true
	} else if strings.HasPrefix(str, prefix) {
		return true
	}

	return false
}

// RandomStringElement Returns random element from string array
func RandomStringElement(array []string) string {
	return array[rand.Intn(len(array))]
}

// ClampInt clamps int
func ClampInt(num int, min int, max int) int {
	switch {
	case num < min:
		return min
	case num > max:
		return max
	default:
		return num
	}
}

// IsNSFW returns whether channel with given channelID has enabled NSFW
func IsNSFW(session *discordgo.Session, channelID string) bool {
	channel, err := session.Channel(channelID)
	if err != nil {
		return false
	}
	return channel.NSFW
}

// GetStringVal if string str is not set it returns var notSet, if its set it returns var str
func GetStringVal(str string, notSet string) string {
	switch {
	case str == "":
		return notSet
	default:
		return str
	}
}
