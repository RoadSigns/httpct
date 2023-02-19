package scanners

import "github.com/roadsigns/httpct/pkg/http"

func checkForXFrameOptions(headers http.Headers) SecurityHeader {
	securityHeader := SecurityHeader{
		Title:   "X-Frame-Options",
		Exists:  false,
		Content: "",
		Reason:  "X-Frame-Options tells the browser whether you want to allow your site to be framed or not. By preventing a browser from framing your site you can defend against attacks like clickjacking.",
		Valid:   false,
		Guide:   "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options",
	}

	headerExists := headers.Exists("X-Frame-Options")
	if !headerExists {
		return securityHeader
	}

	securityHeader.Exists = true
	securityHeader.Content = headers.GetContent("X-Frame-Options")

	// Check if the Content being sent in the header is valid
	securityHeader.Valid = validatedXFrameOptionsContent(securityHeader.Content)

	return securityHeader
}

func validatedXFrameOptionsContent(content string) bool {
	if content == "DENY" {
		return true
	}

	if content == "SAMEORIGIN" {
		return true
	}

	return false
}
