package transport

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetInfoResponse struct {
	GoldNum int `json:"gold_num"`
}
type DGetInfo struct{}

func (_ *DGetInfo) DecodeRespose(c context.Context, w *http.Response) (response interface{}, err error) {
	res := GetInfoResponse{}
	err = json.NewDecoder(w.Body).Decode(&res)
	if err != nil {
		return
	}
	response = res
	return
}

func (_ *DGetInfo) EncodeRequset(c context.Context, w *http.Request, req interface{}) error {
	return nil
}

func (_ *DGetInfo) DecodeRequest(c context.Context, w *http.Request) (request interface{}, err error) {
	return nil, nil
}

func (_ *DGetInfo) EncodeResponse(c context.Context, w http.ResponseWriter, respose interface{}) error {
	return json.NewEncoder(w).Encode(respose)
}
