package depot

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) {
	w.Header().Add("content-type", "application/json")
	v := r.URL.Query()


}
