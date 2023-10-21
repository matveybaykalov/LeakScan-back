package app

import (
	"backend/internal/config"
	"backend/internal/config/cli"
	v1 "backend/internal/controller/http/v1"
	postgresrepo "backend/internal/repository/postgresRepo"
	challengeservice "backend/internal/service/challengeService"
	hashservice "backend/internal/service/hashService"
	"backend/internal/usecase"
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func Run() {
	var (
		cfgLoader config.Loader
	)

	cfgLoader = cli.ConfigLoader{}
	config := cfgLoader.Load()

	logger := &logrus.Logger{ //nolint: exhaustruct
		Out:   os.Stdout,
		Level: logrus.DebugLevel,
		Formatter: &prefixed.TextFormatter{ //nolint: exhaustruct
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.PostgresConfig.DSN)))
	pgConnection := bun.NewDB(sqldb, pgdialect.New())

	pgRepo := postgresrepo.New(pgConnection)

	chgSer := challengeservice.New(logger, pgRepo)
	hashSer := hashservice.New(logger, pgRepo)

	uc := usecase.New(logger, hashSer, chgSer)

	gin.DefaultWriter = logger.WriterLevel(logrus.DebugLevel)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	httpController := v1.New(logger, router, uc)

	srv := httpController.StartServer(config.HTTPConfig)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		logger.Info("timeout of 5 seconds.")
	}
	logger.Info("Server exiting")
}
