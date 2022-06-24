package services

func CheckSecurityHeaders(headers Headers) []SecurityHeader {
	var securityHeaders []SecurityHeader
	securityHeaders = append(securityHeaders, checkForXFrameOptions(headers))
	return securityHeaders
}

func checkForXFrameOptions(headers Headers) SecurityHeader {
	securityHeader := SecurityHeader{
		Title:   "X-Frame-Options",
		Exists:  false,
		Content: "",
		Reason:  "X-Frame-Options tells the browser whether you want to allow your site to be framed or not. By preventing a browser from framing your site you can defend against attacks like clickjacking.",
	}

	xFrameOptionHeaderExists := headers.Exists("X-Frame-Options")
	if !xFrameOptionHeaderExists {
		return securityHeader
	}

	securityHeader.Exists = true
	securityHeader.Content = headers.GetContent("X-Frame-Options")

	return securityHeader
}
