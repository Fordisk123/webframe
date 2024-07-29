package main

import (
	"github.com/fordisk/webframe/conf"
	"github.com/fordisk/webframe/db"
	util "github.com/fordisk/webframe/frame"
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

	db.InitDb()

	db.MigrateTable()

	util.KratosServe(Name, Version, func(httpSrv *http.Server) {

	}, nil, nil)
}
