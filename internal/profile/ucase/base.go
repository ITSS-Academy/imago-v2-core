package ucase

import (
	"context"
	"github.com/itss-academy/imago/core/domain/profile"
)

type ProfileUseCase struct {
	repo profile.ProfileRepository
}

func (p ProfileUseCase) GetById(ctx context.Context, id string) (*profile.Profile, error) {
	data, err := p.repo.GetById(ctx, id)
	if err != nil {
		return nil, profile.ErrProfileNotFound
	}
	return data, nil
}

func (p ProfileUseCase) GetAll(ctx context.Context) ([]*profile.Profile, error) {
	profiles, err := p.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func (p ProfileUseCase) Create(ctx context.Context, profileData *profile.Profile) error {
	//get by id to check if profile already exists then return error, else create
	_, err := p.repo.GetById(ctx, profileData.UID)
	if err == nil {
		return profile.ErrProfileExists
	}
	//if any of the fields is empty return error
	if profileData.UserName == "" || profileData.FirstName == "" || profileData.LastName == "" {
		return profile.ErrFieldEmpty
	}
	err = p.repo.Create(ctx, profileData)
	if err != nil {
		return err
	}
	return nil
}

func (p ProfileUseCase) Update(ctx context.Context, profileData *profile.Profile) error {
	//get by id to check if profile already exists then return error, else update
	_, err := p.repo.GetById(ctx, profileData.UID)
	if err != nil {
		return profile.ErrProfileNotFound
	}
	//if any of the fields is empty return error
	if profileData.UserName == "" || profileData.FirstName == "" || profileData.LastName == "" {
		return profile.ErrFieldEmpty
	}
	err = p.repo.Update(ctx, profileData)
	if err != nil {
		return err
	}
	return nil
}

func NewProfileUseCase(repo profile.ProfileRepository) *ProfileUseCase {
	return &ProfileUseCase{
		repo: repo,
	}
}
