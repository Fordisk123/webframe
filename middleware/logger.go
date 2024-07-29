package middleware

import (
	"context"
	"fmt"
	logger "github.com/fordisk/webframe/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
	"reflect"
	"time"
)

func LoggingMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		startTime := time.Now()
		log := logger.WithFields(ctx, "Middleware-Log", "")
		logCtx := logger.WithContext(ctx, log)
		if httpCtx, ok := ctx.(http.Context); ok {
			ctx = context.WithValue(logCtx, "http.request", httpCtx.Request())
			logger.WithFields(logCtx, "path", httpCtx.Request().URL.Path, "method", httpCtx.Request().Method)
			reply, err = handler(ctx, req)
		} else {
			reply, err = handler(logCtx, req)
		}
		duration := time.Since(startTime)
		//接口耗时
		log = logger.WithFields(logCtx, "latency", fmt.Sprintf("%dms", duration.Milliseconds()))
		if err != nil {
			//replay 反射 RtnCode值来判定打info还是error
			replayVal := reflect.ValueOf(err)
			rtnCodeFiled := replayVal.Elem().FieldByName("RtnCode")
			//认为是未知错误
			if rtnCodeFiled.String() == "<invalid Value>" {
				log.Errorf("请求调用出现未知错误:%s", err.Error())
			} else {
				if rtnCodeFiled.String() == "000000" {
					log.Infof("请求调用成功")
				} else {
					log.Errorf("请求调用失败," + err.Error())
				}
			}
		} else {
			log.Infof("请求调用成功")
		}
		return
	}
}
