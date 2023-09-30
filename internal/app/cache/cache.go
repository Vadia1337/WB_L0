package cache

import (
	"WB_L0/internal/app/model"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func New(redisAddr string) (*Cache, error) {
	conn := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	_, err := conn.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &Cache{
		client: conn,
	}, nil
}

func (c *Cache) Save(ctx context.Context, key string, order *model.Order) error {
	value, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = c.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache) Get(ctx context.Context, key string) (*model.Order, error) {
	order := model.Order{}

	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(val), &order)
	if err != nil {
		return nil, err
	}

	return &order, err
}
