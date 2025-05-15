package utils

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartServerWithGracefulShutdown(server *http.Server, shutdownTimeout time.Duration) {
	go func() {
		log.Printf("\t\u2705 Server is running: http://127.0.0.1%s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("\t\u274c Server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("\t\u2705 Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("\t\u274c Server forced to shutdown: %v", err)
	}

	log.Println("\t\u2705 Server is ended gracefully")
}
