package main

import (
	"fmt"
	"my-apps/api/cms/pkg"
	"my-apps/api/cms/routers"
	"net/http"
	"time"
)

func init() {
	err := initDBEngine()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	r := routers.NewRouter()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: 1 << 20,
	}
	_ = server.ListenAndServe()
}

func initDBEngine() error {
	var err error
	pkg.DBEngine, err = pkg.NewDBEngine()
	if err != nil {
		return err
	}
	return nil
}
