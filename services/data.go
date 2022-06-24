package services

type SecurityHeader struct {
	Title   string
	Exists  bool
	Content string
	Reason  string
}

type Header struct {
	Title   string
	Content string
}

type Headers struct {
	headers map[string]Header
}

func (headers Headers) Exists(key string) bool {
	if _, ok := headers.headers[key]; ok {
		return true
	}

	return false
}

func (headers Headers) GetContent(key string) string {
	if headers.Exists(key) {
		return headers.headers[key].Content
	}
	return ""
}

func (headers Headers) AddHeader(header Header) {
	headers.headers[header.Title] = header
}

type HttpHeaderScanResults struct {
	Result                      Result
	RawHeaders                  []RawHeader
	MissingHeaders              []MissingHeader
	AdditionalInformationHeader []AdditionalInformationHeader
}

type Result struct {
	Site        string
	TimeScanned string
	Rank        string
}

type RawHeader struct {
	Title   string
	Content string
}

type MissingHeader struct {
	Title   string
	Exists  bool
	Content string
	Reason  string
}

type AdditionalInformationHeader struct {
	Title   string
	Exists  bool
	Content string
	Reason  string
}
