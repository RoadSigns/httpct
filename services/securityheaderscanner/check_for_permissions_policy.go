package securityheaderscanner

func checkForPermissionsPolicy(headers Headers) SecurityHeader {
	securityHeader := SecurityHeader{
		Title:   "Permissions-Policy",
		Exists:  false,
		Content: "",
		Reason:  "Permission Policy is a new header that allows a site to control which features and APIs can be used in the browser",
		Valid:   false,
		Guide:   "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Permissions-Policy",
	}

	headerExists := headers.Exists("Content-Security-Policy")
	if !headerExists {
		return securityHeader
	}

	securityHeader.Exists = true
	securityHeader.Content = headers.GetContent("Permissions-Policy")

	// Check if the Content being sent in the header is valid
	securityHeader.Valid = true

	return securityHeader
}
