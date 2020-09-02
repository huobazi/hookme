package routes

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

type MethodCollection []string

type route struct {
	methods []string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

type routeTable []route

type router struct {
	routes routeTable
}

var Router = &router{}

func (router *router) AddRoute(path string, methods MethodCollection, handler http.HandlerFunc) *router {
	router.routes = append(router.routes,
		route{
			methods: methods,
			regex:   regexp.MustCompile("^" + path + "$"),
			handler: handler,
		},
	)

	return router
}

type contextKey struct{}

func (table routeTable) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range table {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			exists := false
			for _, m := range route.methods {
				if r.Method == m {
					exists = true
				}
			}
			if !exists {
				w.Header().Set("Allow", strings.Join(route.methods, ", "))
				http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
				return
			} else {
				ctx := context.WithValue(r.Context(), contextKey{}, matches[1:])
				route.handler(w, r.WithContext(ctx))
				return
			}
		}
	}

	http.NotFound(w, r)
}

func (router *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.routes.ServeHTTP(w, r)
}

func GetParam(r *http.Request, index int) string {
	fields := r.Context().Value(contextKey{}).([]string)
	return fields[index]
}
