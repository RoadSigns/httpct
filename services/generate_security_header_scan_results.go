package services

import (
	"github.com/roadsigns/http-header-scanner/services/securityheaderscanner"
	"sort"
)

func GenerateHttpHeaderHeaderScanResults(url string, headers Headers, securityHeaders []securityheaderscanner.SecurityHeader) HttpHeaderScanResults {
	factory := HttpHeaderScanResultFactory{
		Url:             url,
		Headers:         headers,
		SecurityHeaders: securityHeaders,
	}

	return factory.generate()
}

type HttpHeaderScanResultFactory struct {
	Url             string
	Headers         Headers
	SecurityHeaders []securityheaderscanner.SecurityHeader
}

func (factory HttpHeaderScanResultFactory) generate() HttpHeaderScanResults {
	// Create Raw Title Slice
	var rawHeaders []RawHeader
	for _, header := range factory.Headers.headers {
		rawHeaders = append(rawHeaders, RawHeader{
			Title:   header.Title,
			Content: header.Content,
		})
	}

	sort.Strings(rawHeaders)

	// Get the Missing and Additional Information Security Headers
	var missingHeaders []MissingHeader
	var additionalInformationHeaders []AdditionalInformationHeader

	for _, header := range factory.SecurityHeaders {
		if header.Exists && header.Valid {
			additionalInformationHeaders = append(additionalInformationHeaders, AdditionalInformationHeader{
				Title:   header.Title,
				Exists:  header.Exists,
				Content: header.Content,
				Reason:  header.Reason,
				Valid:   header.Valid,
				Guide:   header.Guide,
			})
			continue
		}

		missingHeaders = append(missingHeaders, MissingHeader{
			Title:   header.Title,
			Exists:  header.Exists,
			Content: header.Content,
			Reason:  header.Reason,
			Valid:   header.Valid,
			Guide:   header.Guide,
		})
	}

	return HttpHeaderScanResults{
		Url:                         factory.Url,
		RawHeaders:                  rawHeaders,
		MissingHeaders:              missingHeaders,
		AdditionalInformationHeader: additionalInformationHeaders,
	}
}
