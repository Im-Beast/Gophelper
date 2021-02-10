package gophelper

import (
	"fmt"
	"time"
)

// RateLimit Structure for command to configure their own rate limit timings
type RateLimit struct {
	Limit    int
	Duration time.Duration
}

// RateLimiter Main RateLimiter structure containing information about currently cooldowned players etc
type RateLimiter struct {
	Initialized    bool
	CmdCountMap    map[string]int
	LastRequestMap map[string]int64
	CooldownMap    map[string]int64
}

var rateLimiters = make(map[string]map[*Command]RateLimiter)

// RateLimiterMiddleware Middleware that checks user rate limit
func RateLimiterMiddleware(context *CommandContext) (bool, func(*CommandContext)) {
	command := context.Command
	authorID := context.Event.Author.ID

	routerLanguage := context.Router.Config.Language
	session := context.Session

	if rateLimiters[authorID] == nil {
		rateLimiters[authorID] = make(map[*Command]RateLimiter)
	}

	rateLimiter := rateLimiters[authorID][command]

	if rateLimiter.Initialized == false {
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
		return false, func(context *CommandContext) {
			session.ChannelMessageSend(context.Event.ChannelID, fmt.Sprintf(routerLanguage.RateLimitError, cooldownTime-now))
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
