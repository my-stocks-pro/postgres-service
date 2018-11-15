package router

import "net/http"

func (r TypeRouter) HandlerVersion(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("HandlerVersion"))
}
