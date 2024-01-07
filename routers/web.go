package routers

import (
	"net/http"
	"time"

	"github.com/huxulm/auth-template/assets"
)

func serveWeb() http.HandlerFunc {
	var fs, _ = assets.Assets()
	h := http.FileServer(http.FS(fs))

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// set cache headers
			w.Header().Add("Cache-Control", "max-age=84600, public")
			w.Header().Set("Expires", time.Now().Add(time.Hour*24).Format(http.TimeFormat))
			w.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
