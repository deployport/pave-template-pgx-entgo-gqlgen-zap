package web

import (
	"fmt"
	"net/http"
)

func homeRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
		<h2>mainprocess server!</h2>
		<div>
			GraphQL Endpoint at <a href="/graphql">/graphql</a>
		</div>
	`)
}
