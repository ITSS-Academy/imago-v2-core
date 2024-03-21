package profile

import (
	"errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UID       string   `json:"id" gorm:"primaryKey"`
	UserName  string   `json:"username"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Bio       string   `json:"bio"`
	Email     string   `json:"email"`
	PhotoUrl  string   `json:"photo_url"`
	Category  []string `json:"category"`
	Followers []string `json:"followers"`
	Following []string `json:"following"`
}

type ProfileRepository interface {
	GetById(ctx context.Context, id string) (*Profile, error)
	GetAll(ctx context.Context) ([]*Profile, error)
	Create(ctx context.Context, profile *Profile) error
	Update(ctx context.Context, profile *Profile) error
}

type ProfileUseCase interface {
	GetById(ctx context.Context, id string) (*Profile, error)
	GetAll(ctx context.Context) ([]*Profile, error)
	Create(ctx context.Context, profile *Profile) error
	Update(ctx context.Context, profile *Profile) error
}

type ProfileInterop interface {
	GetById(ctx context.Context, token string, id string) (*Profile, error)
	GetAll(ctx context.Context, token string) ([]*Profile, error)
	Create(ctx context.Context, token string, profile *Profile) error
	Update(ctx context.Context, token string, profile *Profile) error
	Follow(ctx context.Context, token string, profileId string, profileOther string) error
	Unfollow(ctx context.Context, token string, profileId string, profileOther string) error
}

var (
	ErrProfileExists     = errors.New("profile already exists")
	ErrProfileNotFound   = errors.New("profile not found")
	ErrIdEmpty           = errors.New("id is empty")
	ErrFieldEmpty        = errors.New("field is empty")
	ErrProfileNotCreated = errors.New("profile not created")
)
