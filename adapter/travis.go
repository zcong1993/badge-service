package adapter

import (
	"fmt"
	"github.com/zcong1993/badge-service/utils"
	"regexp"
)

var (
	// RegPass is regexp for matching success build
	RegPass = regexp.MustCompile(`passed|passing`)
	// RegFail is regexp for matching failed build
	RegFail = regexp.MustCompile(`failed|failing`)
)

// TravisApi is travis api provider
func TravisApi(args ...string) BadgeInput {
	if len(args) != 3 {
		return ErrorInput
	}
	username := args[0]
	project := args[1]
	branch := utils.StringOrDefault(args[2], "master")

	endpoint := fmt.Sprintf("https://api.travis-ci.org/%s/%s.svg?branch=%s", username, project, branch)
	resp, err := utils.Get(endpoint)
	if err != nil {
		return ErrorInput
	}
	if RegPass.Match(resp) {
		return BadgeInput{
			Subject: "travis",
			Status:  "passing",
			Color:   "green",
		}
	}
	if RegFail.Match(resp) {
		return BadgeInput{
			Subject: "travis",
			Status:  "failed",
			Color:   "red",
		}
	}
	return BadgeInput{
		Subject: "travis",
		Status:  "unknown",
		Color:   "grey",
	}
}
