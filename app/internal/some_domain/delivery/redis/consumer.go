package redis

import (
	"context"
	"example-svc/internal/some_domain/delivery"
	"log"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type ConsumerRedis struct {
	redisPS *redis.PubSub
	logger  log.Logger
	uc      delivery.SomeDomain
}

func NewConsumerRedis(
	redis *redis.PubSub,
	logger log.Logger,
	uc delivery.SomeDomain,
) *ConsumerRedis {
	return &ConsumerRedis{
		redisPS: redis,
		logger:  logger,
		uc:      uc,
	}
}

func (c *ConsumerRedis) ConsumeMessage(ctx context.Context) {
	for {
		select {
		case msg, ok := <-c.redisPS.Channel():
			if ok {
				go func() {
					clientOrderID, err := uuid.Parse(msg.Payload)
					if err != nil {
						// c.logger.Error()
						c.logger.Println(
							"ConsumeMessage.UUID.Parse(). Unable convert msg to uuid",
							"clientOrderID", clientOrderID,
						)
						return
					}
					id, err := c.uc.ProcessSomeEvent(ctx, clientOrderID)
					if err != nil {
						// c.logger.Error()
						c.logger.Println(
							"ConsumeMessage.TryToMatchOrdersHandler.Handle(). Error",
							"error", err,
							"clientOrderID", clientOrderID,
						)
						return
					}
					if id == uuid.Nil {
						// c.logger.Error()
						c.logger.Println(
							"ConsumeMessage.TryToMatchOrdersHandler.Handle(). Error",
							"error", err,
							"clientOrderID", clientOrderID,
						)
						return
					}
					// c.logger.Info()
					c.logger.Println(
						"ConsumeMessage.TryToMatchOrdersHandler.Handle(). Success",
						"error", err,
						"clientOrderID", clientOrderID,
					)
				}()
			}
		}
	}
}
