package ucase

import (
	"context"
	firebaseAuth "firebase.google.com/go/v4/auth"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/auth"
)

type AuthUseCase struct {
	repo       auth.AuthRepository
	authClient *firebaseAuth.Client
}

func (a AuthUseCase) Create(ctx context.Context, authData *auth.Auth) error {
	err := a.Validate(authData)
	if err != nil {
		return err
	}
	err = a.repo.Create(ctx, authData)
	if err != nil {
		return auth.ErrAuthNotCreated
	}
	return nil
}

func (a AuthUseCase) GetById(ctx context.Context, id string) (*auth.Auth, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) Get(ctx context.Context, opts *common.QueryOpts) ([]*auth.Auth, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) Update(ctx context.Context, auth *auth.Auth) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) Verify(ctx context.Context, token string) (*firebaseAuth.UserRecord, error) {
	idToken, err := a.authClient.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, err
	}
	record, err := a.authClient.GetUser(ctx, idToken.UID)
	return record, nil
}

func (a AuthUseCase) Validate(authData *auth.Auth) error {
	if authData.Email == "" {
		return auth.ErrEmailEmpty
	}
	if authData.RoleId == "" {
		authData.RoleId = auth.RoleUser
	}
	return nil
}

func NewAuthUseCase(repo auth.AuthRepository, authClient *firebaseAuth.Client) *AuthUseCase {

	return &AuthUseCase{repo: repo, authClient: authClient}
}
