package scanners

import "github.com/roadsigns/httpct/pkg/http"

func CheckSecurityHeaders(headers http.Headers) []SecurityHeader {
	return Scan(headers)
}
