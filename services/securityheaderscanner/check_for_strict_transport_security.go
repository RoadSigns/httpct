package securityheaderscanner

func checkForStrictTransportSecurity(headers Headers) SecurityHeader {
	securityHeader := SecurityHeader{
		Title:   "Strict-Transport-Security",
		Exists:  false,
		Content: "",
		Reason:  "HTTP Strict Transport Security is an excellent feature to support on your site and strengthens your implementation of TLS by getting the User Agent to enforce the use of HTTPS.",
		Valid:   false,
		Guide:   "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security",
	}

	headerExists := headers.Exists("Strict-Transport-Security")
	if !headerExists {
		return securityHeader
	}

	securityHeader.Exists = true
	securityHeader.Content = headers.GetContent("Strict-Transport-Security")

	return securityHeader
}
