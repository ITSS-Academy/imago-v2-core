package repo

import (
	"context"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/comment"
	"gorm.io/gorm"
	"math"
)

type CommentRepository struct {
	db *gorm.DB
}

func (c CommentRepository) CreateComment(ctx context.Context, comment *comment.Comment) error {
	tx := c.db.Create(comment)
	return tx.Error
}

func (c CommentRepository) GetCommentById(ctx context.Context, id string) (*comment.Comment, error) {
	commentData := &comment.Comment{}
	tx := c.db.Where("id = ?", id).First(commentData)
	return commentData, tx.Error
}

func (c CommentRepository) GetCommentByPostId(ctx context.Context, postId string, opts *common.QueryOpts) (*common.ListResult[*comment.Comment], error) {
	commentData := make([]*comment.Comment, 0)
	limit := opts.Size
	offset := opts.Size * (opts.Page - 1)
	tx := c.db.Where("post_id = ?", postId).Limit(limit).Offset(offset).Find(&commentData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	count := int64(0)
	tx = c.db.Model(&comment.Comment{}).Where("post_id = ?", postId).Count(&count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	pageNum := int64(0)
	if count > 0 {
		pageNum = int64(math.Ceil(float64(count) / float64(limit)))
	}
	return &common.ListResult[*comment.Comment]{Data: commentData, EndPage: int64(int(pageNum))}, tx.Error
}

func (c CommentRepository) GetComment(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*comment.Comment], error) {
	commentData := make([]*comment.Comment, 0)
	limit := opts.Size
	offset := opts.Size * (opts.Page - 1)
	tx := c.db.Limit(limit).Offset(offset).Find(&commentData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	count := int64(0)
	tx = c.db.Model(&comment.Comment{}).Count(&count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	pageNum := int(math.Ceil(float64(count) / float64(limit)))
	return &common.ListResult[*comment.Comment]{Data: commentData, EndPage: int64(pageNum)}, tx.Error
}

func (c CommentRepository) UpdateComment(ctx context.Context, id string, comment *comment.Comment) error {
	tx := c.db.Where("id = ?", id).Updates(comment)
	return tx.Error
}

func (c CommentRepository) DeleteComment(ctx context.Context, id string) error {
	tx := c.db.Where("id = ?", id).Delete(&comment.Comment{})
	return tx.Error
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	err := db.AutoMigrate(&comment.Comment{})
	if err != nil {
		panic(err)
	}
	return &CommentRepository{db: db}
}
