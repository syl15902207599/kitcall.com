package util

import (
	"context"
	"flag"
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	khttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/hashicorp/consul/api"
	"kitcall.com/transport"
)

type DiscoverClient struct {
	Address string
	Port    int
	Config  *api.Config
	Client  consul.Client
}

var DisClient *DiscoverClient
var Logger = log.NewNopLogger()
var address = flag.String("address", "127.0.0.1", "Input Service Address")
var port = flag.Int("port", 8080, "Input Service Port")

func init() {
	flag.Parse()
	DisClient = &DiscoverClient{}
	DisClient.Address = *address
	DisClient.Port = *port
	//链接注册中心
	DisClient.Config = api.DefaultConfig()
	DisClient.Config.Address = "http://" + DisClient.Address + ":" + strconv.Itoa(DisClient.Port)
	client, err := api.NewClient(DisClient.Config)
	if err != nil {
		panic("url error : " + err.Error())
	}
	DisClient.Client = consul.NewClient(client)
}

func ArithmeticFactory(_ context.Context, tsp transport.ITransport, method, path string) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, err error) {
		if !strings.HasPrefix(instance, "http") {
			instance = "http://" + instance
		}

		tgt, err := url.Parse(instance)
		if err != nil {
			return nil, nil, err
		}
		tgt.Path = path
		return khttp.NewClient(method, tgt, tsp.EncodeRequset, tsp.DecodeRespose).Endpoint(), nil, nil
	}
}
