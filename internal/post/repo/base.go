package repo

import (
	"context"
	"math"

	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/post"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func (p PostRepository) Create(ctx context.Context, post *post.Post) error {
	tx := p.db.Create(post)
	return tx.Error
}

//func (p PostRepository) GetById(ctx context.Context, id string) (*post.Post, error) {
//	found := &post.Post{}
//	tx := p.db.WithContext(ctx).Where("id = ?", id).First(&found)
//	if tx.Error != nil {
//		if tx.Error.Error() == "not found" {
//			return nil, post.ErrPostNotFound
//		}
//		return nil, tx.Error
//	}
//	return found, nil
//}
//
//func (p PostRepository) GetByUid(ctx context.Context, uid string, opts *common.QueryOpts) (*common.ListResult[*post.Post], error) {
//
//	panic("implement me")
//}

func (p PostRepository) List(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*post.Post], error) {
	offset := (opts.Page - 1) * opts.Size
	result := make([]*post.Post, 0)
	tx := p.db.WithContext(ctx).Offset(offset).Limit(opts.Size).Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	count := int64(0)
	tx = p.db.WithContext(ctx).Model(&post.Post{}).Count(&count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	page := int(math.Ceil(float64(count) / float64(opts.Size)))
	return &common.ListResult[*post.Post]{
		Data:    result,
		EndPage: int(page),
	}, nil
}

//func (p PostRepository) UpdatePostComment(ctx context.Context, id string, comment []string) error {
//	tx := p.db.WithContext(ctx).Model(&post.Post{}).Where("id = ?", id).Update("comment", comment)
//	return tx.Error
//}

func NewPostRepository(db *gorm.DB) *PostRepository {
	err := db.AutoMigrate(&post.Post{})
	if err != nil {
		panic(err)
	}
	return &PostRepository{db: db}
}
