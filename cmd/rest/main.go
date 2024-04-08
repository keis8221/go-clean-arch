package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	adaptorHTTP "github.com/keis8221/go-clean-arch/pkg/adaptor/http"
	"github.com/keis8221/go-clean-arch/pkg/config"
)

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	addr := config.LoadConfig().HTTPInfo.Addr
	router := adaptorHTTP.InitRouter()
	srv := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// graceful shutdown
	go func() {
		log.Println("server is running! addr: ", addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve %v", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	stop()
	log.Println("shutdown down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
