package daemon

import (
	"fmt"
	"net/http"
)

func HandleHello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello from the server")
}
