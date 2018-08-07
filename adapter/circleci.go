package adapter

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/zcong1993/badge-service/utils"
	"net/url"
	"strings"
)

// CircleciApi is circleci api provider
func CircleciApi(args ...string) BadgeInput {
	if len(args) != 4 {
		return ErrorInput
	}

	vcsType := args[0]
	username := args[1]
	project := args[2]
	branch := args[3]

	var b string
	if branch == "" {
		b = ""
	} else {
		b = fmt.Sprintf("/tree/%s", url.QueryEscape(branch))
	}

	endpoint := fmt.Sprintf("https://circleci.com/api/v1.1/project/%s/%s/%s%s", vcsType, username, project, b)
	resp, err := utils.Get(endpoint)
	if err != nil {
		return ErrorInput
	}
	status := gjson.Get(string(resp), "0.status").String()
	return BadgeInput{
		Subject: "circleci",
		Status:  utils.StringOrDefault(strings.Replace(status, "_", " ", -1), "not found"),
		Color:   getColor(status),
	}
}

func getColor(status string) string {
	switch status {
	case "infrastructure_fail", "timedout", "failed", "no_tests":
		return "red"
	case "canceled", "not_run", "not_running":
		return "grey"
	case "queued", "scheduled":
		return "yellow"
	case "retried", "running":
		return "orange"
	case "fixed", "success":
		return "green"
	default:
		return "grey"
	}
}
