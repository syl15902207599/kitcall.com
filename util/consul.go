package util

import (
	"encoding/json"
	"errors"
	"flag"
	"os"
	"strconv"

	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/log"
	"github.com/hashicorp/consul/api"
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
var port = flag.Int("port", 8500, "Input Service Port")
var HPort = flag.Int("hport", 8080, "Input Service Port")
var LogC = make(chan *[]interface{}, 100)

func init() {
	initConsulClient()
}

type CommonResposne struct {
	Err  bool        `json:"err"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(err bool, s string, data interface{}) (interface{}, error) {
	r := CommonResposne{Err: err, Msg: s, Data: data}
	if err {
		j, _ := json.Marshal(r)
		return data, errors.New(string(j))
	}
	return r, nil
}

func initConsulClient() {
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

func initLog() {
	f, err := os.OpenFile("./logs/logFile/go-kit.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("log file open fataled")
	}
	Logger = log.NewLogfmtLogger(f)
	Logger = log.With(Logger, "time", log.DefaultTimestamp)
	Logger = log.With(Logger, "caller", log.DefaultCaller)

}
