package securityheaderscanner

func checkForContentSecurityPolicy(headers Headers) SecurityHeader {
	securityHeader := SecurityHeader{
		Title:   "Content-Security-Policy",
		Exists:  false,
		Content: "",
		Reason:  "Content Security Policy is an effective measure to protect your site from XSS attacks. By whitelisting sources of approved content, you can prevent the browser from loading malicious assets. Analyse this policy in more detail. You can sign up for a free account on Report URI to collect reports about problems on your site.",
		Valid:   false,
		Guide:   "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy",
	}

	headerExists := headers.Exists("Content-Security-Policy")
	if !headerExists {
		return securityHeader
	}

	securityHeader.Exists = true
	securityHeader.Content = headers.GetContent("Content-Security-Policy")

	// Check if the Content being sent in the header is valid
	securityHeader.Valid = true

	return securityHeader
}
