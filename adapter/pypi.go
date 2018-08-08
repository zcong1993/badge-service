package adapter

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/zcong1993/badge-service/utils"
)

var (
	pypiUnknownInput = makeUnknownTopicInput("pypi")
	pypiTopics       = []string{"version", "license"}
)

// PypiApi is pypi api provider
func PypiApi(args ...string) BadgeInput {
	if len(args) != 2 {
		return ErrorInput
	}

	topic := args[0]
	name := args[1]

	if !utils.IsOneOf(pypiTopics, topic) {
		return pypiUnknownInput
	}

	endpoint := fmt.Sprintf("https://pypi.org/pypi/%s/json", name)
	resp, err := utils.Get(endpoint)
	if err != nil {
		return ErrorInput
	}

	switch topic {
	case "version":
		v := gjson.Get(string(resp), "info.version").String()
		return BadgeInput{
			Subject: "pypi",
			Status:  utils.StringOrDefault(v, "unknown"),
			Color:   utils.VersionColor(v),
		}
	}

	return ErrorInput
}
