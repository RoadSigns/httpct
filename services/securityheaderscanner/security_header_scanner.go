package securityheaderscanner

// Scan looks to see if the headers contain the ones requested and generated a new slice of SecurityHeaders with their status
func Scan(headers Headers) []SecurityHeader {
	var securityHeaders []SecurityHeader
	securityHeaders = append(securityHeaders, checkForXFrameOptions(headers))
	securityHeaders = append(securityHeaders, checkForXXssProtection(headers))
	securityHeaders = append(securityHeaders, checkForContentSecurityPolicy(headers))
	//securityHeaders = append(securityHeaders, checkForStrictTransportSecurity(headers))
	return securityHeaders
}
