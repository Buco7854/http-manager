package router

import (
	"github.com/Buco7854/http-shutdown/errors"
	"net/http"
)

type Router struct {
	Routes []Route
}

func (router *Router) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	for _, route := range router.Routes {
		pathMatch, methodMatch := route.isMatching(req)
		if pathMatch {
			if methodMatch {
				route.Handler(writer, req)
			} else {
				errors.JSONError(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
			return
		}
	}
	errors.JSONError(writer, "Not Found", http.StatusNotFound)
}

func (router *Router) AddRoute(method string, path string, handler http.HandlerFunc) {
	route := Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
	router.Routes = append(router.Routes, route)
}
