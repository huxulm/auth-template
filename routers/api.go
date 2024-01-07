package routers

import (
	"encoding/gob"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/huxulm/auth-template/models"
	UserService "github.com/huxulm/auth-template/services"
)

var sid = "__sid"
var store sessions.Store

func init() {

}

func SetStore(s sessions.Store) {
	gob.Register(&models.User{})
	store = s
}

func SidName(name string) {
	sid = name
}

func AuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, sid)
		if err != nil {
			writeHeader(w, http.StatusUnauthorized)
			return
		}
		if v, ok := session.Values["user"]; !ok {
			writeHeader(w, http.StatusUnauthorized)
			return
		} else {
			log.Printf("session: %#v\n", v)
			c, _ := r.Cookie(sid)
			newC := &http.Cookie{
				Name:    c.Name,
				Value:   c.Value,
				Domain:  c.Domain,
				Path:    c.Path,
				Expires: time.Now().Add(time.Duration(session.Options.MaxAge) * time.Second),
			}
			http.SetCookie(w, newC)
			// continue
			f(w, r)
		}
	}
}

func writeHeader(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

type H map[string]interface{}

// Validate user login credentials
// if ok then add `user` to session values
var Login http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeHeader(w, http.StatusNotFound)
		return
	}
	name, pass := r.FormValue("name"), r.FormValue("password")
	if u, err := UserService.Authenticate(name, pass); err != nil {
		// authenticate failed
		writeHeader(w, http.StatusUnauthorized)
	} else {
		session, _ := store.Get(r, sid)
		session.Values["user"] = u
		session.Options.MaxAge = 3600
		sessions.Save(r, w)
		// write data to response
		json.NewEncoder(w).Encode(H{"user": u})
	}
}

func Me(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, sid)
	if u, ok := session.Values["user"]; ok {
		// write data to response
		json.NewEncoder(w).Encode(H{"user": u})
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, sid)
	session.Options.MaxAge = -1 // mark delete
	session.Save(r, w)
	json.NewEncoder(w).Encode(H{"logout": "ok"})
}
