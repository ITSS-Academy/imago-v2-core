package auth

import (
	"context"
	"errors"
	"firebase.google.com/go/v4/auth"
	"github.com/itss-academy/imago/core/common"
	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	ID     string `json:"id" gorm:"primaryKey"`
	Email  string `json:"email"`
	RoleId string `json:"role_id"`
	Status string `json:"status"`
}

type AuthRepository interface {
	Create(ctx context.Context, auth *Auth) error
	GetById(ctx context.Context, id string) (*Auth, error)
	Get(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Auth], error)
	Update(ctx context.Context, auth *Auth) error
	Delete(ctx context.Context, id string) error
}

type AuthUseCase interface {
	Create(ctx context.Context, auth *Auth) error
	GetById(ctx context.Context, id string) (*Auth, error)
	Get(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Auth], error)
	Update(ctx context.Context, auth *Auth) error

	Delete(ctx context.Context, id string) error
	Verify(ctx context.Context, token string) (*auth.UserRecord, error)
}

type AuthInterop interface {
	Create(ctx context.Context, token string, auth *Auth) error
	GetById(ctx context.Context, token string, id string) (*Auth, error)
	Get(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*Auth], error)
	Update(ctx context.Context, token string, auth *Auth) error
	ChangeRole(ctx context.Context, token string, roleId string) error
	Delete(ctx context.Context, token string, id string) error
}

const RoleAdmin = "admin"
const RoleUser = "user"

var (
	ErrEmailEmpty        = errors.New("email is empty")
	ErrRoleIdEmpty       = errors.New("role id is empty")
	ErrAuthNotFound      = errors.New("auth not found")
	ErrAuthNotValid      = errors.New("auth not valid")
	ErrAuthNotAuthorized = errors.New("auth not authorized")
	ErrAuthNotDeleted    = errors.New("auth not deleted")
	ErrAuthNotUpdated    = errors.New("auth not updated")
	ErrAuthNotCreated    = errors.New("auth not created")
	ErrInvalidAuthPage   = errors.New("invalid auth page")
	ErrInvalidAuthSize   = errors.New("invalid auth size")
)
