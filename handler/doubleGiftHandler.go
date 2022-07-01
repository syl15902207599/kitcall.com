package handler

import (
	"context"
	"fmt"

	khttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	medp "kitcall.com/endpoint"
	t "kitcall.com/transport"
)

func makeDGetInfoHandler(c context.Context, r *mux.Router) *mux.Router {
	fmt.Println(1)
	fmt.Println(medp.MakeDoubleGiftGetInfoHandlers(c, &t.DGetInfo{}, medp.GET, "/activity/doublegift/info"))
	r.Methods("GET").Path("/get_info").Handler(khttp.NewServer(
		medp.MakeDoubleGiftGetInfoHandlers(c, &t.DGetInfo{}, medp.GET, "/activity/doublegift/info"),
		t.DGetInfoDecodeRequest,
		t.DGetInfoEncodeResponse,
	))
	fmt.Println(2)
	fmt.Println(medp.MakeDoubleGiftGetInfoHandlers(c, &t.DExchange{}, medp.GET, "/activity/doublegift/exchange"))
	r.Methods("GET").Path("/exchange/{idx:[0-9]}").Handler(khttp.NewServer(
		medp.MakeDoubleGiftGetInfoHandlers(c, &t.DExchange{}, medp.GET, "/activity/doublegift/exchange"),
		t.ExchangeDecodeRequest,
		t.ExchangeEncodeResponse,
	))

	// r.Methods("GET").Path("/exchange").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("1111111"))
	// })
	return r
}
