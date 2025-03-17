package rush

import (
	"encoding/json"
	"encoding/xml"
)

type Request struct {
	body  []byte
	error error
}

func (request *Request) Bytes() []byte {
	return request.body
}

func (request *Request) JSON(v interface{}) error {
	if request.error != nil {
		return request.error
	}
	return json.Unmarshal(request.body, v)
}

func (request *Request) XML(v interface{}) error {
	if request.error != nil {
		return request.error
	}
	return xml.Unmarshal(request.body, v)
}
