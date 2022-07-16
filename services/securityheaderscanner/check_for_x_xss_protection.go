package securityheaderscanner

func checkForXXssProtection(headers Headers) SecurityHeader {
	securityHeader := SecurityHeader{
		Title:   "X-Xss-Protection",
		Exists:  false,
		Content: "",
		Reason:  "X-XSS-Protection sets the configuration for the XSS Auditor built into older browsers. The recommended value was \"X-XSS-Protection: 1; mode=block\" but you should now look at Content Security Policy instead.",
		Valid:   false,
		Guide:   "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection",
	}

	headerExists := headers.Exists("X-Xss-Protection")
	if !headerExists {
		return securityHeader
	}

	securityHeader.Exists = true
	securityHeader.Content = headers.GetContent("X-Xss-Protection")

	// Check if the Content being sent in the header is valid
	securityHeader.Valid = validatedXXssProtectionContent(securityHeader.Content)

	return securityHeader
}

func validatedXXssProtectionContent(content string) bool {
	if content == "0" {
		return true
	}

	if content == "1" {
		return true
	}

	if content == "1; mode=block" {
		return true
	}

	return false
}
