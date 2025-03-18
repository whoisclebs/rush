package rush

type Context struct {
	Request  *Request
	Response *Response
}

type HandlerFunc func(Request *Request, Response *Response) error
