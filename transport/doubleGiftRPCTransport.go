package transport

import (
	"context"
	"io"

	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	kgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	pb "kitcall.com/pbs"
	"kitcall.com/util"
)

type DGetInfoRpc struct{}

func (d *DGetInfoRpc) EncodeRequset(c context.Context, req interface{}) (interface{}, error) {
	return nil, nil
}

func (d *DGetInfoRpc) DecodeResponse(c context.Context, res interface{}) (interface{}, error) {
	r, ok := res.(*pb.GetInfoRes)
	if !ok {
		return util.Response(true, "getInfo.GetInfoRes：解析错误", nil)
	}
	return util.Response(false, "success", GetInfoResponse{GoldNum: int(r.GoldNum)})
}

type DExchangeRpc struct{}

func (d *DExchangeRpc) EncodeRequset(c context.Context, req interface{}) (interface{}, error) {
	r, ok := req.(*pb.ExchangeReq)
	if !ok {
		return nil, errors.New("exchange.ExchangeReq：解析错误")
	}
	return &pb.ExchangeReq{Idx: r.Idx}, nil
}

func (d *DExchangeRpc) DecodeResponse(c context.Context, res interface{}) (interface{}, error) {
	r, ok := res.(*pb.ExchangeRes)
	if !ok {
		return util.Response(true, "exchange.ExchangeReq：解析错误", nil)
	}
	return ExchangeResponse{Gotten: int(r.Gotten)}, nil
}

func RpcArithmeticFactory(_ context.Context, tsp IRPCTransport, serviceName, method string, response interface{}) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, err error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		return kgrpc.NewClient(conn, serviceName, method, tsp.EncodeRequset, tsp.DecodeResponse, response).Endpoint(), conn, nil
	}
}
