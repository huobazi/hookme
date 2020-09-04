package routes

import (
	"net/http"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

type MethodCollection []string

type Router interface {
	AddRoute(path string, methods MethodCollection, handler http.HandlerFunc) Router
	GetParam(r *http.Request, index int) string
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

var defaultRouter = NewRegexRouter()

func AddRoute(path string, methods MethodCollection, handler http.HandlerFunc) Router {
	return defaultRouter.AddRoute(path, methods, handler)
}

func GetParam(r *http.Request, index int) string {
	return defaultRouter.GetParam(r, index)
}
