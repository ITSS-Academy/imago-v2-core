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
	if data.Content == "" {
		return post.ErrPostRequiredContent
	}
	if len(data.PhotoUrl) == 0 {
		return post.ErrPostRequiredPhoto
	}

	err := p.postRepo.Create(ctx, data)
	if err != nil {
		return err
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

func NewPostUseCase(postRepo post.PostRepository) *PostUseCase {
	return &PostUseCase{
		postRepo: postRepo,
	}
}
