package router

import "net/http"

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func (route Route) isMatching(req *http.Request) (bool, bool) {
	return req.URL.Path == route.Path, req.Method == route.Method
}
