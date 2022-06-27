package securityheaderscanner

import "strings"

type SecurityHeader struct {
	Title   string
	Exists  bool
	Content string
	Reason  string
	Valid   bool
	Guide   string
}

type Header struct {
	Title   string
	Content string
}

type Headers struct {
	headers map[string]Header
}

func (headers Headers) Exists(key string) bool {
	key = strings.ToLower(key)
	if _, ok := headers.headers[key]; ok {
		return true
	}

	return false
}

func (headers Headers) GetContent(key string) string {
	key = strings.ToLower(key)
	if headers.Exists(key) {
		return headers.headers[key].Content
	}
	return ""
}

func (headers *Headers) AddHeader(header Header) {
	if headers.headers == nil {
		headers.headers = make(map[string]Header)
	}

	headers.headers[strings.ToLower(header.Title)] = header
}
