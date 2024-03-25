package Report

import (
	"context"
	"errors"
	"github.com/itss-academy/imago/core/common"
)

type Report struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Type      string `json:"type"`
	Reason    string `json:"reason"`
	Status    string `json:"status"`
	Content   string `json:"content"`
	CreatorID string `json:"creator_id"`
	UpdatedAt int64  `json:"updated_at"`
	CreatedAt int64  `json:"created_at"`
	TypeID    string `json:"type_id"`
}
type ReportRepository interface {
	Create(ctx context.Context, report *Report) error
	Get(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Report], error)
	GetById(ctx context.Context, id string) (*Report, error)
	GetAllByStatusCompleted(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Report], error)
	GetAllByStatusPending(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Report], error)
	Update(ctx context.Context, report *Report, id string) error
	Delete(ctx context.Context, id string) error
}

type ReportUseCase interface {
	Create(ctx context.Context, report *Report) error
	Get(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Report], error)
	GetById(ctx context.Context, id string) (*Report, error)
	GetAllByStatusCompleted(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Report], error)
	GetAllByStatusPending(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Report], error)
	Update(ctx context.Context, report *Report, id string) error
	Delete(ctx context.Context, id string) error
}

type ReportInterop interface {
	Create(ctx context.Context, token string, report *Report) error
	Get(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*Report], error)
	GetById(ctx context.Context, token string, id string) (*Report, error)
	GetAllByStatusCompleted(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*Report], error)
	GetAllByStatusPending(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*Report], error)
	Update(ctx context.Context, token string, report *Report, id string) error
	Delete(ctx context.Context, token string, id string) error
}

var (
	ErrReportNotFound    = errors.New("report not found")
	ErrReportNotValid    = errors.New("report not valid")
	ErrReportNotUpdated  = errors.New("report not updated")
	ErrReportNotCreated  = errors.New("report not created")
	ErrInvalidReportPage = errors.New("invalid report page")
	ErrInvalidReportSize = errors.New("invalid report size")
	ErrTypeEmpty         = errors.New("type is empty")
	ErrReasonEmpty       = errors.New("reason is empty")
	ErrContentEmpty      = errors.New("content is empty")
	ErrCreatorIDEmpty    = errors.New("creator id is empty")
	ErrTypeIDEmpty       = errors.New("type id is empty")
	ErrIDEmpty           = errors.New("id is empty")
	ErrStatusEmpty       = errors.New("status is empty")
)
