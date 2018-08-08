package adapter

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/zcong1993/badge-service/utils"
)

// AppveyorApi is appveyor api provider
func AppveyorApi(args ...string) BadgeInput {
	if len(args) != 3 {
		return ErrorInput
	}
	username := args[0]
	project := args[1]
	branch := utils.StringOrDefault(args[2], "master")

	endpoint := fmt.Sprintf("https://ci.appveyor.com/api/projects/%s/%s/branch/%s", username, project, branch)
	resp, err := utils.Get(endpoint)
	if err != nil {
		return ErrorInput
	}

	message := gjson.Get(string(resp), "message").String()

	if message != "" {
		return ErrorInput
	}

	status := gjson.Get(string(resp), "build.status").String()

	input := BadgeInput{
		Subject: "appveyor",
		Status:  status,
		Color:   "red",
	}

	if status == "success" {
		input.Color = "green"
	}

	return input
}
