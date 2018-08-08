package adapter

import "github.com/zcong1993/badge-service/utils"

func ChromeWebStoreApi(args ...string) BadgeInput {
	if len(args) != 2 {
		return ErrorInput
	}
	topic := args[0]
	id := args[1]

	body := map[string]string{"id": id}

	resp, err := utils.Post("https://chrome.google.com/webstore/ajax/detail", body, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return ErrorInput
	}
	println(topic, string(resp), id)
	return ErrorInput
}
