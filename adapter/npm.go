package adapter

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/tidwall/gjson"
	"github.com/zcong1993/badge-service/utils"
	"strings"
	"time"
)

func getPkg(topic string, args ...string) BadgeInput {
	pkg := args[0]
	tag := "latest"

	if strings.HasPrefix(pkg, "@") {
		pkg = fmt.Sprintf("%s/%s", args[0], args[1])
		tag = utils.StringOrDefault(args[2], tag)
	} else {
		tag = utils.StringOrDefault(args[1], tag)
	}

	endpoint := fmt.Sprintf("https://unpkg.com/%s@%s/package.json", pkg, tag)

	resp, err := utils.Get(endpoint)
	if err != nil {
		return ErrorInput
	}

	switch topic {
	case "version":
		v := gjson.Get(string(resp), "version").String()
		sub := "npm"
		if tag != "latest" {
			sub += "@" + tag
		}
		return BadgeInput{
			Subject: sub,
			Status:  utils.NormalizeVersion(v),
			Color:   utils.VersionColor(v),
		}
	case "license":
		v := gjson.Get(string(resp), "license").String()
		return BadgeInput{
			Subject: "license",
			Status:  utils.StringOrDefault(v, "unknown"),
			Color:   utils.LicenseColor(v),
		}
	}
	return ErrorInput
}

func getNpm(period string, args ...string) BadgeInput {
	nextYear := time.Now().Year() + 1
	endpoint := "https://api.npmjs.org/downloads/"
	switch period {
	case "downloads":
		endpoint += fmt.Sprintf("range/2005-01-01:%d-01-01", nextYear)
		break
	default:
		endpoint += fmt.Sprintf("point/%s", period)
	}
	for _, v := range args {
		if v == "" {
			continue
		}
		endpoint += "/" + v
	}

	resp, err := utils.Get(endpoint)
	if err != nil {
		fmt.Printf("%s - %+v\n", endpoint, err)
		return ErrorInput
	}

	if gjson.Get(string(resp), "error").String() != "" {
		return ErrorInput
	}

	if period == "downloads" {
		return BadgeInput{
			Subject: "downloads",
			Status:  humanize.SIWithDigits(countTotalDownloads(string(resp)), 2, ""),
			Color:   "blue",
		}
	}

	count := utils.StringOrDefault(humanize.SIWithDigits(gjson.Get(string(resp), "downloads").Float(), 2, ""), "0")
	status := count + strings.Replace(period, "last-", "/", -1)

	return BadgeInput{
		Subject: "downloads",
		Status:  status,
		Color:   "blue",
	}
}

func countTotalDownloads(resp string) float64 {
	count := float64(0)
	for _, c := range gjson.Get(resp, "downloads.#.downloads").Array() {
		count += c.Float()
	}
	return count
}

// NPM_VALID_TOPICS is valid topic docker api support
var NPM_VALID_TOPICS = []string{"version", "downloads", "license", "dw", "dm", "dy"}

// NpmApi is npm api provider
func NpmApi(args ...string) BadgeInput {
	if len(args) != 4 {
		return ErrorInput
	}

	topic := args[0]
	name := args[1]
	tag := args[2]
	extra := args[3]

	if !utils.IsOneOf(NPM_VALID_TOPICS, topic) {
		return defaultErrorResp
	}

	switch topic {
	case "version", "license":
		return getPkg(topic, name, tag, extra)
	case "downloads":
		return getNpm(topic, name, tag, extra)
	case "dw":
		return getNpm("last-week", name, tag, extra)
	case "dm":
		return getNpm("last-month", name, tag, extra)
	case "dy":
		return getNpm("last-year", name, tag, extra)
	}

	return defaultErrorResp
}
