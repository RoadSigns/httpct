package securityheaderscanner

func checkForXXssProtection(headers Headers) SecurityHeader {
	securityHeader := SecurityHeader{
		Title:   "X-Xss-Protection",
		Exists:  false,
		Content: "",
		Reason:  "X-Xss-Protection response header is a feature of Internet Explorer, Chrome and Safari that stops pages from loading when they detect reflected cross-site scripting (XSS) attacks.",
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
