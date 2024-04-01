package repo

import (
	"context"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/auth"
	"github.com/itss-academy/imago/core/domain/post"
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

func (p ProfileRepository) GetAllAuthNoProfile(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*auth.Auth], error) {
	var auths []*auth.Auth
	var profiles []*profile.Profile
	tx := p.db.Find(&auths)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = p.db.Find(&profiles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var result []*auth.Auth
	for i := 0; i < len(auths); i++ {
		isExist := false
		for j := 0; j < len(profiles); j++ {
			if auths[i].ID == profiles[j].UID {
				isExist = true
				break
			}
		}
		if !isExist {
			result = append(result, auths[i])
		}
	}
	count := int64(0)
	tx = p.db.Model(&auth.Auth{}).Count(&count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	pageNum := int64(0)
	if count > 0 {
		pageNum = int64(math.Ceil(float64(count) / float64(opts.Size)))
	}
	return &common.ListResult[*auth.Auth]{Data: result, EndPage: int(pageNum)}, nil
}

func (p ProfileRepository) GetAllAuthProfile(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*profile.AuthProfile], error) {
	var auths []*auth.Auth
	var profiles []*profile.Profile
	var posts []*post.Post
	tx := p.db.Find(&auths)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = p.db.Find(&profiles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = p.db.Find(&posts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []*profile.AuthProfile
	var countPost int
	for i := 0; i < len(auths); i++ {
		for j := 0; j < len(profiles); j++ {
			if auths[i].ID == profiles[j].UID {
				countPost = 0
				for k := 0; k < len(posts); k++ {
					if profiles[j].UID == posts[k].CreatorId {
						countPost++
					}
				}
				result = append(result, &profile.AuthProfile{Auth: auths[i], Profile: profiles[j], NumberPost: countPost})
			}
		}
	}
	count := int64(0)
	tx = p.db.Model(&auth.Auth{}).Count(&count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	pageNum := int64(0)
	if count > 0 {
		pageNum = int64(math.Ceil(float64(count) / float64(opts.Size)))
	}
	return &common.ListResult[*profile.AuthProfile]{Data: result, EndPage: int(pageNum)}, nil
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	err := db.AutoMigrate(&profile.Profile{})
	if err != nil {
		panic(err)
	}
	return &ProfileRepository{db: db}
}
