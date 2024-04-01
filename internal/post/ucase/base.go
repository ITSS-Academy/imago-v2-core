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

func (p PostUseCase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return post.ErrPostRequiredID
	}
	return p.postRepo.Delete(ctx, id)
}

func (p PostUseCase) List(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*post.Post], error) {
	err := p.OptsValidate(opts)
	if err != nil {
		return nil, err
	}
	result, err := p.postRepo.List(ctx, opts)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (p PostUseCase) GetByUid(ctx context.Context, uid string, opts *common.QueryOpts, style string) (*common.ListResult[*post.Post], error) {
	err := p.OptsValidate(opts)
	if err != nil {
		return nil, err
	}
	if style == "mine" {
		result, err := p.postRepo.GetMine(ctx, uid, opts)
		if err != nil {
			return nil, err

		}
		return result, nil
	} else if style == "share" {
		result, err := p.postRepo.GetShared(ctx, uid, opts)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, post.ErrPostInvalidStyle

}

func (p PostUseCase) GetOther(ctx context.Context, uid string, opts *common.QueryOpts) (*common.ListResult[*post.Post], error) {
	err := p.OptsValidate(opts)
	if err != nil {
		return nil, err
	}
	result, err := p.postRepo.GetOther(ctx, uid, opts)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (p PostUseCase) GetDetail(ctx context.Context, id string) (*post.Post, error) {
	if id == "" {
		return nil, post.ErrPostRequiredID
	}

	return p.postRepo.GetDetail(ctx, id)
}

func (p PostUseCase) GetByCategory(ctx context.Context, categoryId string, opts *common.QueryOpts) (*common.ListResult[*post.Post], error) {
	err := p.OptsValidate(opts)
	if err != nil {
		return nil, err
	}
	result, err := p.postRepo.GetByCategory(ctx, categoryId, opts)
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

func (p PostUseCase) OptsValidate(opts *common.QueryOpts) error {
	if opts.Page <= 0 {
		return post.ErrPostInvalidPage
	}
	if opts.Size <= 0 {
		return post.ErrPostInvalidSize
	}
	return nil

}

func NewPostUseCase(postRepo post.PostRepository) *PostUseCase {
	return &PostUseCase{
		postRepo: postRepo,
	}
}
