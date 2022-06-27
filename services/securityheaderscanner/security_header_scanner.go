package securityheaderscanner

// Replace these with dependencies from the same package due to "import cycle not allowed"
func Scan(headers Headers) []SecurityHeader {
	var securityHeaders []SecurityHeader
	securityHeaders = append(securityHeaders, checkForXFrameOptions(headers))
	//securityHeaders = append(securityHeaders, checkForXXssProtection(headers))
	//securityHeaders = append(securityHeaders, checkForContentSecurityPolicy(headers))
	//securityHeaders = append(securityHeaders, checkForStrictTransportSecurity(headers))
	return securityHeaders
}
