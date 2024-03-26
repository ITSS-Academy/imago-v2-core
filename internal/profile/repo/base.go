package repo

import (
	"context"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/auth"
	"github.com/itss-academy/imago/core/domain/profile"
	"gorm.io/gorm"
	"math"
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

func (p ProfileRepository) GetAllAuthNoProfile(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*any], error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileRepository) GetAllAuthProfile(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*any], error) {
	auth := &auth.Auth{}
	profile := &profile.Profile{}
	limit := opts.Size
	offset := opts.Size * (opts.Page - 1)
	var result []*any
	if auth.ID == profile.UID {
		result = append(result)
	}
	tx := p.db.Find(&result).Limit(limit).Offset(offset)
	if tx.Error != nil {
		return nil, tx.Error
	}
	count := int64(0)
	if tx.Error != nil {
		return nil, tx.Error
	}
	pageNum := int64(0)
	if count > 0 {
		pageNum = int64(math.Ceil(float64(count) / float64(limit)))
	}
	return &common.ListResult[*any]{Data: result, EndPage: int(pageNum)}, tx.Error
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	err := db.AutoMigrate(&profile.Profile{})
	if err != nil {
		panic(err)
	}
	return &ProfileRepository{db: db}
}
