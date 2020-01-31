package photo

import (
	"fmt"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	fmt.Fprintf(w, "{}")
}
