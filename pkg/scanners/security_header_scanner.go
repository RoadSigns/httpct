package scanners

import "github.com/roadsigns/httpct/pkg/http"

// Scan looks to see if the headers contain the ones requested and generated a new slice of SecurityHeaders with their status
func Scan(headers http.Headers) []SecurityHeader {
	var securityHeaders []SecurityHeader
	securityHeaders = append(securityHeaders, checkForXFrameOptions(headers))
	securityHeaders = append(securityHeaders, checkForXXssProtection(headers))
	securityHeaders = append(securityHeaders, checkForContentSecurityPolicy(headers))
	securityHeaders = append(securityHeaders, checkForStrictTransportSecurity(headers))
	securityHeaders = append(securityHeaders, checkForPermissionsPolicy(headers))
	return securityHeaders
}