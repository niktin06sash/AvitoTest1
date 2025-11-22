package main

import (
	"AvitoTest1/config"
	"AvitoTest1/internal/handler"
	"AvitoTest1/internal/logger"
	"AvitoTest1/internal/server"
	"AvitoTest1/internal/service"
	"AvitoTest1/internal/storage"
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	loger, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	defer loger.Sync()
	db, err := storage.NewDBObject(conf.Database)
	if err != nil {
		panic(err)
	}
	loger.ZapLogger.Debug("Successful connect to database")
	defer db.Close(loger)
	st := storage.NewStorage(db)
	srvc := service.NewService(loger, st.Usst, st.PRst, st.Tst, st.Usst, st.TxMan, st.Usst, st.PRst)
	handler := handler.NewHandler(loger, srvc.UserService, srvc.TeamService, srvc.PullRequestService)
	server := server.NewServer(conf.Server, handler)
	//gracefull shutdown
	serverError := make(chan error, 1)
	go func() {
		if err := server.Run(); err != nil {
			serverError <- err
			return
		}
		loger.ZapLogger.Debug("Successful start server")
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sig := <-quit:
		loger.ZapLogger.Debug("Server shutting down with signal", zap.Any("signal", sig))
	case err := <-serverError:
		loger.ZapLogger.Debug("Server startup failed", zap.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), conf.Server.GracefulShutdown)
	defer cancel()
	loger.ZapLogger.Debug("Server is shutting down...")
	if err := server.Shutdown(ctx); err != nil {
		loger.ZapLogger.Debug("Server shutdown error", zap.Error(err))
		return
	}
	loger.ZapLogger.Debug("Server has shutted down successfully!")
}
