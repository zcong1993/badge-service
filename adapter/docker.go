package adapter

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/tidwall/gjson"
	"github.com/zcong1993/badge-service/utils"
)

var defaultErrorResp = BadgeInput{
	Subject: "docker",
	Status:  "unknown topic",
	Color:   "grey",
}

var VALID_TOPICS = []string{"stars", "pulls"}

func DockerApi(topic, namespace, name string) BadgeInput {
	if !utils.IsOneOf(VALID_TOPICS, topic) {
		return defaultErrorResp
	}

	endpoint := fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/%s", namespace, name)
	resp, err := utils.Get(endpoint)
	if err != nil {
		return ErrorInput
	}
	pullCount := gjson.Get(string(resp), "pull_count").Float()
	starCount := gjson.Get(string(resp), "star_count").Float()
	switch topic {
	case "stars":
		return BadgeInput{
			Subject: "docker stars",
			Status:  utils.StringOrDefault(humanize.SI(starCount, ""), "0"),
			Color:   "blue",
		}
	case "pulls":
		return BadgeInput{
			Subject: "docker pulls",
			Status:  utils.StringOrDefault(humanize.SI(pullCount, ""), "0"),
			Color:   "blue",
		}
	default:
		return defaultErrorResp
	}
}
