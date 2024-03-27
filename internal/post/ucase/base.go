package ucase

import (
	"context"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/post"
)

type PostUseCase struct {
	postRepo post.PostRepository
}

func (p PostUseCase) Create(ctx context.Context, data *post.Post) error {
	err := p.Validate(data)
	if err != nil {
		return err
	}
	err = p.postRepo.Create(ctx, data)
	if err != nil {
		return post.ErrPostNotCreated
	}
	return nil
}

func (p PostUseCase) List(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*post.Post], error) {
	if opts.Page <= 0 {
		return nil, post.ErrPostInvalidPage
	}
	if opts.Size <= 0 {
		return nil, post.ErrPostInvalidSize
	}
	result, err := p.postRepo.List(ctx, opts)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//func (p PostUseCase) UpdatePostComment(ctx context.Context, id string, data *post.Post) error {
//	err := p.Validate(data)
//	if err != nil {
//		return err
//	}
//	err = p.postRepo.UpdatePostComment(ctx, id, data.Comment)
//	if err != nil {
//		return post.ErrPostCommentNotUpdated
//	}
//	return nil
//}

func (p PostUseCase) Validate(postData *post.Post) error {
	if postData.Content == "" {
		return post.ErrPostRequiredContent
	}
	if len(postData.PhotoUrl) == 0 {
		return post.ErrPostRequiredPhoto
	}
	return nil
}

func NewPostUseCase(postRepo post.PostRepository) *PostUseCase {
	return &PostUseCase{
		postRepo: postRepo,
	}
}