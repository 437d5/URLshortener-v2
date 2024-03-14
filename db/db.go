package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/437d5/URLshortener-v2/model"
	"github.com/redis/go-redis/v9"
	"time"
)

var ErrNotExist = errors.New("this url does not exit")

type RedisRepo struct {
	Client *redis.Client
}

// Create function adds token instance into redis db
func (r *RedisRepo) Create(ctx context.Context, shorten model.ShortenURL) error {
	data, err := json.Marshal(shorten)
	if err != nil {
		return fmt.Errorf("failed to encode: %w", err)
	}
	// we will get our full url using short token
	key := shorten.Token

	txn := r.Client.TxPipeline()
	// set the value only if not existed
	res := txn.SetNX(ctx, key, string(data), 1*time.Hour)
	if err := res.Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("cannot create: %w", err)
	}
	// adds token to array with name "tokens"
	if err := txn.SAdd(ctx, "tokens", key).Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to add token to tokens set: %w", err)
	}
	// execute transaction
	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("failed to execute create transaction: %w", err)
	}

	return nil
}

// GetURL function gets context and token and tries to get the model.ShortenURL object with its token
func (r *RedisRepo) GetURL(ctx context.Context, token string) (model.ShortenURL, error) {
	value, err := r.Client.Get(ctx, token).Result()
	if errors.Is(err, redis.Nil) {
		return model.ShortenURL{}, ErrNotExist
	} else if err != nil {
		return model.ShortenURL{}, fmt.Errorf("url get error: %w", err)
	}

	var url model.ShortenURL
	err = json.Unmarshal([]byte(value), &url)
	if err != nil {
		return model.ShortenURL{}, fmt.Errorf("error unmarshalling json: %w", err)
	}

	return url, nil
}

// DeleteURL function tries to delete token from tokens array
func (r *RedisRepo) DeleteURL(ctx context.Context, token string) error {
	txn := r.Client.TxPipeline()

	err := txn.Del(ctx, token).Err()
	if errors.Is(err, redis.Nil) {
		txn.Discard()
		return ErrNotExist
	} else if err != nil {
		txn.Discard()
		return fmt.Errorf("failed to delete: %w", err)
	}

	if err := txn.SRem(ctx, "tokens", token).Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to remove token from token set: %w", err)
	}

	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("failed to execute del transaction: %w", err)
	}

	return nil
}
