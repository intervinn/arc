package arc

import "strings"

func ConnectRoutes(vals ...string) string {
	res := ""

	for i := 0; i < len(vals)-1; i++ {
		a := vals[i]
		b := vals[i+1]

		if strings.HasSuffix(a, "/") && strings.HasPrefix(b, "/") {
			res += a[:len(a)-1] + b
		}
	}

	return res

}
