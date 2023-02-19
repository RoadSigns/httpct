package actions

import (
	"errors"
	"net/url"
)

type GetHttpSecurityHeadersData struct {
	Domain string
}

func (data GetHttpSecurityHeadersData) ValidatedUrl() (bool, error) {
	if data.Domain == "" {
		return false, errors.New("domain is unable to be empty")
	}

	if !containsProtocol(data.Domain) {
		return false, errors.New("domain requires the protocol")
	}

	return true, nil
}

func containsProtocol(domain string) bool {
	_, err := url.ParseRequestURI(domain)
	if err != nil {
		return false
	}

	return true
}
