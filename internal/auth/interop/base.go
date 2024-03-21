package interop

import (
	"context"
	"fmt"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/auth"
)

type AuthInterop struct {
	ucase auth.AuthUseCase
}

func (a AuthInterop) Create(ctx context.Context, token string, aut *auth.Auth) error {
	record, err := a.ucase.Verify(ctx, token)
	fmt.Print(record)
	if err != nil {
		return err
	}
	aut.ID = record.UID
	aut.Email = record.Email
	aut.RoleId = auth.RoleUser
	return a.ucase.Create(ctx, aut)
}

func (a AuthInterop) GetById(ctx context.Context, token string, id string) (*auth.Auth, error) {
	return a.ucase.GetById(ctx, id)
}

func (a AuthInterop) Get(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*auth.Auth], error) {
	_, err := a.ucase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return a.ucase.Get(ctx, opts)
}

func (a AuthInterop) Update(ctx context.Context, token string, auth *auth.Auth) error {
	record, err := a.ucase.Verify(ctx, token)
	if err != nil {
		return err
	}
	authData, err := a.ucase.GetById(ctx, record.UID)
	authData.Email = auth.Email
	authData.RoleId = auth.RoleId
	authData.Status = auth.Status
	if err != nil {
		return err
	}
	return a.ucase.Update(ctx, authData)
}

func (a AuthInterop) ChangeRole(ctx context.Context, token string, roleId string) error {
	record, err := a.ucase.Verify(ctx, token)
	if err != nil {
		return err
	}
	authData, err := a.ucase.GetById(ctx, record.UID)
	if err != nil {
		return err

	}
	if authData.RoleId != auth.RoleAdmin {
		return auth.ErrAuthNotAuthorized
	}
	return a.ucase.Update(ctx, authData)
}

func (a AuthInterop) Delete(ctx context.Context, token string, id string) error {
	return a.ucase.Delete(ctx, id)
}

func NewAuthInterop(ucase auth.AuthUseCase) *AuthInterop {
	return &AuthInterop{ucase: ucase}
}
