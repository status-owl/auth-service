package store

import "context"

type AuthStore interface {
	Create(ctx context.Context, userId string, pwd []byte) (string, error)
	Delete(ctx context.Context, id string) error
}
