package transport

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	khttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"kitcall.com/util"
)

//获取基本信息
type DGetInfo struct{}

type GetInfoResponse struct {
	GoldNum int `json:"gold_num"`
}

func (s *DGetInfo) DecodeResponse(c context.Context, w *http.Response) (response interface{}, err error) {
	res := GetInfoResponse{}
	err = json.NewDecoder(w.Body).Decode(&res)
	if err != nil {
		return util.Response(true, err.Error(), nil)
	}

	return util.Response(false, "success", res)
}

func (s *DGetInfo) EncodeRequset(c context.Context, r *http.Request, req interface{}) error {
	setHeader(r)
	return nil
}

func DGetInfoDecodeRequest(c context.Context, w *http.Request) (request interface{}, err error) {
	return nil, nil
}

func DGetInfoEncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

//兑换物品
type DExchange struct{}

type ExchangeRequest struct {
	Index int `json:"idx"`
}

type ExchangeResponse struct {
	Gotten int `json:"gotten"`
}

func (s *DExchange) DecodeRespose(c context.Context, w *http.Response) (response interface{}, err error) {
	res := ExchangeRequest{}
	err = json.NewDecoder(w.Body).Decode(&res)
	if err != nil {
		return util.Response(true, err.Error(), nil)
	}
	response = res
	return util.Response(false, "success", res)
}

func (s *DExchange) EncodeRequset(c context.Context, r *http.Request, req interface{}) error {
	request := req.(ExchangeRequest)
	q := r.URL.Query()
	q.Set("idx", strconv.Itoa(request.Index))
	r.URL.RawQuery = q.Encode()
	setHeader(r)
	return nil
}

func ExchangeDecodeRequest(c context.Context, r *http.Request) (request interface{}, err error) {
	req := mux.Vars(r)
	if _, ok := req["idx"]; ok {
		idx, _ := strconv.Atoi(req["idx"])
		return ExchangeRequest{Index: idx}, nil
	}
	return ExchangeRequest{}, errors.New("参数错误")
}

func ExchangeEncodeResponse(c context.Context, w http.ResponseWriter, respose interface{}) error {
	return json.NewEncoder(w).Encode(respose)
}

func setHeader(r *http.Request) {
	r.Header.Add("HTTP_GAME_DB", "pirate90030001")
}

func HttpArithmeticFactory(_ context.Context, tsp IHTTTPTransport, method, path string) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, err error) {
		if !strings.HasPrefix(instance, "http") {
			instance = "http://" + instance
		}

		tgt, err := url.Parse(instance)
		if err != nil {
			return nil, nil, err
		}
		tgt.Path = path
		return khttp.NewClient(method, tgt, tsp.EncodeRequset, tsp.DecodeResponse).Endpoint(), nil, nil
	}
}
