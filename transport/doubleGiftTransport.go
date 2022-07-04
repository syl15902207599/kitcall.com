package transport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CommonResposne struct {
	Err  bool        `json:"err"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//获取基本信息
type DGetInfo struct{}

type GetInfoResponse struct {
	GoldNum int `json:"gold_num"`
}

func (s *DGetInfo) DecodeRespose(c context.Context, w *http.Response) (response interface{}, err error) {
	res := CommonResposne{Data: GetInfoResponse{}}
	err = json.NewDecoder(w.Body).Decode(&res)
	if err != nil {
		return
	}
	response = res
	return
}

func (s *DGetInfo) EncodeRequset(c context.Context, w *http.Request, req interface{}) error {
	return nil
}

func DGetInfoDecodeRequest(c context.Context, w *http.Request) (request interface{}, err error) {
	return nil, nil
}

func DGetInfoEncodeResponse(c context.Context, w http.ResponseWriter, respose interface{}) error {
	return json.NewEncoder(w).Encode(respose)
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
	res := CommonResposne{Data: ExchangeRequest{}}
	err = json.NewDecoder(w.Body).Decode(&res)
	if err != nil {
		return
	}
	response = res
	return
}

func (s *DExchange) EncodeRequset(c context.Context, r *http.Request, req interface{}) error {
	request := req.(ExchangeRequest)
	r.Form = r.URL.Query()
	r.Form.Add("idx", strconv.Itoa(request.Index))
	fmt.Println(r)

	return nil
}

func ExchangeDecodeRequest(c context.Context, r *http.Request) (request interface{}, err error) {
	req := mux.Vars(r)
	fmt.Println(req)
	if _, ok := req["idx"]; ok {
		idx, _ := strconv.Atoi(req["idx"])
		return ExchangeRequest{Index: idx}, nil
	}
	return ExchangeRequest{}, errors.New("参数错误")
}

func ExchangeEncodeResponse(c context.Context, w http.ResponseWriter, respose interface{}) error {
	return json.NewEncoder(w).Encode(respose)
}
