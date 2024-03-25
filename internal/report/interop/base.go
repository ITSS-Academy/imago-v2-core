package interop

import (
	"context"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/Report"
	"github.com/itss-academy/imago/core/domain/auth"
	"time"
)

type ReportInterop struct {
	ucase     Report.ReportUseCase
	authUcase auth.AuthUseCase
}

func (r ReportInterop) Create(ctx context.Context, token string, report *Report.Report) error {
	record, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return err
	}
	report.CreatorID = record.UID
	currentTime := time.Now()
	formattedTime := currentTime.Format("20060102150405")
	report.ID = formattedTime + report.CreatorID
	data := r.ucase.Create(ctx, report)
	if data != nil {
		return data
	}
	return nil
}
func (r ReportInterop) GetById(ctx context.Context, token string, id string) (*Report.Report, error) {
	_, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return r.ucase.GetById(ctx, id)

}

func (r ReportInterop) Get(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*Report.Report], error) {
	_, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return r.ucase.Get(ctx, opts)

}

func (r ReportInterop) GetAllByStatusCompleted(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*Report.Report], error) {

	_, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return r.ucase.GetAllByStatusCompleted(ctx, opts)
}
func (r ReportInterop) GetAllByStatusPending(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*Report.Report], error) {

	_, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return r.ucase.GetAllByStatusCompleted(ctx, opts)
}

// Update report by id
func (r ReportInterop) Update(ctx context.Context, token string, report *Report.Report, id string) error {
	_, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return err
	}
	//check id exist
	_, err = r.ucase.GetById(ctx, id)
	if err != nil {
		return Report.ErrReportNotFound
	}
	return r.ucase.Update(ctx, report, id)
}

// Delete report by id
func (r ReportInterop) Delete(ctx context.Context, token string, id string) error {
	_, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return err
	}
	return r.ucase.Delete(ctx, id)

}
func NewReportInterop(ucase Report.ReportUseCase, authUcase auth.AuthUseCase) *ReportInterop {
	return &ReportInterop{ucase: ucase, authUcase: authUcase}

}
