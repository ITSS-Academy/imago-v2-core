package repo

import (
	"context"
	"github.com/itss-academy/imago/core/domain/profile"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func (p ProfileRepository) GetById(ctx context.Context, id string) (*profile.Profile, error) {
	profileData := &profile.Profile{}
	tx := p.db.Where("uid = ?", id).First(profileData)
	//if profile not found return error
	if tx.Error != nil {
		return nil, tx.Error
	}
	return profileData, nil
}

func (p ProfileRepository) GetAll(ctx context.Context) ([]*profile.Profile, error) {
	var profiles []*profile.Profile
	tx := p.db.Find(&profiles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return profiles, nil
}

func (p ProfileRepository) Create(ctx context.Context, profileData *profile.Profile) error {
	tx := p.db.Create(profileData)
	return tx.Error
}

func (p ProfileRepository) Update(ctx context.Context, profile *profile.Profile) error {
	tx := p.db.Save(profile)
	return tx.Error
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	err := db.AutoMigrate(&profile.Profile{})
	if err != nil {
		panic(err)
	}
	return &ProfileRepository{db: db}
}
