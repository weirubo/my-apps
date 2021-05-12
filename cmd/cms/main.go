package main

import (
	"fmt"
	"my-apps/api/cms/common"
	"my-apps/api/cms/pkg"
	"my-apps/api/cms/routers"
	"net/http"
	"time"
)

func init() {
	err := initConfig()
	if err != nil {
		fmt.Println("initConfig() error:", err)
		return
	}
	err = initDBEngine()
	if err != nil {
		fmt.Println("initDBEngine() error:", err)
		return
	}
}

func main() {
	r := routers.NewRouter()
	server := &http.Server{
		Addr:           ":" + common.ServerConfig.HttpPort,
		Handler:        r,
		ReadTimeout:    common.ServerConfig.ReadTimeout,
		WriteTimeout:   common.ServerConfig.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = server.ListenAndServe()
}

func initConfig () error {
	config, err := pkg.NewViper()
	if err != nil {
		return err
	}
	err = config.ReadConfig("Server", &common.ServerConfig)
	if err != nil {
		return err
	}
	err = config.ReadConfig("Database", &common.DatabaseConfig)
	if err != nil {
		return err
	}
	common.ServerConfig.ReadTimeout *= time.Second
	common.ServerConfig.WriteTimeout *= time.Second
	return nil
}

func initDBEngine() error {
	var err error
	common.DBEngine, err = pkg.NewDBEngine(common.DatabaseConfig)
	if err != nil {
		return err
	}
	return nil
}
