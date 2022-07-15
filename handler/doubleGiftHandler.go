package handler

import (
	"context"

	khttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	medp "kitcall.com/endpoint"
	"kitcall.com/logs"
	pb "kitcall.com/pbs"
	t "kitcall.com/transport"
	"kitcall.com/util"
)

func makeDoubleGiftHandler(c context.Context, r *mux.Router) *mux.Router {
	h := logs.HttpLogMiddleware(util.Logger)
	hg := h(&t.DGetInfo{})
	he := h(&t.DGetInfo{})
	r.Methods("GET").Path("/get_info").Handler(khttp.NewServer(
		medp.MakeDoubleGiftHandlers(c, &hg, medp.GET, "/activity/doublegift/info"),
		t.DGetInfoDecodeRequest,
		t.DGetInfoEncodeResponse,
	))
	r.Methods("GET").Path("/exchange/{idx:[0-9]}").Handler(khttp.NewServer(
		medp.MakeDoubleGiftHandlers(c, &he, medp.GET, "/activity/doublegift/exchange"),
		t.ExchangeDecodeRequest,
		t.ExchangeEncodeResponse,
	))
	rpc := logs.RpcLogMiddleware(util.Logger)
	rg := rpc(&t.DGetInfoRpc{})
	re := rpc(&t.DExchangeRpc{})
	r.Methods("GET").Path("/rpc/get_info").Handler(khttp.NewServer(
		medp.MakeRPCDoubleGiftHandlers(c, &rg, "message.GetInfo", "GetInfoRpc", pb.GetInfoRes{}),
		t.DGetInfoDecodeRequest,
		t.DGetInfoEncodeResponse,
	))
	r.Methods("GET").Path("/rpc/exchange/{idx:[0-9]}").Handler(khttp.NewServer(
		medp.MakeRPCDoubleGiftHandlers(c, &re, "message.Exchange", "ExchangeRpc", pb.ExchangeRes{}),
		t.ExchangeDecodeRequest,
		t.ExchangeEncodeResponse,
	))

	return r
}
