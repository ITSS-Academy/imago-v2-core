package usecase

import (
	"context"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/report"
)

type ReportUseCase struct {
	repo report.ReportRepository
}

func (r ReportUseCase) Create(ctx context.Context, reportData *report.Report) error {
	err := r.Validate(reportData)
	if err != nil {
		return err
	}
	err = r.repo.Create(ctx, reportData)
	if err != nil {
		return report.ErrReportNotCreated
	}
	return nil
}

func (r ReportUseCase) Get(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*report.Report], error) {
	if opts.Page < 1 {
		return nil, report.ErrInvalidReportPage
	}
	if opts.Size < 0 {
		return nil, report.ErrInvalidReportSize
	}
	return r.repo.Get(ctx, opts)
}

func (r ReportUseCase) GetById(ctx context.Context, id string) (*report.Report, error) {
	data, err := r.repo.GetById(ctx, id)
	if err != nil {
		return nil, report.ErrReportNotFound
	}
	return data, nil
}

func (r ReportUseCase) GetAllByStatusApproved(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*report.Report], error) {
	data, err := r.repo.GetAllByStatusApproved(ctx, opts)
	if err != nil {
		return nil, report.ErrReportNotFound
	}
	return data, nil
}

func (r ReportUseCase) GetAllByStatusPending(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*report.Report], error) {
	data, err := r.repo.GetAllByStatusPending(ctx, opts)
	if err != nil {
		return nil, report.ErrReportNotFound
	}
	return data, nil
}

// Update report by id
func (r ReportUseCase) Update(ctx context.Context, reportData *report.Report, id string) error {
	err := r.Validate(reportData)
	if err != nil {
		return err
	}
	err = r.repo.Update(ctx, reportData, id)
	if err != nil {
		return report.ErrReportNotUpdated
	}
	return nil

}

// Delete report by id
func (r ReportUseCase) Delete(ctx context.Context, id string) error {
	err := r.repo.Delete(ctx, id)
	if err != nil {
		return report.ErrReportNotFound
	}
	return nil

}

func (r ReportUseCase) Validate(data *report.Report) error {
	if data.ID == "" {
		return report.ErrIDEmpty
	}
	if data.Content == "" {
		return report.ErrContentEmpty
	}
	if data.Type == "" {
		return report.ErrTypeEmpty
	}
	if data.TypeID == "" {
		return report.ErrTypeIDEmpty
	}
	if data.Reason == "" {
		return report.ErrReasonEmpty
	}
	if data.Status == "" {
		data.Status = report.StatusPending
	}
	if data.CreatorID == "" {
		return report.ErrCreatorIDEmpty
	}
	return nil

}

func NewReportUseCase(repo report.ReportRepository) *ReportUseCase {
	return &ReportUseCase{
		repo: repo,
	}
}
