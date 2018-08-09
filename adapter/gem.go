package adapter

import (
	"bytes"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/tidwall/gjson"
	"github.com/zcong1993/badge-service/utils"
	"strings"
)

var (
	gemUnknownInput = makeUnknownTopicInput("gem")
	gemTopics       = []string{"version", "name", "downloads", "dv"}
	preKeys         = []string{".rc", ".beta", "-rc", "-beta"}
)

func isPreVersion(v string) bool {
	for _, pp := range preKeys {
		if strings.Contains(v, pp) {
			return true
		}
	}
	return false
}

// GemApi is gem api provider
func GemApi(args ...string) BadgeInput {
	if len(args) != 3 {
		return ErrorInput
	}

	topic := args[0]
	name := args[1]
	extra := args[2]

	if !utils.IsOneOf(gemTopics, topic) {
		return gemUnknownInput
	}

	endpoint := "https://rubygems.org/api/v1/%s.json"
	if topic != "version" {
		endpoint = fmt.Sprintf(endpoint, "gems/"+name)
	} else {
		endpoint = fmt.Sprintf(endpoint, "versions/"+name)
	}
	resp, err := utils.Get(endpoint)
	if err != nil {
		return ErrorInput
	}

	if bytes.Contains(resp, []byte("This rubygem could not be found.")) {
		r := gemUnknownInput
		r.Status = "unknown"
		return r
	}

	switch topic {
	case "version":
		vs := gjson.GetBytes(resp, "#.number").Array()
		v := "0.0.0"
		if len(vs) > 0 {
			switch extra {
			case "pre", "stable":
				for _, vv := range vs {
					if extra == "pre" && isPreVersion(vv.String()) {
						v = vv.String()
						break
					}
					if !isPreVersion(vv.String()) {
						v = vv.String()
						break
					}
				}
			default:
				v = vs[0].String()
			}
		}
		return BadgeInput{
			Subject: "gem",
			Status:  utils.NormalizeVersion(v),
			Color:   utils.VersionColor(v),
		}
	case "name":
		name := gjson.GetBytes(resp, "name").String()
		return BadgeInput{
			Subject: "gem",
			Status:  name,
			Color:   "green",
		}
	case "downloads", "dv":
		key := topic
		if topic == "dv" {
			key = "version_downloads"
		}
		c := gjson.GetBytes(resp, key).Float()
		status := humanize.SIWithDigits(c, 1, "")
		if topic == "dv" {
			status += " / version"
		}
		return BadgeInput{
			Subject: "downloads",
			Status:  status,
			Color:   "blue",
		}
	}

	return ErrorInput
}
