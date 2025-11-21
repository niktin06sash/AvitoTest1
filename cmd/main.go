package main

import (
	"AvitoTest1/internal/handler"
	"AvitoTest1/internal/server"
	"AvitoTest1/internal/service"
	"AvitoTest1/internal/storage"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db, err := storage.NewDBObject("")
	if err != nil {
		log.Println("Error while connect to Database")
		return
	}
	defer db.Close()
	st := storage.NewStorage(db)
	srvc := service.NewService(st.Usst, st.PRst, st.Tst, st.Usst, st.TxMan, st.Usst, st.PRst)
	handler := handler.NewHandler(srvc.UserService, srvc.TeamService, srvc.PullRequestService)
	server := server.NewServer(handler)
	//gracefull shutdown
	serverError := make(chan error, 1)
	go func() {
		if err := server.Run(); err != nil {
			serverError <- fmt.Errorf("server run failed: %w", err)
			return
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sig := <-quit:
		log.Printf("Server shutting down with signal: %v", sig)
	case err := <-serverError:
		log.Printf("Server startup failed: %v", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println("Server is shutting down...")
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
		return
	}
	log.Println("Server has shutted down successfully")
}
