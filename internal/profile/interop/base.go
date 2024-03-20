package interop

import (
	"context"
	"github.com/itss-academy/imago/core/domain/auth"
	"github.com/itss-academy/imago/core/domain/profile"
)

type ProfileInterop struct {
	ucase   profile.ProfileUseCase
	usecase auth.AuthUseCase
}

func (p ProfileInterop) GetById(ctx context.Context, token string, id string) (*profile.Profile, error) {
	//verify token
	_, err := p.usecase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	//get profile by id
	profileData, err := p.ucase.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return profileData, nil
}

func (p ProfileInterop) GetAll(ctx context.Context, token string) ([]*profile.Profile, error) {
	//verify token
	_, err := p.usecase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	//get all profiles
	profileList, err := p.ucase.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return profileList, nil
}

func (p ProfileInterop) Create(ctx context.Context, token string, profileData *profile.Profile) error {
	record, err := p.usecase.Verify(ctx, token)
	if err != nil {
		return err
	}
	profileData.ID = record.UID
	profileData.Email = record.Email
	profileData.UserName = ""
	profileData.FirstName = ""
	profileData.LastName = ""
	profileData.Bio = ""
	profileData.PhotoURL = ""
	profileData.Category = []string{}
	profileData.Followers = []string{}
	profileData.Following = []string{}

	return p.ucase.Create(ctx, profileData)
}

func (p ProfileInterop) Update(ctx context.Context, token string, profileData *profile.Profile) error {
	//verify token
	decodeToken, err := p.usecase.Verify(ctx, token)
	if err != nil {
		return err
	}
	existingProfile, err := p.ucase.GetById(ctx, decodeToken.UID)
	if err != nil {
		return err
	}
	existingProfile.Bio = profileData.Bio
	existingProfile.PhotoURL = profileData.PhotoURL
	existingProfile.UserName = profileData.UserName
	existingProfile.FirstName = profileData.FirstName
	existingProfile.LastName = profileData.LastName

	err = p.ucase.Update(ctx, existingProfile)
	if err != nil {
		return err
	}
	return nil
}

func (p ProfileInterop) Follow(ctx context.Context, token string, profileId string, profileOther string) error {
	//verify token
	_, err := p.usecase.Verify(ctx, token)

	if err != nil {
		return err
	}
	return nil
}

func (p ProfileInterop) Unfollow(ctx context.Context, token string, profileId string, profileOther string) error {
	//TODO implement me
	panic("implement me")
}

func NewProfileInterop(ucase profile.ProfileUseCase) *ProfileInterop {
	return &ProfileInterop{ucase: ucase}
}

// func check id profile exists in array
func checkProfileExists(profileList []string, profileId string) bool {
	for _, id := range profileList {
		if id == profileId {
			return true
		}
	}
	return false
}
