package app

import (
	"context"
	"persserv/cmd/server"
	"persserv/internal/config"
	"persserv/internal/handler"
	"persserv/internal/repository"
	"persserv/internal/repository/postgres"
	"persserv/internal/usecase"

	"github.com/sirupsen/logrus"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func Run(cfg config.Config) {
	// logger
	setupLoggerGlobally(cfg.Env)

	logrus.Info("starting app")
	logrus.Debug("debug messages are enabled")

	// db
	db, err := postgres.NewPostgresDB(cfg.DB)
	if err != nil {
		logrus.Fatalf("failed to init db: %s", err.Error())
	}

	// layers
	repos := repository.NewRepository(db)
	useCases := usecase.NewUseCase(repos)
	handlers := handler.NewHandler(useCases)

	// server
	serv := new(server.Server)

	if err := serv.Run(cfg.HTTPServer, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error running server: %s", err.Error())
	}

	logrus.Info("app started")

	if err := serv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error shutdown server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error close db: %s", err.Error())
	}
}

func setupLoggerGlobally(env string) {
	switch env {
	case envLocal:
		logrus.SetLevel(logrus.DebugLevel)
	case envDev:
		logrus.SetLevel(logrus.DebugLevel)
	case envProd:
		logrus.SetLevel(logrus.InfoLevel)
	}
}

// func setupLogger(env string) *logrus.Logger {
// 	log := logrus.New()

// 	switch env {
// 	case envLocal:
// 		log.SetLevel(logrus.DebugLevel)
// 	case envDev:
// 		log.SetLevel(logrus.DebugLevel)
// 	case envProd:
// 		log.SetLevel(logrus.InfoLevel)
// 	}

// 	return log
// }

//// Anothger Logger
// func setupLogger(env string) *slog.Logger {
// 	var log *slog.Logger

// 	switch env {
// 	case envLocal:
// 		log = slog.New(
// 			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
// 		)
// 	case envDev:
// 		log = slog.New(
// 			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
// 		)
// 	case envProd:
// 		log = slog.New(
// 			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
// 		)
// 	}

// 	return log
// }
