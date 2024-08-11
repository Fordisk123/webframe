package main

import (
	"github.com/Fordisk123/webframe/conf"
	"github.com/Fordisk123/webframe/db"
	util "github.com/Fordisk123/webframe/frame"
	"github.com/Fordisk123/webframe/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "example"
	// Version is the version of the compiled software.
	Version = "v1.0.0"
)

func main() {
	conf.InitConf("./conf")

	log.NewDefaultLogger(Name, Version)

	db.InitDb()

	db.MigrateTable()

	util.KratosServe(Name, Version, func(httpSrv *http.Server) {

	}, nil, nil)
}
