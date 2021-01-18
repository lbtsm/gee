package gee

import "net/http"

type HandlerFunc func(ctx *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) Get(path string, handler HandlerFunc) {
	e.router.addRoute(http.MethodGet, path, handler)
}

func (e *Engine) Post(path string, handler HandlerFunc) {
	e.router.addRoute(http.MethodPost, path, handler)
}

func (e *Engine) Delete(path string, handler HandlerFunc) {
	e.router.addRoute(http.MethodDelete, path, handler)
}

func (e *Engine) Put(path string, handler HandlerFunc) {
	e.router.addRoute(http.MethodPut, path, handler)
}

func (e *Engine) Option(path string, handler HandlerFunc) {
	e.router.addRoute(http.MethodOptions, path, handler)
}

func (e *Engine) Run(port string) error {
	return http.ListenAndServe(port, e)
}

/*
重写 ServeHTTP 的必要性，其实也是写这个框架的原因
为什么要写这个框架，那肯定是go src内的http server，有什么缺陷
*/
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.handle(c)
}

/*
	MethodHead    = "HEAD"
	MethodPatch   = "PATCH" // RFC 5789
	MethodConnect = "CONNECT"
	MethodTrace   = "TRACE"
*/
