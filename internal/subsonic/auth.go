package subsonic

import (
	"crypto/md5"
	"crypto/subtle"
	"fmt"
	"net/http"
	"strings"
)

//you first create a middleware with username and password
//which returns u handler which accepts a handler
//with this middleware, u pass the handler of the func u need authorization for
//that will become next, if all validations pass. next will be served.

func AuthMiddleWare(username, password string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !strings.HasPrefix(r.URL.Path, "/rest/") {
				next.ServeHTTP(w,r)
				return
			}
			//GET /rest/ping?u=admin&s=abc123&t=HASH&v=1.16.1&c=tempo&f=json HTTP/1.1
			//Host: localhost:8080
			query := r.URL.Query()
			user:=query.Get("u")
			salt:=query.Get("s")
			token:=query.Get("t")

			if user == "" || salt == "" || token == "" {
				WriteError(w, 10, "missing authentication params")
				return
			}

			if len(salt) < 6 {
				WriteError(w, 40, "invalid authentication")
				return
			}

			if user != username {
				WriteError(w, 40, "wrong username or pass")
				return
			}

			hash:=md5.Sum([]byte(password+salt))
			expected := fmt.Sprintf("%x", hash)
			if len(expected)!=len(token) || subtle.ConstantTimeCompare([]byte(token), []byte(expected)) != 1  {
				WriteError(w, 40, "Wrong username or password")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}