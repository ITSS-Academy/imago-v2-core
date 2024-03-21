package usecase

import (
	"context"
	firebaseAuth "firebase.google.com/go/v4/auth"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/Report"
)

type ReportUseCase struct {
	repo       Report.ReportRepository
	authClient *firebaseAuth.Client
}

func (r ReportUseCase) Create(ctx context.Context, reportData *Report.Report) error {
	err := r.Validate(reportData)
	if err != nil {
		return err
	}
	err = r.Create(ctx, reportData)
	if err != nil {
		return Report.ErrReportNotCreated
	}
	return nil
}

func (r ReportUseCase) Get(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Report.Report], error) {
	if opts.Page < 1 {
		return nil, Report.ErrInvalidReportPage
	}
	if opts.Size < 0 {
		return nil, Report.ErrInvalidReportSize
	}
	return r.repo.Get(ctx, opts)
}

func (r ReportUseCase) GetById(ctx context.Context, id string) (*Report.Report, error) {
	data, err := r.repo.GetById(ctx, id)
	if err != nil {
		return nil, Report.ErrReportNotFound
	}
	return data, nil
}

func (r ReportUseCase) GetAllByStatusCompleted(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Report.Report], error) {
	data, err := r.repo.GetAllByStatusCompleted(ctx, opts)
	if err != nil {
		return nil, Report.ErrReportNotFound
	}
	return data, nil
}

func (r ReportUseCase) GetAllByStatusPending(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Report.Report], error) {
	data, err := r.repo.GetAllByStatusPending(ctx, opts)
	if err != nil {
		return nil, Report.ErrReportNotFound
	}
	return data, nil
}

func (r ReportUseCase) Update(ctx context.Context, report *Report.Report) error {
	data := r.repo.Update(ctx, report)
	if data != nil {
		return Report.ErrReportNotUpdated
	}
	return data
}

func (r ReportUseCase) Validate(data *Report.Report) error {
	if data.ID == "" {
		return Report.ErrIDEmpty
	}
	if data.Content == "" {
		return Report.ErrContentEmpty
	}
	if data.Type == "" {
		return Report.ErrTypeEmpty
	}
	if data.TypeID == "" {
		return Report.ErrTypeIDEmpty
	}
	if data.Reason == "" {
		return Report.ErrReasonEmpty
	}
	if data.Status == "" {
		return Report.ErrStatusEmpty
	}
	if data.CreatorID == "" {
		return Report.ErrCreatorIDEmpty
	}
	return nil

}

func NewReportUseCase(repo Report.ReportRepository, authClient *firebaseAuth.Client) Report.ReportUseCase {
	return &ReportUseCase{
		repo:       repo,
		authClient: authClient,
	}
}
