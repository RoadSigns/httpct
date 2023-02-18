package actions

import (
	"github.com/roadsigns/httpct/services"
)

func CheckSecurityHttpHeaders(data GetHttpSecurityHeadersData) services.HttpHeaderScanResults {
	// Get the headers
	headers := services.GetHttpHeaders(data.Domain)

	// Pass the headers to the service to check for the security ones
	securityHeaders := services.CheckSecurityHeaders(headers)

	// Map to a struct in the Actions layer
	return services.GenerateHttpHeaderHeaderScanResults(data.Domain, headers, securityHeaders)
}
