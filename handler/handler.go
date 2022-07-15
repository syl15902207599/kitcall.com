package handler

import (
	"context"

	"github.com/gorilla/mux"
)

func MakeHttpHandlers(c context.Context) *mux.Router {
	r := mux.NewRouter()
	r = makeDoubleGiftHandler(c, r)
	return r
}
