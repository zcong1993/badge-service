package adapter

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/tidwall/gjson"
	"github.com/zcong1993/badge-service/utils"
)

var (
	bundlephobiaUnknownInput = makeUnknownTopicInput("bundle size")
	bundlephobiaTopics       = []string{"size", "gzip"}
	bundlephobiaUnknown      = BadgeInput{
		Subject: "bundle size",
		Status:  "unknown",
		Color:   "grey",
	}
)

// BundlephobiaApi is bundlephobia api provider
func BundlephobiaApi(args ...string) BadgeInput {
	if len(args) != 3 {
		return ErrorInput
	}
	topic := args[0]
	project := args[1]
	extra := args[2]

	println(topic)

	if !utils.IsOneOf(bundlephobiaTopics, topic) {
		return bundlephobiaUnknownInput
	}

	if extra != "" {
		project += "/" + extra
	}

	endpoint := fmt.Sprintf("https://bundlephobia.com/api/size?package=%s", project)

	resp, err := utils.Get(endpoint)
	if err != nil {
		return ErrorInput
	}
	if gjson.Get(string(resp), "error").String() != "" {
		return bundlephobiaUnknown
	}

	input := BadgeInput{
		Color: "blue",
	}

	size := gjson.Get(string(resp), "size").Uint()
	gzip := gjson.Get(string(resp), "gzip").Uint()

	switch topic {
	case "size":
		input.Subject = "size"
		input.Status = humanize.Bytes(size)
		break
	case "gzip":
		input.Subject = "gzip"
		input.Status = humanize.Bytes(gzip)
		break
	default:
		return ErrorInput
	}

	return input
}
