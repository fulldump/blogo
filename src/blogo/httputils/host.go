package httputils

import "net/http"

func GetHost(r *http.Request) string {

	h := r.Header.Get("X-Forwarded-Host")

	if "" != h {
		return h
	}

	return r.Host
}
