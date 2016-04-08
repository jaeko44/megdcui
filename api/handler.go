package api

import (
	"net/http"

	"github.com/megamsys/megdcui/api/context"
)

type Handler func(http.ResponseWriter, *http.Request) error

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context.AddRequestError(r, fn(w, r))
}
