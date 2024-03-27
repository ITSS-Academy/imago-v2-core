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

func (p PostBaseInterop) List(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*post.Post], error) {
	return p.postUseCase.List(ctx, opts)
}

func (p PostBaseInterop) Create(ctx context.Context, token string, data *post.Post) error {
	record, err := p.authUseCase.Verify(ctx, token)
	if err != nil {
		return err
	}
	data.CreatorId = record.UID
	data.ID = record.UID[:10] + strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	data.Comment = make([]string, 0)
	data.Like = make([]string, 0)
	if data.HashTag == nil {
		data.HashTag = make([]string, 0)
	}
	if data.Mention == nil {
		data.Mention = make([]string, 0)
	}
	data.Share = make([]string, 0)
	data.Status = "active"
	return p.postUseCase.Create(ctx, data)
}

//func (p PostBaseInterop) UpdatePostComment(ctx context.Context, token string, id string, data *post.Post) error {
//	_, err := p.authUseCase.Verify(ctx, token)
//	if err != nil {
//		return err
//	}
//	//postCommentData, err := p.postUseCase.GetPostById(ctx, id)
//	if err != nil {
//		return err
//	}
//	if data.CreatorId != data.CreatorId {
//		return post.ErrPostCommentNotUpdated
//	}
//	data.ID = id
//	return p.postUseCase.UpdatePostComment(ctx, id, data.Comment)
//}

func NewPostBaseInterop(postUseCase post.PostUseCase, authUseCase auth.AuthUseCase) PostBaseInterop {
	return PostBaseInterop{
		postUseCase: postUseCase,
		authUseCase: authUseCase,
	}
}