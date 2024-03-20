package ucase

import (
	"context"
	"github.com/itss-academy/imago/core/domain/profile"
)

type ProfileUseCase struct {
	repo profile.ProfileRepository
}

func (p ProfileUseCase) GetById(ctx context.Context, id string) (*profile.Profile, error) {
	if id == "" {
		return nil, profile.ErrIdEmpty
	}
	profileData, err := p.repo.GetById(ctx, id)
	if err != nil {
		return nil, profile.ErrProfileNotFound
	}
	return profileData, nil
}

func (p ProfileUseCase) GetAll(ctx context.Context) ([]*profile.Profile, error) {
	profileList, err := p.repo.GetAll(ctx)
	if err != nil {
		return nil, profile.ErrProfileNotFound
	}
	return profileList, nil
}

func (p ProfileUseCase) Create(ctx context.Context, profileData *profile.Profile) error {
	existedProfile, err := p.repo.GetById(ctx, profileData.ID)
	if existedProfile != nil {
		return profile.ErrProfileExists
	}
	err = p.Validate(profileData)
	if err != nil {
		return err
	}
	err = p.repo.Create(ctx, profileData)
	if err != nil {
		return profile.ErrProfileNotCreated

	}
	return nil

}

func (p ProfileUseCase) Update(ctx context.Context, profileData *profile.Profile) error {
	err := p.Validate(profileData)
	if err != nil {
		return err
	}
	err = p.repo.Update(ctx, profileData)
	if err != nil {
		return profile.ErrProfileNotCreated
	}
	return nil
}

func NewProfileUseCase(repo profile.ProfileRepository) *ProfileUseCase {
	return &ProfileUseCase{repo: repo}
}

func (p ProfileUseCase) Validate(profileData *profile.Profile) error {
	if profileData.UserName == "" {
		return profile.ErrFieldEmpty
	}
	if profileData.FirstName == "" {
		return profile.ErrFieldEmpty
	}
	if profileData.LastName == "" {
		return profile.ErrFieldEmpty
	}
	return nil
}
