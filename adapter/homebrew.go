package adapter

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/zcong1993/badge-service/utils"
)

var (
	homebrewUnknownInput = makeUnknownTopicInput("homebrew")
	homebresTopics       = []string{"version"}
)

// HomebrewApi is homebrew api provider
func HomebrewApi(args ...string) BadgeInput {
	if len(args) != 2 {
		return ErrorInput
	}

	topic := args[0]
	name := args[1]

	if !utils.IsOneOf(homebresTopics, topic) {
		return homebrewUnknownInput
	}

	endpoint := fmt.Sprintf("https://formulae.brew.sh/api/formula/%s.json", name)
	resp, err := utils.Get(endpoint)
	if err != nil {
		return ErrorInput
	}

	switch topic {
	case "version":
		v := gjson.Get(string(resp), "versions.stable").String()
		return BadgeInput{
			Subject: "homebrew",
			Status:  utils.NormalizeVersion(v),
			Color:   utils.VersionColor(v),
		}
	}

	return ErrorInput
}
