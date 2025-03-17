package rush

import "net/http"

type Router struct {
	mux *http.ServeMux
	app *App
}

func (r *Router) handle(method, pattern string, handler HandlerFunc) {
	r.mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		ctx := &Context{writer: w, Request: req}
		if req.Method != method {
			r.app.ErrorHandler(ctx, ErrorRush{Code: 405, Message: "Method Not Allowed"})
			return
		}
		if err := handler(ctx); err != nil {
			r.app.ErrorHandler(ctx, err)
			return
		}
	})
}

func (r *Router) GET(pattern string, handler HandlerFunc) {
	r.handle(http.MethodGet, pattern, handler)
}

func (r *Router) POST(pattern string, handler HandlerFunc) {
	r.handle(http.MethodPost, pattern, handler)
}

func (r *Router) PUT(pattern string, handler HandlerFunc) {
	r.handle(http.MethodPut, pattern, handler)
}

func (r *Router) DELETE(pattern string, handler HandlerFunc) {
	r.handle(http.MethodDelete, pattern, handler)
}
