package adapter

import (
	"errors"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/tidwall/gjson"
	"github.com/zcong1993/badge-service/utils"
	"os"
)

// GITHUB_TOKEN is our github api token cause v4 need it
var GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")

// GITHUB_TOPICS is github api support topic
var GITHUB_TOPICS = []string{"stars", "forks", "watchers", "release", "tag", "issues", "open-issues", "license"}

var defaultGithubErrorResp = BadgeInput{
	Subject: "github",
	Status:  "unknown topic",
	Color:   "grey",
}

func graphqlQuery(q string) ([]byte, error) {
	query := map[string]string{
		"query": q,
	}
	header := map[string]string{
		"Accept": "application/vnd.github.hawkgirl-preview+json",
	}
	if GITHUB_TOKEN != "" {
		header["Authorization"] = fmt.Sprintf("bearer %s", GITHUB_TOKEN)
	}
	return utils.Post("https://api.github.com/graphql", query, header)
}

func getInfo(tp, user, repo string) (string, error) {
	resp, err := graphqlQuery(fmt.Sprintf(`
    query {
      repository(owner:"%s", name:"%s") {
        releases(last: 1) {
          nodes {
        	tag {
              name
            }
          }
        }

		tags:refs(refPrefix:"refs/tags/", last:1) {
          edges {
            tag:node {
              name
            }
          }
        }

		licenseInfo {
		  name
		}
      }
    }
`, user, repo))

	if err != nil {
		return "0", err
	}

	errs := gjson.Get(string(resp), "errors").String()
	if errs != "" {
		return "0", errors.New(errs)
	}

	tag := "unknown"

	switch tp {
	case "release":
		tag = gjson.Get(string(resp), fmt.Sprintf("data.repository.%s.%s.0.tag.name", "releases", "nodes")).String()
	case "tag":
		tag = gjson.Get(string(resp), fmt.Sprintf("data.repository.%s.%s.0.tag.name", "tags", "edges")).String()
	case "license":
		tag = gjson.Get(string(resp), "data.repository.licenseInfo.name").String()
	}

	return utils.StringOrDefault(tag, "unknown"), nil
}

func getCount(tp, user, repo string) (string, error) {
	resp, err := graphqlQuery(fmt.Sprintf(`
    query {
      repository(owner:"%s", name:"%s") {
        stargazers {
          totalCount
        }
		forks {
          totalCount
        }
		watchers {
		  totalCount
		}
		openIssues:issues(states:OPEN) {
   		  totalCount
    	}
    	issues() {
          totalCount
		}
      }
    }
`, user, repo))

	if err != nil {
		return "0", err
	}

	errs := gjson.Get(string(resp), "errors").String()
	if errs != "" {
		return "0", errors.New(errs)
	}
	if tp == "open-issues" {
		tp = "openIssues"
	}
	stars := gjson.Get(string(resp), fmt.Sprintf("data.repository.%s.totalCount", tp)).Float()
	return utils.StringOrDefault(humanize.SI(stars, ""), "0"), nil
}

// GithubApi is github api provider
func GithubApi(args ...string) BadgeInput {
	if len(args) != 3 {
		return ErrorInput
	}

	topic := args[0]
	user := args[1]
	repo := args[2]

	if !utils.IsOneOf(GITHUB_TOPICS, topic) {
		return defaultGithubErrorResp
	}

	switch topic {
	case "stars", "forks", "watchers", "issues", "open-issues":
		tp := topic
		if tp == "stars" {
			tp = "stargazers"
		}
		resp, err := getCount(tp, user, repo)
		if err != nil {
			return defaultErrorResp
		}
		return BadgeInput{
			Subject: topic,
			Status:  resp,
			Color:   "blue",
		}
	case "release", "tag", "license":
		resp, err := getInfo(topic, user, repo)
		if err != nil {
			return defaultErrorResp
		}
		return BadgeInput{
			Subject: topic,
			Status:  resp,
			Color:   "blue",
		}
	default:
		return defaultGithubErrorResp
	}
}
