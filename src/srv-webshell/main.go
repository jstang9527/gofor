package main

import (
	"github.com/jstang9527/gofor/src/share/config"
	"github.com/jstang9527/gofor/src/share/log"
	"github.com/jstang9527/gofor/src/share/pb"
	"github.com/jstang9527/gofor/src/srv-webshell/handler"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"go.uber.org/zap"
)

func main() {
	log.Init("webshell", false)
	// consulReg := consul.NewRegistry(registry.Addrs(config.ConsulAddr))
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameWebshell),
		micro.Version("latest"),
		// micro.Registry(consulReg),
	)
	// 定义Service动作操作
	service.Init(
		micro.Action(func(c *cli.Context) {
			logger.Info("Info", zap.Any("webshell-srv", "webshell-srv is start ..."))
			// dao.Init(config.MysqlDSN)
			pb.RegisterWebshellServiceExtHandler(service.Server(), handler.NewWebshellServiceExtHandler(), server.InternalHandler(true))
		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("webshell-srv", "webshell-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			logger.Info("Info", zap.Any("webshell-srv", "webshell-srv is AfterStart ..."))
			return nil
		}),
	)
	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("webshell-srv服务启动失败 ...")
	}
}

// [*] Example:
// curl
// -H "Content-Type:application/json;charset=utf-8"
// -X POST
// -d '{"target":"http://172.31.50.248:8080/ant.php","language":"php","cmd":"ls"}' http://172.31.50.249:8888/webshell/WebshellServiceExt/RunCmdWithOutput
