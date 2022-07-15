package endpoint

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	"kitcall.com/transport"
	"kitcall.com/util"
)

const (
	GET = iota
	POST
	PUT
	DELETE
)

var methods = [4]string{"GET", "POST", "PUT", "DELETE"}

func MakeDoubleGiftHandlers(c context.Context, tsp transport.IHTTTPTransport, method int, path string) endpoint.Endpoint {
	instancer := consul.NewInstancer(util.DisClient.Client, util.Logger, "doubleGiftService", []string{"default"}, false)
	factory := transport.HttpArithmeticFactory(c, tsp, methods[method], path)

	//创建端点管理器， 此管理器根据Factory和监听的到实例创建endPoint并订阅instancer的变化动态更新Factory创建的endPoint
	endpointer := sd.NewEndpointer(instancer, factory, util.Logger)
	//创建RoundRibbon负载均衡器
	balancer := lb.NewRoundRobin(endpointer)
	//为负载均衡器增加重试功能，同时该对象为endpoint.Endpoint
	return lb.Retry(1, 3*time.Second, balancer)
}

func MakeRPCDoubleGiftHandlers(c context.Context, tsp transport.IRPCTransport, rpcService, method string, response interface{}) endpoint.Endpoint {
	instancer := consul.NewInstancer(util.DisClient.Client, util.Logger, "doubleGiftService", []string{"rpc"}, false)
	factory := transport.RpcArithmeticFactory(c, tsp, rpcService, method, response)

	//创建端点管理器， 此管理器根据Factory和监听的到实例创建endPoint并订阅instancer的变化动态更新Factory创建的endPoint
	endpointer := sd.NewEndpointer(instancer, factory, util.Logger)
	//创建RoundRibbon负载均衡器
	balancer := lb.NewRoundRobin(endpointer)
	//为负载均衡器增加重试功能，同时该对象为endpoint.Endpoint
	return lb.Retry(1, 3*time.Second, balancer)
}
