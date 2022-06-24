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

func MakeDoubleGiftGetInfoHandlers(c context.Context, tsp transport.ITransport) endpoint.Endpoint {
	instancer := consul.NewInstancer(util.DisClient.Client, util.Logger, "doubleGiftService", []string{"default"}, false)
	factory := util.ArithmeticFactory(c, tsp, "GET", "/activity/doublegift/info")

	//创建端点管理器， 此管理器根据Factory和监听的到实例创建endPoint并订阅instancer的变化动态更新Factory创建的endPoint
	endpointer := sd.NewEndpointer(instancer, factory, util.Logger)
	//创建RoundRibbon负载均衡器
	balancer := lb.NewRoundRobin(endpointer)
	//为负载均衡器增加重试功能，同时该对象为endpoint.Endpoint
	return lb.Retry(5, 3*time.Second, balancer)
}
