package adapter

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/tidwall/gjson"
	"github.com/zcong1993/badge-service/utils"
	"regexp"
	"strings"
)

var (
	cwsUnknown = BadgeInput{
		Subject: "chrome web store",
		Status:  "unknown",
		Color:   "grey",
	}
	cwsUnknownTopic = makeUnknownTopicInput("chrome web store")
	cwsTopics       = []string{"version", "users", "price", "rating", "stars", "rating-count"}
	notFoundReg     = regexp.MustCompile(`(Not Found)`)
)

// ChromeWebStoreApi is chrome web store api provider
func ChromeWebStoreApi(args ...string) BadgeInput {
	if len(args) != 2 {
		return ErrorInput
	}
	topic := args[0]
	id := args[1]

	if !utils.IsOneOf(cwsTopics, topic) {
		return cwsUnknownTopic
	}

	resp, err := utils.Post(fmt.Sprintf("https://chrome.google.com/webstore/ajax/detail?id=%s&pv=20180301", id), nil, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return ErrorInput
	}

	if notFoundReg.Match(resp) {
		return cwsUnknown
	}
	json := string(resp)[5:]
	//println(topic, json, id)
	switch topic {
	case "version":
		v := gjson.Get(json, "1.1.6").String()

		return BadgeInput{
			Subject: "chrome web store",
			Status:  utils.NormalizeVersion(v),
			Color:   utils.VersionColor(v),
		}
	case "users":
		c := gjson.Get(json, "1.1.0.23").String()
		count := utils.String2Float64(strings.Replace(c, ",", "", -1))

		return BadgeInput{
			Subject: "users",
			Status:  humanize.SIWithDigits(count, 2, ""),
			Color:   "green",
		}
	case "price":
		c := gjson.Get(json, "1.1.0.30").String()
		return BadgeInput{
			Subject: "price",
			Status:  utils.StringOrDefault(c, "0"),
			Color:   "green",
		}
	case "rating", "stars":
		c := gjson.Get(json, "1.1.0.12").Float()
		status := humanize.FtoaWithDigits(c, 1)
		if topic == "stars" {
			status = utils.Star(c, 5)
		}
		return BadgeInput{
			Subject: topic,
			Status:  utils.StringOrDefault(status, "0"),
			Color:   "blue",
		}
	case "rating-count":
		c := gjson.Get(json, "1.1.0.22").Float()
		return BadgeInput{
			Subject: topic,
			Status:  humanize.SIWithDigits(c, 2, ""),
			Color:   "green",
		}
	}

	return ErrorInput
}
