package post

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/itss-academy/imago/core/common"
	"gorm.io/gorm"
)

type MultiString []string

func (s *MultiString) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, s)
	case string:
		return json.Unmarshal([]byte(v), s)
	default:
		return errors.New("unsupported type for ShareType")
	}
}

func (s MultiString) Value() (driver.Value, error) {
	return json.Marshal(s)
}

type Post struct {
	gorm.Model
	ID         string      `json:"id" gorm:"primaryKey"`
	Content    string      `json:"content"`
	CreatorId  string      `json:"creator_id"`
	CategoryId MultiString `json:"category_id" gorm:"type:text" `
	PhotoUrl   MultiString `json:"photo_url" gorm:"type:text" `
	Like       MultiString `json:"like" gorm:"type:text"`
	Comment    MultiString `json:"comment"gorm:"type:text" `
	HashTag    MultiString `json:"hash_tag" gorm:"type:text" `
	Share      MultiString `json:"share" gorm:"type:text" `
	Status     string      `json:"status"`
	Mention    MultiString `json:"mention" gorm:"type:text"`
}

type JSONStringArray []string

type PostRepository interface {
	Create(ctx context.Context, post *Post) error
	GetDetail(ctx context.Context, id string) (*Post, error)
	GetOther(ctx context.Context, uid string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	GetMine(ctx context.Context, uid string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	GetShared(ctx context.Context, uid string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	List(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	//Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id string) error
	GetByCategory(ctx context.Context, categoryId string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	//UpdatePostComment(ctx context.Context, id string, comment []string) error
	//Delete(ctx context.Context, id string) error
	//GetByCategory(ctx context.Context, categoryId string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	UpdatePostComment(ctx context.Context, id string, post *Post) error
}

type PostUseCase interface {
	Create(ctx context.Context, post *Post) error
	GetDetail(ctx context.Context, id string) (*Post, error)
	GetByUid(ctx context.Context, uid string, opts *common.QueryOpts, style string) (*common.ListResult[*Post], error)
	GetOther(ctx context.Context, uid string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	List(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	//Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id string) error
	GetByCategory(ctx context.Context, categoryId string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	//UpdatePostComment(ctx context.Context, id string, comment []string) error
	//Delete(ctx context.Context, id string) error
	//GetByCategory(ctx context.Context, categoryId string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	UpdatePostComment(ctx context.Context, id string, post *Post) error
}

type PostInterop interface {
	Create(ctx context.Context, token string, post *Post) error
	GetDetail(ctx context.Context, token string, id string) (*Post, error)
	GetByUid(ctx context.Context, token string, opts *common.QueryOpts, style string) (*common.ListResult[*Post], error)
	GetOther(ctx context.Context, token string, uid string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	List(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	//Update(ctx context.Context, token string, post *Post) error
	Delete(ctx context.Context, token string, id string) error
	GetByCategory(ctx context.Context, token string, categoryId string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	//UpdatePostComment(ctx context.Context, token string, id string, comment []string) error
	//Delete(ctx context.Context, token string, id string) error
	//GetByCategory(ctx context.Context, token string, categoryId string, opts *common.QueryOpts) (*common.ListResult[*Post], error)
	UpdatePostComment(ctx context.Context, token string, id string, post *Post) error
}

var (
	ErrPostNotFound          = errors.New("post not found")
	ErrPostInvalidSize       = errors.New("post invalid size")
	ErrPostInvalidPage       = errors.New("post invalid page")
	ErrPostNotCreated        = errors.New("post not created")
	ErrPostRequiredID        = errors.New("post required id")
	ErrPostRequiredCreatorID = errors.New("post required creator id")
	ErrPostRequiredContent   = errors.New("post required content")
	ErrPostRequiredPhoto     = errors.New("post required photo")
	ErrPostRequiredComment   = errors.New("post required comment")
	ErrPostCommentNotUpdated = errors.New("post comment not updated")
	ErrPostInvalidStyle      = errors.New("post invalid style")
)
