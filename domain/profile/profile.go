package profile

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/auth"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UID       string          `json:"uid" gorm:"primaryKey" `
	UserName  string          `json:"username"`
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Bio       string          `json:"bio"`
	Email     string          `json:"email"`
	PhotoUrl  string          `json:"photo_url"`
	Category  JSONStringArray `json:"category" gorm:"type:json"`
	Followers JSONStringArray `json:"followers" gorm:"type:json"`
	Following JSONStringArray `json:"following" gorm:"type:json"`
}
type JSONStringArray []string

type AuthProfile struct {
	*auth.Auth
	Profile    *Profile `json:"profile"`
	NumberPost int
}

func (a *JSONStringArray) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &a)
}

func (a JSONStringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type ProfileRepository interface {
	GetById(ctx context.Context, id string) (*Profile, error)
	GetAll(ctx context.Context) ([]*Profile, error)
	Create(ctx context.Context, profile *Profile) error
	Update(ctx context.Context, profile *Profile) error
	GetAllAuthNoProfile(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*auth.Auth], error)
	GetAllAuthProfile(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*AuthProfile], error)
}

type ProfileUseCase interface {
	GetById(ctx context.Context, id string) (*Profile, error)
	GetAll(ctx context.Context) ([]*Profile, error)
	Create(ctx context.Context, profile *Profile) error
	Update(ctx context.Context, profile *Profile) error
	GetAllAuthNoProfile(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*auth.Auth], error)
	GetAllAuthProfile(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*AuthProfile], error)
}

type ProfileInterop interface {
	GetById(ctx context.Context, token string, id string) (*Profile, error)
	GetMine(ctx context.Context, token string) (*Profile, error)
	GetAll(ctx context.Context, token string) ([]*Profile, error)
	Create(ctx context.Context, token string, profile *Profile) error
	Update(ctx context.Context, token string, profile *Profile) error
	Follow(ctx context.Context, token string, profileId string, profileOtherId string) error
	Unfollow(ctx context.Context, token string, profileId string, profileOtherId string) error
	GetAllAuthNoProfile(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*auth.Auth], error)
	GetAllAuthProfile(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*AuthProfile], error)
	GetAllExceptMine(ctx context.Context, token string) ([]*Profile, error)
}

var (
	ErrTokenEmpty        = errors.New("token is empty")
	ErrProfileExists     = errors.New("profile already exists")
	ErrProfileNotFound   = errors.New("profile not found")
	ErrIdEmpty           = errors.New("id is empty")
	ErrFieldEmpty        = errors.New("field is empty")
	ErrProfileNotCreated = errors.New("profile not created")
)
