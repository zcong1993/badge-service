package utils

import (
	"github.com/dustin/go-humanize"
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

// Colors is color name to hex map
var Colors = map[string]string{
	"brightgreen": "44CC11",
	"limegreen":   "97CA00",
	"yellow":      "DFB317",
	"orange":      "FE7D37",
	"red":         "E05D44",
	"blue":        "007EC6",
	"pink":        "FF69B4",
}

// SizeColor return different color by file size
func SizeColor(s uint64) string {
	if s < humanize.KiByte*100 {
		return Colors["brightgreen"]
	}
	if s < humanize.MiByte {
		return Colors["limegreen"]
	}
	if s < humanize.MiByte*5 {
		return Colors["blue"]
	}
	if s < humanize.MiByte*30 {
		println("hahaha")
		return Colors["yellow"]
	}
	if s < humanize.MiByte*100 {
		return Colors["orange"]
	}
	if s < humanize.MiByte*500 {
		return Colors["red"]
	}

	return Colors["pink"]
}

var licenseColorsMap = map[string][]string{
	"green": {
		"AFL-3.0",
		"Apache-2.0",
		"Artistic-2.0",
		"BSD-2-Clause",
		"BSD-3-Clause",
		"BSD-3-Clause-Clear",
		"BSL-1.0",
		"CC-BY-4.0",
		"ECL-2.0",
		"ISC",
		"MIT",
		"MS-PL",
		"NCSA",
		"PostgreSQL",
		"Zlib",
	},
	"orange": {
		"AGPL-3.0",
		"CC-BY-SA-4.0",
		"EPL-1.0",
		"EUPL-1.1",
		"GPL-2.0",
		"GPL-3.0",
		"LGPL-2.1",
		"LGPL-3.0",
		"LPPL-1.3c",
		"MPL-2.0",
		"MS-RL",
		"OFL-1.1",
		"OSL-3.0",
	},
	"7cd958": {"CC0-1.0", "Unlicense", "WTFPL"},
}

// LicenseColor return different color by license type
func LicenseColor(l string) string {
	for k, v := range licenseColorsMap {
		for _, vv := range v {
			if strings.Contains(l, vv) {
				return k
			}
		}
	}
	return "grey"
}
