package usecases

import "errors"

var (
	ErrNoAccess      = errors.New("no access")
	ErrCacheNotFound = errors.New("cache not found")
)
