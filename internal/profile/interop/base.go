package interop

import (
	"context"
	"github.com/itss-academy/imago/core/domain/auth"
	"github.com/itss-academy/imago/core/domain/profile"
)

type ProfileInterop struct {
	ucase     profile.ProfileUseCase
	authucase auth.AuthUseCase
}

func (p ProfileInterop) GetById(ctx context.Context, token string, id string) (*profile.Profile, error) {
	return p.ucase.GetById(ctx, id)
}

func (p ProfileInterop) GetAll(ctx context.Context, token string) ([]*profile.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileInterop) Create(ctx context.Context, token string, profileData *profile.Profile) error {
	record, err := p.authucase.Verify(ctx, token)
	if err != nil {
		return err
	}
	profileData.UID = record.UID
	profileData.Email = record.Email
	return p.ucase.Create(ctx, profileData)
}

func (p ProfileInterop) Update(ctx context.Context, token string, profile *profile.Profile) error {
	//TODO implement me
	panic("implement me")
}

func (p ProfileInterop) Follow(ctx context.Context, token string, profileId string, profileOther string) error {
	//TODO implement me
	panic("implement me")
}

func (p ProfileInterop) Unfollow(ctx context.Context, token string, profileId string, profileOther string) error {
	//TODO implement me
	panic("implement me")
}

func NewProfileInterop(ucase profile.ProfileUseCase, authucase auth.AuthUseCase) *ProfileInterop {
	return &ProfileInterop{
		ucase:     ucase,
		authucase: authucase,
	}
}
