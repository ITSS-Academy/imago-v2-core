package post

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/itss-academy/imago/core/common"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID         string          `json:"id" gorm:"primaryKey"`
	Content    string          `json:"content"`
	CreatorId  string          `json:"creator_id"`
	CategoryId json.RawMessage `json:"category_id" gorm:"type:json"`
	PhotoUrl   json.RawMessage `json:"photo_url" gorm:"type:json"`
	Like       json.RawMessage `json:"like" gorm:"type:json"`
	Comment    json.RawMessage `json:"comment" gorm:"type:json"`
	HashTag    json.RawMessage `json:"hash_tag" gorm:"type:json"`
	Share      json.RawMessage `json:"share" gorm:"type:json"`
	Status     string          `json:"status"`
	Mention    json.RawMessage `json:"mention" gorm:"type:json"`
}

type PostRepository interface {
	Create(ctx context.Context, post *Post) error
	//GetById(ctx context.Context, id string) (*Post, error)
	//GetByUid(ctx context.Context, uid string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	List(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	//Update(ctx context.Context, post *Post) error
	//Delete(ctx context.Context, id string) error
	//GetByCategory(ctx context.Context, categoryId string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
}

type PostUseCase interface {
	Create(ctx context.Context, post *Post) error
	//GetById(ctx context.Context, id string) (*Post, error)
	//GetByUid(ctx context.Context, uid string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	List(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	//Update(ctx context.Context, post *Post) error
	//Delete(ctx context.Context, id string) error
	//GetByCategory(ctx context.Context, categoryId string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
}

type PostInterop interface {
	Create(ctx context.Context, token string, post *Post) error
	//GetById(ctx context.Context, token string, id string) (*Post, error)
	//GetByUid(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	List(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	//Update(ctx context.Context, token string, post *Post) error
	//Delete(ctx context.Context, token string, id string) error
	//GetByCategory(ctx context.Context, token string, categoryId string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
}

var (
	ErrPostNotFound        = errors.New("post not found")
	ErrPostInvalidSize     = errors.New("post invalid size")
	ErrPostInvalidPage     = errors.New("post invalid page")
	ErrPostRequiredContent = errors.New("post required content")
	ErrPostRequiredPhoto   = errors.New("post required photo")
)
