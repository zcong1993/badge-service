package adapter

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/tidwall/gjson"
	"github.com/zcong1993/badge-service/utils"
)

var (
	cratesUnknownInput = makeUnknownTopicInput("crates")
	cratesTopics       = []string{"version", "downloads", "dl"}
)

// CratesApi is crates api provider
func CratesApi(args ...string) BadgeInput {
	if len(args) != 2 {
		return ErrorInput
	}

	topic := args[0]
	name := args[1]

	if !utils.IsOneOf(cratesTopics, topic) {
		return cratesUnknownInput
	}

	endpoint := fmt.Sprintf("https://crates.io/api/v1/crates/%s", name)
	resp, err := utils.Get(endpoint)
	if err != nil {
		return ErrorInput
	}

	if gjson.Get(string(resp), "errors").String() != "" {
		reason := gjson.Get(string(resp), "errors.0.detail").String()
		return BadgeInput{
			Subject: "crates",
			Status:  utils.StringOrDefault(reason, "unknown"),
			Color:   "grey",
		}
	}

	input := ErrorInput

	switch topic {
	case "version":
		v := gjson.Get(string(resp), "crate.max_version").String()
		input.Subject = "version"
		input.Status = utils.NormalizeVersion(v)
		input.Color = utils.VersionColor(v)
	case "downloads", "dl":
		d := gjson.Get(string(resp), "crate.downloads").Float()
		if topic == "dl" {
			d = gjson.Get(string(resp), "crate.recent_downloads").Float()
		}
		input.Subject = "downloads"
		input.Status = utils.StringOrDefault(humanize.SIWithDigits(d, 2, ""), "0")
		if topic == "dl" {
			input.Status += " latest version"
		}
		input.Color = "green"
	}

	return input
}
