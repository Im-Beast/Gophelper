package middleware

import (
	"fmt"
	"log"
	"time"

	gophelper "github.com/Im-Beast/Gophelper/internal"
)

// RateLimit structure for commands to configure their own rate limit timings
type RateLimit struct {
	Limit    int
	Duration time.Duration
}

// RateLimiter structure containing neccessary information about cooldowned users
type RateLimiter struct {
	Initialized    bool
	CmdCountMap    map[string]int
	LastRequestMap map[string]int64
	CooldownMap    map[string]int64
}

var rateLimiters = make(map[string]map[*gophelper.Command]RateLimiter)

// Maintains user rate limits
func RateLimiterMiddleware(context *gophelper.CommandContext) (bool, func(*gophelper.CommandContext)) {
	command := context.Command
	authorID := context.Event.Author.ID

	routerLanguage := context.Router.Config.Language
	session := context.Session

	if rateLimiters[authorID] == nil {
		rateLimiters[authorID] = make(map[*gophelper.Command]RateLimiter)
	}

	rateLimiter := rateLimiters[authorID][command]

	if !rateLimiter.Initialized {
		rateLimiters[authorID][command] = RateLimiter{
			Initialized:    true,
			CmdCountMap:    make(map[string]int),
			LastRequestMap: make(map[string]int64),
			CooldownMap:    make(map[string]int64),
		}

		rateLimiter = rateLimiters[authorID][command]
	}

	now := time.Now().Unix()
	cooldownTime := rateLimiter.CooldownMap[authorID]

	if cooldownTime > now {
		return false, func(context *gophelper.CommandContext) {
			_, err := session.ChannelMessageSend(context.Event.ChannelID, fmt.Sprintf(routerLanguage.Errors.RateLimit, cooldownTime-now))
			if err != nil {
				log.Printf("Middleware RateLimiter errored while sending RateLimit error message: %s\n", err.Error())
			}
		}
	} else if now-rateLimiter.LastRequestMap[authorID] >= int64(command.RateLimit.Duration.Seconds()) {
		rateLimiter.CmdCountMap[authorID] = 0
	}

	rateLimiter.CmdCountMap[authorID]++
	rateLimiter.LastRequestMap[authorID] = now

	if rateLimiter.CmdCountMap[authorID] >= command.RateLimit.Limit {
		rateLimiter.CooldownMap[authorID] = time.Now().Add(command.RateLimit.Duration).Unix()
	}

	return true, nil
}
