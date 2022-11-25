package services

import "github.com/roadsigns/http-header-scanner/services/securityheaderscanner"

func CheckSecurityHeaders(headers Headers) []securityheaderscanner.SecurityHeader {
	var scannerHeaders securityheaderscanner.Headers
	for _, header := range headers.headers {
		scannerHeaders.AddHeader(securityheaderscanner.Header{
			Title:   header.Title,
			Content: header.Content,
		})
	}

	return securityheaderscanner.Scan(scannerHeaders)
}
