package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	gophelper "github.com/Im-Beast/Gophelper/internal"
	middleware "github.com/Im-Beast/Gophelper/middleware"
	"github.com/Im-Beast/Gophelper/utils"
	"github.com/bwmarrin/discordgo"
)

// Pinterest cool pics
var Pinterest = &gophelper.Command{
	ID: "Pinterest",

	Name:    "📷 Pinterest",
	Aliases: []string{"pinterest"},

	Category: gophelper.CATEGORY_FUN,

	Description: "Cool pics",

	RateLimit: middleware.RateLimit{
		Limit:    1,
		Duration: time.Second * 5,
	},

	Handler: func(context *gophelper.CommandContext) {
		msg := context.Event.Message
		session := context.Session
		args := context.Arguments
		lang := context.Router.Config.Language

		if len(args) < 1 {
			return
		}

		queryPhrase := strings.Join(args, " ")

		images := getPinterestQuery(queryPhrase).ResourceResponse.Data.Results

		if len(images) == 0 {
			_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
				Title:       lang.Errors.NoResultsFound.Title,
				Description: lang.Errors.NoResultsFound.Message,
			})

			if err != nil {
				log.Printf("Command Pinterest errored while sending embed message: %s\n", err.Error())
			}
			return
		}

		img := images[rand.Intn(len(images))]

		_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
			Title:       img.RichSummary.DisplayName,
			Description: img.Description,
			Image: &discordgo.MessageEmbedImage{
				URL: img.Images[PinterestSizes.Original].Url,
			},
			Footer: &discordgo.MessageEmbedFooter{
				IconURL: "https://s.pinimg.com/webapp/favicon-54a5b2af.png",
				Text:    fmt.Sprintf("Made by %s (%s)", img.Pinner.FullName, img.Pinner.Username),
			},
		})

		if err != nil {
			log.Printf("Command Pinterest errored while sending embed message: %s\n", err.Error())
		}
	},
}

type Sizes struct {
	Tiny     string
	Small    string
	Medium   string
	Big      string
	Huge     string
	Original string
}

var PinterestSizes = &Sizes{
	Tiny:     "136x136",
	Small:    "170x",
	Medium:   "236x",
	Big:      "474x",
	Huge:     "736x",
	Original: "orig",
}

type PinterestResult struct {
	Description string `json:"description"`

	RichSummary struct {
		URL         string `json:"url"`
		DisplayName string `json:"display_name"`
	} `json:"rich_summary"`

	Pinner struct {
		FullName string `json:"full_name"`
		Username string `json:"username"`
	} `json:"pinner"`

	Images map[string]struct {
		Height int    `json:"height"`
		Width  int    `json:"width"`
		Url    string `json:"url"`
	} `json:"images"`
}

type PinterestResponse struct {
	ResourceResponse struct {
		Data struct {
			Results []PinterestResult `json:"results"`
		} `json:"data"`
	} `json:"resource_response"`
}

func getPinterestQuery(query string) *PinterestResponse {
	query = strings.Join(strings.Fields(query), "_")

	url := `https://www.pinterest.com/resource/BaseSearchResource/get/?source_url=/search/pins/?q=` + query + `&rs=typed&term_meta[]=` + query + `%7Ctyped&data={%22options%22:{%22page_size%22:25,%22query%22:%22` + query + `%22,%22scope%22:%22pins%22,%22bookmarks%22:[%22Y2JVSG81V2sxcmNHRlpWM1J5VFVad1YxWllaR3hpVmxwSldWVmFkMkZXV2xWV2JteFhUVzVTY2xWNlNrdFNhelZaVld4V1YxSlZjR2hYVjNoaFZqQTFWMVZzWkZaaVJuQlBXV3RvUTJWR1drZFZiR1JWWWxWd1dGVnNVa05YUmxwR1kwVmtWV0pHVmpSV2JGcExWbFpPY2s5V1pGTmlXR041Vm1wS01HRXhWWGxUYkdScFUwWktWVlpyVm1GVU1XeFlUVmR3YkZKc1NqRlpNRlV4Vkd4S1ZWSnNiRmROVmtwWVZrZDRZVk5IVmtsUmJGWnBZbXRLU0Zkc1dtRmtNbEpIVm14c2FGSnVRbk5aYTFaM1pVWmFTR1JHVG1wTmExWXpWRlpvUjFkSFNsaGxTRkpXWWtaS1dGVnFSbUZqVmxKeFZHeEdWbFpFUVRWYWExcFhVMWRLTmxWdGVGTk5XRUpIVmxSSmVHTXhVbk5UV0doWVltdEtWbGxyWkZOVVJteFdWbFJHV0ZKck5UQlVWbVJIVmpGS2NtTkVRbGRTUlZwVVdUSnpNVlpyT1ZaV2JGSllVMFZLVWxadGRHRlNhekZYVld4YVlWSnJjSE5WYkZKWFUxWlZlVTFJYUZWaVJuQkhWbTF3VjFkSFNrZFRhMDVoVmpOTk1WVXdXa3RrUjBaR1RsZDRhRTFJUWpSV2Frb3dWVEZKZVZKc1pHcFNiRnBYVm10YVMxVldWbkphUms1cVlrWktNRmt3Vmt0aVIwWTJVbTVvVm1KR1NrUldNbk40WTJzeFJWSnNWbWhoTTBKUlYxZDRWbVZIVWtkWGJrWm9VbXhhYjFSV1duZFhiR1IwWkVWYVVGWnJTbE5WUmxGNFQwVXdlVk5VUWs1U01VVjNWMjB4VW1WV2NIRmlSMmhPVmtkb2NGZFhjRTVsYkhCMFZGaG9VRlpHVlhwVWFrcFRZVlpzTmxkdGFFOWlWV3cxVjFSS1MyRXhiSEZhUjJoUVVrVndjMVJxU2s1bFZUVTJWMWhvVDFaSFpETlhiVEZYWVZVNVJWSlVSbEJTUlVVeFZGUktUMkZyT1VsbFJUbFRWbTFSTkdaRVVUQmFSMGt3VDFSb2FGbFhTWGhQVkUxNVRtcFpOVmxxYkdwT1YwMHpUMGRPYUZscVZteE5iVVUwV1RKVk0wNVhSWHBPTWtsNlQxUlNiVTlIVVhoYVYwMHdUVlJuZVZwVVVUTk5NbEY2V2tkV2FrMXFUamhVYTFaWVprRTlQUT09fFVIbzVUMUZ1V21oTE1ERlpVek5CY2s1RU1XWk5hbFptVEZSR09GbFhTWGhQUkVab1dXcEthVTFYVlRWWmVrNXNXa2RTYWs1RVdUSk9la1pwVFZSak1scHRUbWhQVkVreVRVUkthbHBxV214TlJGcHFXWHBCZDAxRVVtMU9WMWt6VFVSQk0xcEVWbXBOZWtKb1dXMVpOVnBZZUU5U1ZtUTR8N2EzOTVhZTA5NjA0MTJmODRjN2VmZjU0MzFjODc1YTdlMWJiZGE0YTc2YmRjNGM3OGZlY2JhYzI3YjgzMDkyNXxORVd8%22],%22field_set_key%22:%22unauth_react%22,%22no_fetch_context_on_resource%22:false},%22context%22:{}}&_=1613332989706`
	content := utils.GetWebpageContent(url)

	var response PinterestResponse

	err := json.Unmarshal(content, &response)
	if err != nil {
		println("error while unmarshaling")
	}

	return &response
}
