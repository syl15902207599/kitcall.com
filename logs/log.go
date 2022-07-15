package logs

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/log"

	"kitcall.com/transport"
)

type loggingHttpMiddleware struct {
	logger     log.Logger
	HttpServer transport.IHTTTPTransport
}

type loggingRpcpMiddleware struct {
	logger    log.Logger
	RpcServer transport.IRPCTransport
}

type HttpServiceMiddleware func(transport.IHTTTPTransport) loggingHttpMiddleware

type RpcServiceMiddleware func(transport.IRPCTransport) loggingRpcpMiddleware

func HttpLogMiddleware(l log.Logger) HttpServiceMiddleware {
	return func(s transport.IHTTTPTransport) loggingHttpMiddleware {
		return loggingHttpMiddleware{logger: l, HttpServer: s}
	}
}

func RpcLogMiddleware(l log.Logger) RpcServiceMiddleware {
	return func(s transport.IRPCTransport) loggingRpcpMiddleware {
		return loggingRpcpMiddleware{logger: l, RpcServer: s}
	}
}

func (mw *loggingHttpMiddleware) EncodeRequset(c context.Context, r *http.Request, req interface{}) (err error) {
	defer func(beign time.Time) {
		mw.logger.Log(
			"client", "http",
			"function", "EncodeRequset",
			"result", err,
			"took", time.Since(beign),
		)
	}(time.Now())
	err = mw.HttpServer.EncodeRequset(c, r, req)
	return err
}

func (mw *loggingHttpMiddleware) DecodeResponse(c context.Context, w *http.Response) (response interface{}, err error) {
	defer func(beign time.Time) {
		mw.logger.Log(
			"client", "http",
			"function", "EncodeRequset",
			"result", err,
			"took", time.Since(beign),
		)
	}(time.Now())
	response, err = mw.HttpServer.DecodeResponse(c, w)
	return response, err
}

func (mw *loggingRpcpMiddleware) EncodeRequset(c context.Context, r interface{}) (request interface{}, err error) {
	defer func(beign time.Time) {
		mw.logger.Log(
			"client", "rpc",
			"function", "EncodeRequset",
			"result", err,
			"took", time.Since(beign),
		)
	}(time.Now())
	request, err = mw.RpcServer.EncodeRequset(c, r)
	return request, err
}

func (mw *loggingRpcpMiddleware) DecodeResponse(c context.Context, r interface{}) (response interface{}, err error) {
	defer func(beign time.Time) {
		mw.logger.Log(
			"client", "rpc",
			"function", "EncodeRequset",
			"result", err,
			"took", time.Since(beign),
		)
	}(time.Now())
	response, err = mw.RpcServer.DecodeResponse(c, r)
	return response, err
}
