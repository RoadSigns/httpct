package actions

import (
	"github.com/roadsigns/httpct/pkg/http"
	"github.com/roadsigns/httpct/pkg/scanners"
)

func CheckSecurityHttpHeaders(data GetHttpSecurityHeadersData) scanners.HttpHeaderScanResults {
	// Get the headers
	headers := http.GetHttpHeaders(data.Domain)

	// Pass the headers to the service to check for the security ones
	securityHeaders := scanners.CheckSecurityHeaders(headers)

	// Map to a struct in the Actions layer
	return scanners.GenerateHttpHeaderHeaderScanResults(data.Domain, headers, securityHeaders)
}
