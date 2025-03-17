package rush

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
)

type Context struct {
	Request   *http.Request
	Response  *Response
	bodyCache *Request
	writer    http.ResponseWriter
}

type HandlerFunc func(*Context) error

func (ctx *Context) Status(code int) *Context {
	ctx.Response.statusCode = code
	return ctx
}

func (ctx *Context) JSON(obj interface{}) error {
	ctx.writer.Header().Set("Content-Type", "application/json")
	if ctx.Response.statusCode != 0 {
		ctx.writer.WriteHeader(ctx.Response.statusCode)
	}
	return json.NewEncoder(ctx.writer).Encode(obj)
}

func (ctx *Context) XML(obj interface{}) error {
	ctx.writer.Header().Set("Content-Type", "application/xml")
	if ctx.Response.statusCode != 0 {
		ctx.writer.WriteHeader(ctx.Response.statusCode)
	}
	return xml.NewEncoder(ctx.writer).Encode(obj)
}

func (ctx *Context) Body() *Request {
	if ctx.bodyCache != nil {
		return ctx.bodyCache
	}
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return &Request{error: err}
	}
	ctx.bodyCache = &Request{body: data}
	return ctx.bodyCache
}
