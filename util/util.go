package util

import "strconv"

func StrToInt64(params string) int64 {
	i, _ := strconv.ParseInt(params, 10, 64)
	return i
}

// https://play.golang.org/p/Qg_uv_inCek
// contains checks if a string is present in a slice
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
