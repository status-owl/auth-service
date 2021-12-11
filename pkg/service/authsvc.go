package service

import "context"

type AuthService interface {
	CreateAuthMethod(ctx context.Context, userId string, pwd []byte) (string, error)
	DeleteAuthMethod(ctx context.Context, id string) error
}
