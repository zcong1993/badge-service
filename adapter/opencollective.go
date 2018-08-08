package adapter

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/rhymond/go-money"
	"github.com/tidwall/gjson"
	"github.com/zcong1993/badge-service/utils"
)

var (
	opencUnknownInput = makeUnknownTopicInput("opencollective")
	opencTopics       = []string{"backers", "contributors", "balance", "yearly"}
)

// OpencollectiveApi is opencollective api provider
func OpencollectiveApi(args ...string) BadgeInput {
	if len(args) != 2 {
		return ErrorInput
	}

	topic := args[0]
	name := args[1]

	if !utils.IsOneOf(opencTopics, topic) {
		return opencUnknownInput
	}

	endpoint := fmt.Sprintf("https://opencollective.com/%s.json", name)
	resp, err := utils.Get(endpoint)
	if err != nil {
		return ErrorInput
	}

	if string(resp) == "Not found" {
		r := opencUnknownInput
		r.Status = "Not Found"
		return r
	}

	input := opencUnknownInput
	input.Subject = topic

	switch topic {
	case "backers", "contributors":
		count := gjson.Get(string(resp), topic+"Count").Int()
		input.Status = utils.StringOrDefault(humanize.Comma(count), "0")
		input.Color = "green"
		break
	case "balance", "yearly":
		currency := utils.StringOrDefault(gjson.Get(string(resp), "currency").String(), "USD")
		t := topic
		if topic == "yearly" {
			t = "yearlyIncome"
			input.Subject = "yearly income"
		}
		c := gjson.Get(string(resp), t).Int()
		input.Status = money.New(c, currency).Display()
		input.Color = "green"
		break
	}

	return input
}
