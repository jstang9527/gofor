package handler

import (
	"context"

	"net/url"

	"github.com/jstang9527/gofor/src/share/config"
	"github.com/jstang9527/gofor/src/share/log"
	"github.com/jstang9527/gofor/src/share/pb"
	"github.com/jstang9527/gofor/src/srv-webshell/entity"
	"github.com/micro/go-micro/errors"
	"go.uber.org/zap"
)

type WebshellServiceExtHandler struct {
	logger *zap.Logger
}

func NewWebshellServiceExtHandler() *WebshellServiceExtHandler {
	return &WebshellServiceExtHandler{
		logger: log.Instance(),
	}
}

func (w *WebshellServiceExtHandler) RunCmdWithOutput(ctx context.Context, req *pb.RunCmdWithOutputReq, rsp *pb.RunCmdWithOutputRsp) error {
	// 1. url正规性
	if _, err := url.Parse(req.Target); err != nil {
		w.logger.Error("error", zap.Error(err))
		return errors.New(config.ServiceNameUser, "invalid url", 200)
	}
	// 2. 获取对应的实体
	obj := entity.LoadLanguageEntity(entity.PHP_ENV, req.Target, req.Cmd)
	if obj == nil {
		err := errors.New(config.ServiceNameUser, "unsupport lang", 200)
		w.logger.Error("error", zap.Error(err))
		return err
	}
	// 3. 执行命令并返回
	out, err := obj.RunCmdWithOutput(req.Cmd)
	if err != nil {
		w.logger.Error("error", zap.Error(err))
		return errors.New(config.ServiceNameUser, "failed operation", 200)
	}
	rsp.Output = out
	return nil
}
