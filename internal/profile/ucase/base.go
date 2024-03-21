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
	//TODO implement me
	panic("implement me")
}

func (p ProfileUseCase) Create(ctx context.Context, profileData *profile.Profile) error {
	//get by id to check if profile already exists then return error, else create
	_, err := p.repo.GetById(ctx, profileData.UID)
	if err == nil {
		return profile.ErrProfileExists
	}
	err = p.repo.Create(ctx, profileData)
	return nil
}

func (p ProfileUseCase) Update(ctx context.Context, profile *profile.Profile) error {
	//TODO implement me
	panic("implement me")
}

func NewProfileUseCase(repo profile.ProfileRepository) *ProfileUseCase {
	return &ProfileUseCase{
		repo: repo,
	}
}
