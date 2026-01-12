package daemon

import (
	"net/http"
	"rxplore/internals/api"
)

func StartServer() {
	http.HandleFunc(api.RouteHello, HandleHello)

	http.ListenAndServe("127.0.0.1:8080", nil)
}
