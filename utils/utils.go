package utils

import (
	"regexp"
	"strings"
)

// ParamsOrDefault ignore the extra params or add "" to make sure params lens equal
func ParamsOrDefault(str string, n int) []string {
	if len(str) != 0 && strings.HasPrefix(str, "/") {
		str = str[1:]
	}
	tmp := strings.Split(str, "/")
	res := make([]string, n)
	for i := 0; i < minInt(len(tmp), n); i++ {
		res[i] = tmp[i]
	}
	return res
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// StringOrDefault return default value if "" given
func StringOrDefault(str, dft string) string {
	if str == "" {
		return dft
	}
	return str
}

// IsOneOf check if array contains an item
func IsOneOf(strArr []string, item string) bool {
	for _, it := range strArr {
		if it == item {
			return true
		}
	}
	return false
}

// VersionRexp is regexp for matching unstable version
var VersionRexp = regexp.MustCompile(`(?i)\b(alpha|beta|canary|rc)`)

// VersionColor return different colors by version
func VersionColor(v string) string {
	if VersionRexp.MatchString(v) {
		return "cyan"
	}
	if strings.HasPrefix(v, "0.") {
		return "orange"
	}
	return "blue"
}
