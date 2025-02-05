package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"example-svc/internal/some_domain/delivery/cron"
	grpcServer "example-svc/internal/some_domain/delivery/grpc"
	httpServer "example-svc/internal/some_domain/delivery/http"
	redisServer "example-svc/internal/some_domain/delivery/redis"
	"example-svc/internal/some_domain/delivery/tg"
	"example-svc/internal/some_domain/infra/clickhouse"
	"example-svc/internal/some_domain/infra/grpc"
	"example-svc/internal/some_domain/infra/http"
	"example-svc/internal/some_domain/infra/postgres"
	"example-svc/internal/some_domain/infra/redis"
	"example-svc/internal/some_domain/usecases"
	"example-svc/pkg/http/middlewares/http_error"
	errProcessor "example-svc/pkg/http/middlewares/http_error"

	"github.com/go-co-op/gocron"
)

func main() {
	// some initialization logic here and boot code here

	chClient := clickhouse.NewSomeClient()
	postgresClient := postgres.NewSomeClient()
	redisClient := redis.NewSomeClient()
	grpcClient := grpc.NewSomeClient()
	httpClient := http.NewSomeClient()

	uc := usecases.NewSomeExampleUsecase(
		httpClient,
		postgresClient,
		chClient,
		redisClient,
		grpcClient,
	)

	grpcHandler := grpcServer.NewSomeGrpcHandler(uc)
	httpHandler := httpServer.NewExampleHandlers(uc)
	redisHandler := redisServer.NewConsumerRedis(nil, log.Logger{}, uc)
	tgHandler := tg.NewTgBot(uc, nil)
	cronHandler := cron.NewHandler(uc)

	middlewareHandler := http_error.NewErrorHandler(
		log.Logger{},
		[]errProcessor.Handler{
			httpServer.ErrorHandlerMiddleware,
		},
	)

	moscowTime, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		panic(err)
	}

	s := gocron.NewScheduler(moscowTime)
	if _, err := s.Cron("* * * * *").Do(
		func() {
			cronHandler.CloseExpiredSession(context.Background())
		},
	); err != nil {
		// some log
	}

	fmt.Println(grpcHandler, httpHandler, redisHandler, tgHandler, middlewareHandler)
}
