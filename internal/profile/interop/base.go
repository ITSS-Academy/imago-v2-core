package interop

import (
	"context"
	"github.com/itss-academy/imago/core/common"
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
	profileData.Category = []string{}
	profileData.Followers = []string{}
	profileData.Following = []string{}
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
	if profileData.Email != "" {
		currentProfile.Email = profileData.Email
	}
	if profileData.PhotoUrl != "" {
		currentProfile.PhotoUrl = profileData.PhotoUrl
	}
	if profileData.Category != nil {
		currentProfile.Category = profileData.Category
	}

	//follow
	if profileData.Following != nil {
		currentProfile.Following = profileData.Following
	}
	if profileData.Followers != nil {
		currentProfile.Followers = profileData.Followers

	}

	err = p.ucase.Update(ctx, currentProfile)
	if err != nil {
		return err
	}
	return nil
}

func (p ProfileInterop) Follow(ctx context.Context, token string, profileId string, profileOtherId string) error {
	//verify the token
	_, err := p.authucase.Verify(ctx, token)
	if err != nil {
		return err
	}

	profile, err := p.ucase.GetById(ctx, profileId)
	if err != nil {
		return err
	}

	otherProfile, err := p.ucase.GetById(ctx, profileOtherId)
	if err != nil {
		return err
	}

	if isExisted(profile.Following, profileOtherId) {
		return nil
	}

	profile.Following = append(profile.Following, profileOtherId)
	otherProfile.Followers = append(otherProfile.Followers, profileId)

	err = p.ucase.Update(ctx, profile)
	if err != nil {
		return err

	}
	err = p.ucase.Update(ctx, otherProfile)
	if err != nil {
		return err
	}
	return nil
}

func (p ProfileInterop) Unfollow(ctx context.Context, token string, profileId string, profileOtherId string) error {
	//verify the token
	_, err := p.authucase.Verify(ctx, token)
	if err != nil {
		return err
	}

	profile, err := p.ucase.GetById(ctx, profileId)
	if err != nil {
		return err
	}

	otherProfile, err := p.ucase.GetById(ctx, profileOtherId)
	if err != nil {
		return err
	}

	if !isExisted(profile.Following, profileOtherId) {
		return nil
	}

	profile.Following = remove(profile.Following, profileOtherId)
	otherProfile.Followers = remove(otherProfile.Followers, profileId)

	err = p.ucase.Update(ctx, profile)
	if err != nil {
		return err

	}
	err = p.ucase.Update(ctx, otherProfile)
	if err != nil {
		return err
	}
	return nil
}

func (p ProfileInterop) GetAllAuthNoProfile(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*auth.Auth], error) {
	_, err := p.authucase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return p.ucase.GetAllAuthNoProfile(ctx, opts)
}

func (p ProfileInterop) GetAllAuthProfile(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*profile.AuthProfile], error) {
	_, err := p.authucase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return p.ucase.GetAllAuthProfile(ctx, opts)
}

func NewProfileInterop(ucase profile.ProfileUseCase, authucase auth.AuthUseCase) *ProfileInterop {
	return &ProfileInterop{
		ucase:     ucase,
		authucase: authucase,
	}
}

func isExisted(checkArray []string, id string) bool {
	for _, item := range checkArray {
		if item == id {
			return true
		}
	}
	return false
}

func remove(slice []string, s string) []string {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
