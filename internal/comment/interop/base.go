package interop

import (
	"context"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/auth"
	"github.com/itss-academy/imago/core/domain/comment"
	"time"
)

type CommentInterop struct {
	ucase     comment.CommentUseCase
	authUcase auth.AuthUseCase
}

func (c CommentInterop) CreateComment(ctx context.Context, token string, commentData *comment.Comment) error {
	record, err := c.authUcase.Verify(ctx, token)
	if err != nil {
		return err
	}
	commentData.CreatorID = record.UID
	currentTime := time.Now()
	formattedTime := currentTime.Format("20060102150405")
	commentData.ID = formattedTime + commentData.CreatorID
	return c.ucase.CreateComment(ctx, commentData)
}

func (c CommentInterop) GetCommentById(ctx context.Context, token string, id string) (*comment.Comment, error) {
	_, err := c.authUcase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return c.ucase.GetCommentById(ctx, id)
}

func (c CommentInterop) GetCommentByPostId(ctx context.Context, token string, postId string, opts *common.QueryOpts) (*common.ListResult[*comment.Comment], error) {
	_, err := c.authUcase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return c.ucase.GetCommentByPostId(ctx, postId, opts)
}

func (c CommentInterop) GetComment(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*comment.Comment], error) {
	_, err := c.authUcase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return c.ucase.GetComment(ctx, opts)
}

func (c CommentInterop) UpdateComment(ctx context.Context, token string, id string, commentUpdate *comment.Comment) error {
	_, err := c.authUcase.Verify(ctx, token)
	if err != nil {
		return err
	}
	commentData, err := c.ucase.GetCommentById(ctx, id)
	if err != nil {
		return err
	}
	if commentUpdate.CreatorID != commentData.CreatorID {
		return comment.ErrCommentNotUpdated
	}
	commentUpdate.ID = id
	return c.ucase.UpdateComment(ctx, id, commentUpdate)
}

func (c CommentInterop) DeleteComment(ctx context.Context, token string, id string) error {
	_, err := c.authUcase.Verify(ctx, token)
	if err != nil {
		return err
	}
	_, err = c.ucase.GetCommentById(ctx, id)
	if err != nil {
		return err
	}
	return c.ucase.DeleteComment(ctx, id)
}

func NewCommentInterop(ucase comment.CommentUseCase, authUcase auth.AuthUseCase) CommentInterop {
	return CommentInterop{
		ucase:     ucase,
		authUcase: authUcase,
	}
}
