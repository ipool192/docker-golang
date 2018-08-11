package helpers

import "strconv"

func StringToInt(v string) int {
	r, e := strconv.Atoi(v)
	if e != nil {
		return 0
	}
	return r
}
