package cachemanager

import "context"

type CacheManager interface {
	GetString(context.Context, string) (string, error)
	GetInt64(context.Context, string) (int64, error)
	GetInterface(context.Context, string) (interface{}, error)
	Add(context.Context, string, interface{}) (bool, error)
	IsExisted(context.Context, string) (bool, error)
	Lock(context.Context, string) (bool, error)
	Release(context.Context, string) (bool, error)
	IsLocked(context.Context, string) (bool, error)
}
