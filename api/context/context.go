package context

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/megamsys/libgo/errors"
	"github.com/megamsys/megdcui/auth"
)

const (
	tokenContextKey int = iota
	errorContextKey
	delayedHandlerKey
)

func Clear(r *http.Request) {
	context.Clear(r)
}

func GetAuthToken(r *http.Request) auth.Token {
	if v := context.Get(r, tokenContextKey); v != nil {
		return v.(auth.Token)
	}
	return nil
}

func SetAuthToken(r *http.Request, t auth.Token) {
	context.Set(r, tokenContextKey, t)
}

func AddRequestError(r *http.Request, err error) {
	if err == nil {
		return
	}
	existingErr := context.Get(r, errorContextKey)
	if existingErr != nil {
		err = &errors.CompositeError{Base: existingErr.(error), Message: err.Error()}
	}
	context.Set(r, errorContextKey, err)
}

func GetRequestError(r *http.Request) error {
	if v := context.Get(r, errorContextKey); v != nil {
		return v.(error)
	}
	return nil
}

func SetDelayedHandler(r *http.Request, h http.Handler) {
	context.Set(r, delayedHandlerKey, h)
}

func GetDelayedHandler(r *http.Request) http.Handler {
	v := context.Get(r, delayedHandlerKey)
	if v != nil {
		return v.(http.Handler)
	}
	return nil
}
