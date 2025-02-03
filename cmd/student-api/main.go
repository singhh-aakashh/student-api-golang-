package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/singhh-aakashh/student-api/internal/config"
)

func main() {
    if err:=godotenv.Load(); err!=nil{
		log.Fatal(err)
	}
	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /",func(w http.ResponseWriter, r *http.Request) { 
		w.Write([]byte("hello from server"))
	})
	fmt.Println("Hello world")

	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started on ",slog.String("port",cfg.Addr))

	done := make(chan os.Signal,1)

	signal.Notify(done,os.Interrupt,syscall.SIGINT,syscall.SIGTERM) 

	go func() {
		if err:= server.ListenAndServe(); err!=nil{
			log.Fatal(err)
		}
	}()

	<-done

	slog.Info("shutting down the server ")
	
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()


	if err:=server.Shutdown(ctx); err!=nil{
		slog.Error("failed to shutdown the server",slog.String("error",err.Error()))
	}
	
	slog.Info("Server shutdown successfully")
}