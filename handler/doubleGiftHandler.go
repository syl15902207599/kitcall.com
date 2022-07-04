package handler

import (
	"context"

	khttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	medp "kitcall.com/endpoint"
	t "kitcall.com/transport"
)

func makeDGetInfoHandler(c context.Context, r *mux.Router) *mux.Router {
	r.Methods("GET").Path("/get_info").Handler(khttp.NewServer(
		medp.MakeDoubleGiftGetInfoHandlers(c, &t.DGetInfo{}, medp.GET, "/activity/doublegift/info"),
		t.DGetInfoDecodeRequest,
		t.DGetInfoEncodeResponse,
	))
	r.Methods("GET").Path("/exchange/{idx:[0-9]}").Handler(khttp.NewServer(
		medp.MakeDoubleGiftGetInfoHandlers(c, &t.DExchange{}, medp.GET, "/activity/doublegift/exchange"),
		t.ExchangeDecodeRequest,
		t.ExchangeEncodeResponse,
	))

	return r
}
