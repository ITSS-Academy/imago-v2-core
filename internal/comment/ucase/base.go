package ucase

import (
	"context"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/comment"
)

type CommentUseCase struct {
	repo comment.CommentRepository
}

func (c CommentUseCase) CreateComment(ctx context.Context, commentData *comment.Comment) error {
	err := c.Validate(commentData)
	if err != nil {
		return err
	}
	err = c.repo.CreateComment(ctx, commentData)
	if err != nil {
		return comment.ErrCommentNotCreated
	}
	return nil
}

func (c CommentUseCase) GetCommentById(ctx context.Context, id string) (*comment.Comment, error) {
	commentData, err := c.repo.GetCommentById(ctx, id)
	if err != nil {
		return nil, comment.ErrCommentNotFound
	}
	err = c.Validate(commentData)
	if err != nil {
		return nil, comment.ErrCommentNotValid
	}
	return commentData, nil
}

func (c CommentUseCase) GetCommentByPostId(ctx context.Context, postId string, opts *common.QueryOpts) (*common.ListResult[*comment.Comment], error) {
	commentData, err := c.repo.GetCommentByPostId(ctx, postId, opts)
	if err != nil {
		return nil, err
	}
	return commentData, nil

}

func (c CommentUseCase) GetComment(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*comment.Comment], error) {
	commentData, err := c.repo.GetComment(ctx, opts)
	if err != nil {
		return nil, err
	}
	return commentData, nil
}

func (c CommentUseCase) UpdateComment(ctx context.Context, commentData *comment.Comment) error {
	err := c.Validate(commentData)
	if err != nil {
		return err
	}
	err = c.repo.UpdateComment(ctx, commentData)
	if err != nil {
		return comment.ErrCommentNotUpdated
	}
	return nil
}

func (c CommentUseCase) DeleteComment(ctx context.Context, id string) error {
	err := c.repo.DeleteComment(ctx, id)
	if err != nil {
		return comment.ErrCommentNotDeleted
	}
	return nil
}

func (c CommentUseCase) Validate(commentData *comment.Comment) error {
	if commentData.PostID == "" {
		return comment.ErrCommentPostId
	}
	if commentData.Content == "" {
		return comment.ErrCommentContentEmpty
	}
	return nil
}

func NewCommentUseCase(repo comment.CommentRepository) *CommentUseCase {
	return &CommentUseCase{repo: repo}
}
