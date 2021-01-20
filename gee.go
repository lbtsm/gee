package gee

import "net/http"

type HandlerFunc func(ctx *Context)

type Engine struct {
	*RouterGroup // 匿名继承
	router       *router
	groups       []*RouterGroup // store all routeGroup
}

func New() *Engine {
	e := &Engine{
		router: newRouter(),
		groups: make([]*RouterGroup, 0),
	}
	e.RouterGroup = &RouterGroup{
		basePath:   "/", // 这里减少使用group时，
		engine:     e,
		middleware: make([]HandlerFunc, 0),
	}
	return e
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
