package interop

import (
	"context"
	"fmt"
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
	_, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return err
	}
	fmt.Print(r.ucase)
	report.ID = time.Now().String()
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
func (r ReportInterop) Update(ctx context.Context, token string, reportData *Report.Report) error {

	_, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return err
	}
	return r.ucase.Update(ctx, reportData)

}

func NewReportInterop(ucase Report.ReportUseCase) *ReportInterop {
	return &ReportInterop{ucase: ucase}

}
