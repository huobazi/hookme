# routes

```go

func hello(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "Hookme server is starting ...\n")
}

func helloByName(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello %s \n", routes.GetParam(r, 0))
}

router := routes.
		AddRoute("/hello", routes.MethodCollection{routes.GET}, hello).
		AddRoute("/hello/([^/]+)", routes.MethodCollection{routes.GET}, helloByName)

http.ListenAndServe("0.0.0.0:7979",router)

```