package util

import (
	"github.com/Fordisk123/webframe/errors"
	"github.com/Fordisk123/webframe/log"
	"github.com/Fordisk123/webframe/middleware"
	"github.com/go-kratos/kratos/v2"
	kMidkdeware "github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/http/pprof"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/ory/viper"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

// 初始化kratos框架函数
// name 和部署的工作负载名称必须一样
// version 用于打印log时版本，跟随发行版本更改
func KratosServe(name, version string, regHook func(httpSrv *http.Server), errorHandler http.EncodeErrorFunc, httpCustomMiddlewares []kMidkdeware.Middleware) {
	log.NewLogger(&log.Config{
		AppName:      name,
		Env:          viper.GetString("run.mode"),
		LogDir:       "logs",
		MaxAge:       10,
		MaxLogFileMB: 10,
	}, "name", name, "version", version)

	//http端口默认8080
	httpPort := viper.GetString("http.port")
	if httpPort == "" {
		httpPort = "8080"
	}

	httpMiddewares := make([]kMidkdeware.Middleware, 0)
	httpMiddewares = append(httpMiddewares,
		recovery.Recovery(),
		middleware.LoggingMiddleware,
		middleware.MetricMiddleware(),
	)
	if httpCustomMiddlewares != nil && len(httpCustomMiddlewares) > 0 {
		httpMiddewares = append(httpMiddewares, httpCustomMiddlewares...)
	}

	if errorHandler == nil {
		errorHandler = errors.HttpErrorHandler
	}

	httpSrv := http.NewServer(
		http.Address(":"+httpPort),
		http.Middleware(httpMiddewares...),
		http.ErrorEncoder(errorHandler),
		http.Timeout(10*time.Minute),
	)
	//swagger
	h := openapiv2.NewHandler()
	// 将/q/路由放在最前匹配
	httpSrv.HandlePrefix("/q/", h)
	httpSrv.Handle("/metrics", promhttp.Handler())

	//grpcMiddewares := make([]kMidkdeware.Middleware, 0)
	//grpcMiddewares = append(grpcMiddewares,
	//	recovery.Recovery(),
	//	tracing.Server(),
	//	tracing.Client(),
	//	middleware.LoggingMiddleware,
	//	middleware.MetricMiddleware(),
	//	middleware.Validator(),
	//)
	//if grpcCustomMiddwares != nil && len(grpcCustomMiddwares) > 0 {
	//	grpcMiddewares = append(grpcMiddewares, grpcCustomMiddwares...)
	//}
	//
	//grpcSrv := grpc.NewServer(
	//	grpc.Address(":"+grpcPort),
	//	grpc.Middleware(grpcMiddewares...),
	//	grpc.Timeout(10*time.Minute),
	//)

	regHook(httpSrv)
	//pprof
	httpSrv.HandlePrefix("/", pprof.NewHandler())

	app := kratos.New(
		kratos.Name(name),
		kratos.Version(version),
		kratos.Server(
			httpSrv,
			//grpcSrv,
		),
		kratos.Logger(&log.StdKratosLog{}),
	)

	if err := app.Run(); err != nil {
		panic(err.Error())
	}
}
