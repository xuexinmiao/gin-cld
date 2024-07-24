package main

import (
	"CLD/dao/mysql"
	"CLD/logger"
	"CLD/routes"
	"CLD/settings"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var err error

func main() {
	if err = settings.Init(); err != nil {
		zap.L().Error("settings init failed", zap.Error(err))
	}

	if err = logger.Init(); err != nil {
		zap.L().Error("logger init failed", zap.Error(err))
	}
	defer zap.L().Sync()

	if err = mysql.Init(); err != nil {
		zap.L().Error("mysql init failed", zap.Error(err))
	}
	defer mysql.Close()
	/*
		少redis配置，当前环境没有redis，避免报错
		if err = redis.Init(); err != nil {
			zap.L().Error("redis init failed", zap.Error(err))
		}
		defer redis.Close()
	*/

	zap.L().Info("start server")
	r := routes.Setup()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen: ", zap.Error(err))
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown: ", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
