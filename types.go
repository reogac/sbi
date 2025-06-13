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

type Response struct {
	contentLength int64
	body          io.ReadCloser
	headers       map[string]string
	status        string
	code          int
}

func (rsp *Response) GetCode() int {
	return rsp.code
}

func (rsp *Response) GetStatus() string {
	return rsp.status
}

func (rsp *Response) GetHeaders() map[string]string {
	return rsp.headers
}

func (rsp *Response) DecodeBody(ie SbiIE) (err error) {
	return Decode(rsp.contentLength, rsp.body, ie)
}

func (rsp *Response) CloseBody() {
	rsp.body.Close()
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
	RequestBody() (int64, io.Reader)
	Param(string) string  // get a parameter from the request (application handler need it)
	Header(string) string // get a header parameter from the request (application handler need it)
	WriteResponse(int, SbiIE, map[string]string)
}
