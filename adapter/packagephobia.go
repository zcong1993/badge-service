package adapter

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/tidwall/gjson"
	"github.com/zcong1993/badge-service/utils"
)

var (
	pkgpUnknownInput = makeUnknownTopicInput("packagephobia")
	pkgpTopics       = []string{"publish", "install"}
)

// PackagephobiaApi is packagephobia api provider
func PackagephobiaApi(args ...string) BadgeInput {
	if len(args) != 3 {
		return ErrorInput
	}
	topic := args[0]
	project := args[1]
	extra := args[2]

	println(topic)

	if !utils.IsOneOf(pkgpTopics, topic) {
		return pkgpUnknownInput
	}

	if extra != "" {
		project += "/" + extra
	}

	endpoint := fmt.Sprintf("https://packagephobia.now.sh/api.json?p=%s", project)

	resp, err := utils.Get(endpoint)
	if err != nil {
		return ErrorInput
	}

	println(string(resp))

	size := gjson.Get(string(resp), topic+"Size").Uint()

	println(utils.SizeColor(size))

	return BadgeInput{
		Subject: topic,
		Status:  humanize.IBytes(size),
		Color:   utils.SizeColor(size),
	}
}
