package main

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"mionext/srv/app/http/engine"
	"mionext/srv/app/logger"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func init() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("[ERROR] Init config failed: %s", err.Error())
		os.Exit(2)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config change: ", e.String())
	})

	if err := logger.Init(); err != nil {
		fmt.Printf("[ERROR] Init logger failed: %s", err.Error())
		os.Exit(2)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	server := &http.Server{Addr: viper.GetString("http.listen"), Handler: engine.Init()}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("[ERROR] Failed to listen and serve: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT)
	<-quit
	logrus.Println("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server shutdown: %s", err.Error())
	}

	// db.Close()
	logrus.Fatalf("Server exiting.")
}
