package routes

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

type RegexRoute struct {
	methods MethodCollection
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

type RegexRouter struct {
	Tables []RegexRoute
}

func NewRegexRouter() Router {
	return &RegexRouter{}
}

func (router *RegexRouter) AddRoute(path string, methods MethodCollection, handler http.HandlerFunc) Router {
	router.Tables = append(router.Tables,
		RegexRoute{
			methods: methods,
			regex:   regexp.MustCompile("^" + path + "$"),
			handler: handler,
		},
	)

	return router
}

type contextKey struct{}

func (router *RegexRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.Tables {
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

func (router *RegexRouter) GetParam(r *http.Request, index int) string {
	fields := r.Context().Value(contextKey{}).([]string)
	return fields[index]
}