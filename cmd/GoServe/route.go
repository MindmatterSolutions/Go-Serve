package GoServe

import (
	"net/http"
)

type Route struct {
	handler http.Handler
}
