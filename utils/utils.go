package utils

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Returns whether a given string is a mention <@...>/<@!...>
func IsMention(str string) bool {
	matches, _ := regexp.MatchString("<@!?.*>", str)
	return matches
}

// Returns whether a given string is number
func IsNumber(str string) bool {
	matches, _ := regexp.MatchString("^\\d+$", str)
	return matches
}

// Converts a mention <@...> to the user ID
func MentionToID(mention string) string {
	return strings.Replace(strings.Replace(strings.Replace(mention, "<@", "", 1), ">", "", 1), "!", "", 1)
}

// Returns whether str1 matches str2
func StringMatches(str string, str2 string, caseSensitive bool) bool {
	if !caseSensitive && strings.EqualFold(str, str2) {
		return true
	} else if str == str2 {
		return true
	}

	return false
}

// Returns whether a string has given prefix
func MatchesPrefix(str string, prefix string, caseSensitive bool) bool {
	switch {
	case strings.HasPrefix(str, prefix):
		fallthrough
	case !caseSensitive && strings.HasPrefix(strings.ToLower(str), strings.ToLower(prefix)):
		return true
	}

	fmt.Printf("str %s prefix %s cs %v | %v", str, prefix, caseSensitive, strings.HasPrefix(str, prefix))

	return false
}

// Returns random element from given string array
func RandomStringElement(array []string) string {
	return array[rand.Intn(len(array))]
}

// Clamp int between min and max values
func ClampInt(num int, min int, max int) int {
	switch {
	case num < min:
		return min
	case num > max:
		return max
	}
	return num
}

// Returns whether a channel with given channelID has NSFW enabled
func IsNSFW(session *discordgo.Session, channelID string) bool {
	channel, err := session.Channel(channelID)
	if err != nil {
		return false
	}
	return channel.NSFW
}

// If str is empty notSet is returned otherwise str is returned
func GetStringVal(str string, notSet string) string {
	if str == "" {
		return notSet
	}
	return str
}

// Returns response body contents from given webpage
func GetWebpageContent(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		return []byte{}
	}

	defer response.Body.Close()

	html, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return []byte{}
	}

	return html
}
