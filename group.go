package gee

import "net/http"

type RouterGroup struct {
	basePath   string        // api分组前缀
	middleware []HandlerFunc // 中间间支持，这个是group最重要的一个功能
	engine     *Engine
}

func (rg *RouterGroup) Group(prefix string) *RouterGroup {
	g := &RouterGroup{
		basePath:   rg.engine.basePath + prefix,
		middleware: make([]HandlerFunc, 0),
		engine:     rg.engine,
	}
	rg.engine.groups = append(rg.engine.groups, g)
	return g
}

func (rg *RouterGroup) addRoute(method, path string, handler HandlerFunc) {
	fullPath := rg.basePath + path
	rg.engine.router.addRoute(method, fullPath, handler)
}

func (rg *RouterGroup) Get(path string, handler HandlerFunc) {
	rg.addRoute(http.MethodGet, path, handler)
}

func (rg *RouterGroup) Post(path string, handler HandlerFunc) {
	rg.addRoute(http.MethodPost, path, handler)
}

func (rg *RouterGroup) Delete(path string, handler HandlerFunc) {
	rg.addRoute(http.MethodDelete, path, handler)
}

func (rg *RouterGroup) Put(path string, handler HandlerFunc) {
	rg.addRoute(http.MethodPut, path, handler)
}

func (rg *RouterGroup) Option(path string, handler HandlerFunc) {
	rg.addRoute(http.MethodOptions, path, handler)
}
