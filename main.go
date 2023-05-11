package main

import (
	"context"
	"fmt"
	"mionext/srv/app/db"
	"mionext/srv/app/http/server"
	"mionext/srv/app/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("[ERROR] Setup config failed: %s", err.Error())
		os.Exit(2)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config change: ", e.String())
	})

	if err := logger.Setup(); err != nil {
		fmt.Printf("[ERROR] Setup logger failed: %s", err.Error())
		os.Exit(2)
	}
}

func main() {
	srv := server.Setup()
	go func() {
		logrus.Infof("Start serving on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Failed to listen and serve: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT)
	sig := <-quit

	logrus.Warnf("Shutdown server on signal: %s", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server shutdown: %s", err.Error())
	}

	logrus.Warnf("Server exiting.")
	// Destruct resources
	defer func() {
		db.Close()
	}()
}
