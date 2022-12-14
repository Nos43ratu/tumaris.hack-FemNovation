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

	_ "github.com/mattn/go-sqlite3"

	"tumaris.hack-FemNovation/backend/internal/delivery"
	"tumaris.hack-FemNovation/backend/internal/repository"
	"tumaris.hack-FemNovation/backend/internal/service"
	"tumaris.hack-FemNovation/backend/pkg/hash"
)

func Run() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Println("error zap: ", err)
		return
	}

	sugar := logger.Sugar()
	defer sugar.Sync()

	sqlite, err := repository.SqliteConnection()
	if err != nil {
		sugar.Error("sqlite connection error: ", err)
		return
	}
	defer sqlite.Close()

	db, err := repository.DBConnection(sugar)
	if err != nil {
		sugar.Errorf("Cannot connect to db: %v", err)
		return
	}
	defer db.Close()

	hasher := hash.NewByCryptHasher("67072341-eb28-4174-a01f-baf72c40b966")

	repositories := repository.New(db, sqlite, 10*time.Second, sugar)
	service := service.New(repositories, hasher, 15*time.Hour, 15*time.Hour)
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
