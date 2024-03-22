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
	profiles, err := p.ucase.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func (p ProfileInterop) GetMine(ctx context.Context, token string) (*profile.Profile, error) {
	record, err := p.authucase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return p.ucase.GetById(ctx, record.UID)

}

func (p ProfileInterop) Create(ctx context.Context, token string, profileData *profile.Profile) error {
	record, err := p.authucase.Verify(ctx, token)
	if err != nil {
		return err
	}
	profileData.UID = record.UID
	profileData.Email = record.Email
	if profileData.UserName == "" || profileData.FirstName == "" || profileData.LastName == "" {
		return profile.ErrFieldEmpty
	}
	return p.ucase.Create(ctx, profileData)
}

func (p ProfileInterop) Update(ctx context.Context, token string, profileData *profile.Profile) error {
	// Verify the token
	record, err := p.authucase.Verify(ctx, token)
	if err != nil {
		return err
	}
	currentProfile, err := p.ucase.GetById(ctx, record.UID)
	if err != nil {
		return profile.ErrProfileNotFound
	}
	if profileData.UserName != "" {
		currentProfile.UserName = profileData.UserName
	}
	if profileData.FirstName != "" {
		currentProfile.FirstName = profileData.FirstName
	}
	if profileData.LastName != "" {
		currentProfile.LastName = profileData.LastName
	}
	if profileData.Bio != "" {
		currentProfile.Bio = profileData.Bio
	}

	err = p.ucase.Update(ctx, currentProfile)
	if err != nil {
		return err
	}
	return nil
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
