package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"go.uber.org/zap"

	"tumaris.hack-FemNovation/backend/internal/repository"
)

func Run() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Println("error zap: ", err)
		return
	}

	sugar := logger.Sugar()
	defer sugar.Sync()

	rdb, err := repository.RedisConn(sugar)
	if err != nil {
		sugar.Errorf("Cannot connect to redis: %v", err)
		return
	}
	defer rdb.Close()

	db, err := repository.DBConnection(sugar)
	if err != nil {
		sugar.Errorf("Cannot connect to db: %v", err)
		return
	}
	defer db.Close()

	repositories := repository.New(db, rdb, sugar)
	service := service.New(repositories)
	handlers := delivery.NewHandler(service, sugar)

	port := 8090
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: handlers.InitRoutes(),
	}

	errChan := make(chan error, 1)
	go func() {
		logger.Sugar().Infof("Starting server on port: %d\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
			return
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-errChan:
		sugar.Error(err.Error())
	case <-quit:
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Sugar().Debugf("Server forced to shutdown: %s", err)
		return
	}
}
