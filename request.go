package rush

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
)

type Request struct {
	http *http.Request
	body *Body
}

type Body struct {
	bytes []byte
	error error
}

func (request *Request) Headers() map[string][]string {
	return request.http.Header
}

func (request *Request) Body() *Body {
	if request.body.bytes != nil {
		return &Body{bytes: request.body.bytes}
	}
	data, err := io.ReadAll(request.http.Body)
	if err != nil {
		return &Body{error: err}
	}
	request.body = &Body{bytes: data}
	return request.body
}

func (body *Body) Bytes() []byte {
	return body.bytes
}

func (body *Body) JSON(v interface{}) error {
	if body.error != nil {
		return body.error
	}
	return json.Unmarshal(body.Bytes(), v)
}

func (body *Body) XML(v interface{}) error {
	if body.error != nil {
		return body.error
	}
	return xml.Unmarshal(body.Bytes(), v)
}
