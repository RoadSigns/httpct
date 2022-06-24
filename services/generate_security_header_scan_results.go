package services

func GenerateHttpHeaderHeaderScanResults(url string, headers Headers, securityHeaders []SecurityHeader) HttpHeaderScanResults {
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
	SecurityHeaders []SecurityHeader
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

	// Get the Missing and Additional Information Security Headers
	var missingHeaders []MissingHeader
	var additionalInformationHeaders []AdditionalInformationHeader

	for _, header := range factory.SecurityHeaders {
		if header.Exists {
			additionalInformationHeaders = append(additionalInformationHeaders, AdditionalInformationHeader{
				Title:   header.Title,
				Exists:  header.Exists,
				Content: header.Content,
				Reason:  header.Reason,
			})
			continue
		}

		missingHeaders = append(missingHeaders, MissingHeader{
			Title:   header.Title,
			Exists:  header.Exists,
			Content: header.Content,
			Reason:  header.Reason,
		})
	}

	return HttpHeaderScanResults{
		Url:                         factory.Url,
		RawHeaders:                  rawHeaders,
		MissingHeaders:              missingHeaders,
		AdditionalInformationHeader: additionalInformationHeaders,
	}
}
