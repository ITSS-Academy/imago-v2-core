package interop

import (
	"context"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/auth"
	"github.com/itss-academy/imago/core/domain/post"
	"strconv"
	"time"
)

type PostBaseInterop struct {
	postUseCase post.PostUseCase
	authUseCase auth.AuthUseCase
}

func (p *PostBaseInterop) List(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*post.Post], error) {
	return p.postUseCase.List(ctx, opts)
}

func (p *PostBaseInterop) Create(ctx context.Context, token string, post *post.Post) error {
	record, err := p.authUseCase.Verify(ctx, token)
	if err != nil {
		return err
	}
	post.CreatorId = record.UID
	post.ID = record.UID[:10] + strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	//post.Comment = make([]string, 0)
	//post.Like = make([]string, 0)
	//post.Share = make([]string, 0)
	//post.Status = "active"
	//if post.Mention == nil {
	//	post.Mention = make([]string, 0)
	//}
	//if post.HashTag == nil {
	//	post.HashTag = make([]string, 0)
	//}

	post.CreatedAt = time.Now()

	return p.postUseCase.Create(ctx, post)
}

func NewPostBaseInterop(postUseCase post.PostUseCase, authUcase auth.AuthUseCase) *PostBaseInterop {
	return &PostBaseInterop{
		postUseCase: postUseCase,
		authUseCase: authUcase,
	}
}
