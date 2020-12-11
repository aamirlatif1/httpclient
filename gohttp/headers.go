package gohttp

import (
	"github.com/aamirlatif1/httpclient/mime"
	"net/http"
)

func makeHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

func (c *httpClient) getRequestHeaders(headers http.Header) http.Header {
	allHeaders := make(http.Header)

	//Add common headers to request
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			allHeaders.Set(header, value[0])
		}
	}

	//Add custom headers to the request
	for header, value := range headers {
		if len(value) > 0 {
			allHeaders.Set(header, value[0])
		}
	}

	//Set User-Agent if not already set
	if c.builder.userAgent != ""{
		if allHeaders.Get(mime.HeadUserAgent) != "" {
			allHeaders.Set(mime.HeadUserAgent, c.builder.userAgent)
		}
	}
	return allHeaders
}