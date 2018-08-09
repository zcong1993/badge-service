package utils

import (
	"math"
	"regexp"
	"strconv"
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
	if v == "" {
		return "grey"
	}
	if VersionRexp.MatchString(v) {
		return "cyan"
	}
	if strings.HasPrefix(v, "0.") {
		return "orange"
	}
	return "blue"
}

// CovColor return different colors by coverage
func CovColor(p float64) string {
	if p < 35 {
		return "red"
	}
	if p < 70 {
		return "orange"
	}
	if p < 85 {
		return "EEAA22"
	}
	if p < 90 {
		return "99CC09"
	}
	return "green"
}

// NormalizeVersion always make version start with 'v'
func NormalizeVersion(v string) string {
	if v == "" {
		return "unknown"
	}
	if strings.HasPrefix(v, "v") {
		return v
	}
	return "v" + v
}

// String2Float64 convert string to float64
func String2Float64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return float64(0)
	}
	return f
}

// Star generate stars string by rate
func Star(rate float64, total int) string {
	res := ""
	fullCount := 0
	halfCount := 0
	if rate >= float64(total) {
		fullCount = total
		halfCount = 0
	} else {
		fullCount = int(rate)
		rest := rate - math.Floor(rate)
		if rest >= 0.5 {
			fullCount++
		}
		halfCount = total - fullCount
	}
	res += strings.Repeat("★", fullCount)
	res += strings.Repeat("☆", halfCount)
	return res
}
