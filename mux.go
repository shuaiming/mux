package mux

import "net/http"

// Mux route http request
type Mux struct {
	*http.ServeMux
}

// New HTTP Servermux
func New(m *http.ServeMux) *Mux {
	return &Mux{m}
}

// ServeHTTP implement pod.Handler
func (m *Mux) ServeHTTP(
	w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	m.ServeMux.ServeHTTP(w, r)
	next(w, r)
}
