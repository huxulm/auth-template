package routers

import (
	"net/http"
	"strings"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/v1") {
			switch r.URL.Path {
			case "/v1/login":
				Login(w, r)
			case "/v1/me":
				Me(w, r)
			case "/v1/logout":
				Logout(w, r)
			}
		} else {
			if r.URL.Path == "/dashboard" || r.URL.Path == "/login" {
				r.URL.Path += ".html"
			}
			serveWeb()(w, r)
		}
	})
}
