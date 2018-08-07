package utils

import "strings"

func ParamsOrDefault(str string, n int) []string {
	if len(str) != 0 && strings.HasPrefix(str, "/") {
		str = str[1:]
	}
	tmp := strings.Split(str, "/")
	if n < len(tmp) {
		n = len(tmp)
	}
	res := make([]string, n)
	for i := 0; i < len(tmp); i++ {
		res[i] = tmp[i]
	}
	return res
}

func StringOrDefault(str, dft string) string {
	if str == "" {
		return dft
	}
	return str
}

func IsOneOf(strArr []string, item string) bool {
	for _, it := range strArr {
		if it == item {
			return true
		}
	}
	return false
}
