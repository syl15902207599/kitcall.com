package transport

import (
	"context"
	"net/http"
)

type IHTTTPTransport interface {
	EncodeRequset(context.Context, *http.Request, interface{}) error
	DecodeResponse(context.Context, *http.Response) (interface{}, error)
}

type IRPCTransport interface {
	EncodeRequset(context.Context, interface{}) (request interface{}, err error)
	DecodeResponse(context.Context, interface{}) (response interface{}, err error)
}
