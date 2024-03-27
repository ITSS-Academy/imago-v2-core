package interop

import (
	"context"
	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/auth"
	"github.com/itss-academy/imago/core/domain/report"
	"time"
)

type ReportInterop struct {
	ucase     report.ReportUseCase
	authUcase auth.AuthUseCase
}

func (r ReportInterop) Create(ctx context.Context, token string, report *report.Report) error {
	record, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return err
	}
	report.CreatorID = record.UID
	currentTime := time.Now()
	formattedTime := currentTime.Format("20060102150405")
	report.ID = formattedTime + report.CreatorID
	report.Status = "pending"
	data := r.ucase.Create(ctx, report)
	if data != nil {
		return data
	}
	return nil
}
func (r ReportInterop) GetById(ctx context.Context, token string, id string) (*report.Report, error) {
	_, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return r.ucase.GetById(ctx, id)

}

func (r ReportInterop) Get(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*report.Report], error) {
	_, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return r.ucase.Get(ctx, opts)

}

func (r ReportInterop) GetAllByStatusApproved(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*report.Report], error) {

	_, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return r.ucase.GetAllByStatusApproved(ctx, opts)
}
func (r ReportInterop) GetAllByStatusPending(ctx context.Context, token string, opts *common.QueryOpts) (*common.ListResult[*report.Report], error) {

	_, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return r.ucase.GetAllByStatusPending(ctx, opts)
}

// Update report by id
func (r ReportInterop) Update(ctx context.Context, token string, reports *report.Report, id string) error {
	_, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return err
	}
	//check id exist
	_, err = r.ucase.GetById(ctx, id)
	if err != nil {
		return report.ErrReportNotFound
	}
	return r.ucase.Update(ctx, reports, id)
}

// Change status report to Approved
func (r ReportInterop) ChangeStatusApproved(ctx context.Context, token string, id string, status string) error {
	_, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return err
	}
	//check id exist
	data, err := r.ucase.GetById(ctx, id)
	if err != nil {
		return report.ErrReportNotFound

	}
	if data.Status == "pending" {
		data.Status = report.StatusApproved
		return r.ucase.Update(ctx, data, id)
	}
	return nil
}

// Change status report to rejected
func (r ReportInterop) ChangeStatusRejected(ctx context.Context, token string, id string, status string) error {
	_, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return err
	}
	//check id exist
	data, err := r.ucase.GetById(ctx, id)
	if err != nil {
		return report.ErrReportNotFound

	}
	if data.Status == "pending" {
		data.Status = report.StatusRejected
		return r.ucase.Update(ctx, data, id)
	}
	return nil

}

// Delete report by id
func (r ReportInterop) Delete(ctx context.Context, token string, id string) error {
	_, err := r.authUcase.Verify(ctx, token)
	if err != nil {
		return err
	}
	return r.ucase.Delete(ctx, id)

}
func NewReportInterop(ucase report.ReportUseCase, authUcase auth.AuthUseCase) *ReportInterop {
	return &ReportInterop{ucase: ucase, authUcase: authUcase}

}
