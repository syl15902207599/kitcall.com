package handler

import (
	"context"

	khttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	medp "kitcall.com/endpoint"
	"kitcall.com/transport"
)

func makeDGetInfoHandler(c context.Context, r *mux.Router) *mux.Router {
	g := transport.DGetInfo{}

	r.Methods("GET").Path("/get_info").Handler(khttp.NewServer(
		medp.MakeDoubleGiftGetInfoHandlers(c, &g),
		g.DecodeRequest,
		g.EncodeResponse,
	))

	return r
}
