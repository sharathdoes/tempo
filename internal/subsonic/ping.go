package subsonic

import "net/http"



func Ping(w http.ResponseWriter, r *http.Request) {
	WriteResult(w, "ok")
}


