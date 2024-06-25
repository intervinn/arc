package arc

import "strings"

func Pieces(r string) []string {
	s := strings.Split(r, "/")
	res := []string{}
	for _, v := range s {
		if !IsEmpty(v) {
			res = append(res, v)
		}
	}

	return res
}

func IsEmpty(s string) bool {
	return s == ""
}
