package handler

import (
	"context"

	"github.com/gorilla/mux"
)

func MakeHttpHandlers(c context.Context) *mux.Router {
	r := mux.NewRouter()
	r = makeDGetInfoHandler(c, r)
	return r
}
