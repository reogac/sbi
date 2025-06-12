package sbi

import (
	"io"
	"net/url"
)

type Request struct {
	path         string
	method       string
	body         SbiIE
	headerParams map[string]string
	queryParams  url.Values
}

func NewRequest(path, method string, body any) *Request {
	ret := &Request{
		path:         path,
		method:       method,
		body:         body,
		headerParams: make(map[string]string),
		queryParams:  make(map[string][]string),
	}
	ret.headerParams["Content-Type"] = "application/json"
	ret.headerParams["Accept"] = "application/json, application/problem+json"
	return ret
}

func (r *Request) AddParam(k, v string) {
	if values, ok := r.queryParams[k]; ok {
		r.queryParams[k] = append(values, k)
	} else {
		r.queryParams[k] = []string{v}
	}
}

func (r *Request) AddHeader(k, v string) {
	r.headerParams[k] = v
}

func NewReceivedResponse(code int, status string, contentLength int64, body io.Reader, headers map[string]string) *Response {
	return &Response{
		code:          code,
		status:        status,
		body:          body,
		contentLength: contentLength,
		headers:       headers,
	}
}

type Response struct {
	contentLength int64
	body          io.Reader
	bodyBytes     []byte

	headers map[string]string
	status  string
	code    int
}

func (resp *Response) GetCode() int {
	return resp.code
}

func (resp *Response) GetStatus() string {
	return resp.status
}

func (resp *Response) GetHeaders() map[string]string {
	return resp.headers
}

// Abstraction of a consumer client
type ConsumerClient interface {
	Send(*Request) (*Response, error)
}

// an abstraction of the context where a request is received at a producer. The
// first handler method (openapi/producers) will inject a correct expected
// body for decoding. The next handler (application layer) will process the
// decoded body then create a response. It then write it to a response writer
// to complete the whole procedure of handling a request.

type RequestContext interface {
	HasRequestBody() bool
	DecodeRequest(body interface{}) error //decode the request to get embeded body
	Param(string) string                  // get a parameter from the request (application handler need it)
	Header(string) string                 // get a header parameter from the request (application handler need it)
	WriteResponse(int, any)
}
