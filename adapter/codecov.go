package adapter

import (
	"fmt"
	"github.com/zcong1993/badge-service/utils"
	"github.com/tidwall/gjson"
	"github.com/dustin/go-humanize"
)

var (
	codecovUnknown = BadgeInput{
		Subject: "codecov",
		Status:  "unknown",
		Color:   "grey",
	}
	codecovTypes = []string{"github", "bitbucket", "gitlab"}
)

// CodecovApi is codecov api provider
func CodecovApi(args ...string) BadgeInput {
	if len(args) != 4 {
		return ErrorInput
	}

	vcsType := args[0]
	username := args[1]
	project := args[2]
	branch := args[3]

	endpoint := fmt.Sprintf("https://codecov.io/api/%s/%s/%s", vcsType, username, project)
	if branch != "" {
		endpoint += fmt.Sprintf("/branch/%s", branch)
	}

	resp, err := utils.Get(endpoint)
	if err != nil {
		return ErrorInput
	}

	if gjson.Get(string(resp), "error").String() != "" {
		return codecovUnknown
	}

	percent := gjson.Get(string(resp), "commit.totals.c").Float()

	return BadgeInput{
		Subject: "codecov",
		Status:  utils.StringOrDefault(humanize.FtoaWithDigits(percent, 2), "0") + "%",
		Color:   utils.CovColor(percent),
	}
}
