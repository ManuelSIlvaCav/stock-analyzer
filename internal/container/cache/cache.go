package cache

import "context"

type Cache interface {
	GetConnection() interface{}
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}) error
}
