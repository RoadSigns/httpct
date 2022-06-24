package services

func CheckSecurityHeaders(headers Headers) []SecurityHeader {
	var securityHeaders []SecurityHeader
	securityHeaders = append(securityHeaders, checkForXFrameOptions(headers))
	securityHeaders = append(securityHeaders, checkForXXssProtection(headers))
	securityHeaders = append(securityHeaders, checkForContentSecurityPolicy(headers))
	securityHeaders = append(securityHeaders, checkForStrictTransportSecurity(headers))
	return securityHeaders
}

func checkForXFrameOptions(headers Headers) SecurityHeader {
	return generateSecurityHeader(
		"X-Frame-Options",
		"X-Frame-Options tells the browser whether you want to allow your site to be framed or not. By preventing a browser from framing your site you can defend against attacks like clickjacking.",
		headers,
	)
}

func checkForXXssProtection(headers Headers) SecurityHeader {
	return generateSecurityHeader(
		"X-Xss-Protection",
		"X-Xss-Protection sets the configuration for the XSS Auditor built into older browsers. The recommended value was \"X-XSS-Protection: 1; mode=block\" but you should now look at Content Security Policy instead.",
		headers,
	)
}

func checkForStrictTransportSecurity(headers Headers) SecurityHeader {
	return generateSecurityHeader(
		"Strict-Transport-Security",
		"HTTP Strict Transport Security is an excellent feature to support on your site and strengthens your implementation of TLS by getting the User Agent to enforce the use of HTTPS.",
		headers,
	)
}

func checkForContentSecurityPolicy(headers Headers) SecurityHeader {
	return generateSecurityHeader(
		"Content-Security-Policy",
		"Content Security Policy is an effective measure to protect your site from XSS attacks. By whitelisting sources of approved content, you can prevent the browser from loading malicious assets.",
		headers,
	)
}

func generateSecurityHeader(title, reason string, headers Headers) SecurityHeader {
	securityHeader := SecurityHeader{
		Title:   title,
		Exists:  false,
		Content: "",
		Reason:  reason,
	}

	headerExists := headers.Exists(title)
	if !headerExists {
		return securityHeader
	}

	securityHeader.Exists = true
	securityHeader.Content = headers.GetContent(title)

	return securityHeader
}
