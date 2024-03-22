package comment

import (
	"context"
	"errors"
	"github.com/itss-academy/imago/core/common"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        string `json:"id" gorm:"primaryKey"`
	Content   string `json:"content"`
	CreatorID string `json:"creator_id"`
	PostID    string `json:"post_id"`
}

type CommentRepository interface {
	CreateComment(ctx context.Context, comment *Comment) error
	GetCommentById(ctx context.Context, id string) (*Comment, error)
	GetCommentByPostId(ctx context.Context, postId string, opts *common.QueryOpts) (*common.ListResult[*Comment], error)
	GetComment(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Comment], error)
	UpdateComment(ctx context.Context, comment *Comment) error
	DeleteComment(ctx context.Context, id string) error
}

type CommentUseCase interface {
	CreateComment(ctx context.Context, comment *Comment) error
	GetCommentById(ctx context.Context, id string) (*Comment, error)
	GetCommentByPostId(ctx context.Context, postId string, opts *common.QueryOpts) (*common.ListResult[*Comment], error)
	GetComment(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Comment], error)
	UpdateComment(ctx context.Context, comment *Comment) error
	DeleteComment(ctx context.Context, id string) error
}

type CommentInterop interface {
	CreateComment(ctx context.Context, token string, comment *Comment) error
	GetCommentById(ctx context.Context, token string, id string) (*Comment, error)
	GetCommentByPostId(ctx context.Context, token string, postId string, opts *common.QueryOpts) (*common.ListResult[*Comment], error)
	GetComment(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*Comment], error)
	UpdateComment(ctx context.Context, token string, comment *Comment) error
	DeleteComment(ctx context.Context, token string, id string) error
}

var (
	ErrCommentIdEmpty      = errors.New("comment id is empty")
	ErrCommentPostId       = errors.New("comment post id is empty")
	ErrCommentContentEmpty = errors.New("comment content is empty")
	ErrCommentNotFound     = errors.New("comment not found")
	ErrCommentNotValid     = errors.New("comment not valid")
	ErrCommentNotCreated   = errors.New("comment not created")
	ErrCommentNotUpdated   = errors.New("comment not updated")
	ErrCommentNotDeleted   = errors.New("comment not deleted")
	ErrInvalidCommentPage  = errors.New("invalid comment page")
	ErrInvalidCommentSize  = errors.New("invalid comment size")
)
