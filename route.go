package arc

import "strings"

func ConnectRoutes(a string, b string) string {
	if strings.HasSuffix(a, "/") && strings.HasPrefix(b, "/") {
		return a[:len(a)-1] + b
	}
	return a + b
}
