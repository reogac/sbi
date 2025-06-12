package sbi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func BuildHttpRequest(sbiRequest *Request, remoteAddr string) (httpRequest *http.Request, err error) {
	var body io.Reader
	var bodyLength int64
	if sbiRequest.body != nil {
		if bodyLength, body, err = Encode(sbiRequest.body); err != nil {
			err = fmt.Errorf("Serialize request body failed: %+v", err)
			return
		}
	}

	var link *url.URL
	if link, err = url.Parse(fmt.Sprintf("http://%s/%s", remoteAddr, sbiRequest.path)); err != nil {
		err = fmt.Errorf("Parse request path failed: %+v", err)
		return
	}
	query := link.Query()
	for k, v := range sbiRequest.queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}
	link.RawQuery = query.Encode()

	if httpRequest, err = http.NewRequest(sbiRequest.method, link.String(), body); err != nil {
		err = fmt.Errorf("Build http request failed: %+v", err)
		return
	}
	httpRequest.ContentLength = bodyLength
	if len(sbiRequest.headerParams) > 0 {
		headers := http.Header{}
		for h, v := range sbiRequest.headerParams {
			headers.Set(h, v)
		}
		httpRequest.Header = headers
	}
	return
}

func (sbiResponse *Response) DecodeBody(body SbiIE) (err error) {
	return Decode(sbiResponse.contentLength, sbiResponse.body, body)
}
