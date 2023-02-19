package http

import (
	httpClient "net/http"
)

func GetHttpHeaders(url string) Headers {
	resp, err := httpClient.Head(url)
	if err != nil {
		panic(err.Error())
	}

	var headers Headers
	for title, content := range resp.Header {
		headers.AddHeader(Header{
			Title:   title,
			Content: content[0],
		})
	}

	return headers
}
