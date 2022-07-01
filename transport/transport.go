package transport

import (
	"context"
	"net/http"
)

type ITransport interface {
	EncodeRequset(context.Context, *http.Request, interface{}) error
	DecodeRespose(context.Context, *http.Response) (interface{}, error)
}
