package repo

import (
	"context"
	"github.com/itss-academy/imago/core/domain/report"
	"math"

	"github.com/itss-academy/imago/core/common"
	"gorm.io/gorm"
)

type ReportRepository struct {
	db *gorm.DB
}

func (r ReportRepository) Create(ctx context.Context, report *report.Report) error {
	tx := r.db.Create(report)
	return tx.Error
}

func (r ReportRepository) GetById(ctx context.Context, id string) (*report.Report, error) {
	reportData := &report.Report{}
	tx := r.db.Where("id = ?", id).First(reportData)
	return reportData, tx.Error
}

func (r ReportRepository) Get(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*report.Report], error) {
	reportData := make([]*report.Report, 0)
	limit := opts.Size
	offset := opts.Size * (opts.Page - 1)
	tx := r.db.Limit(limit).Offset(offset).Find(&reportData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	count := int64(0)
	tx = r.db.Model(&report.Report{}).Count(&count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	pageNum := int(math.Ceil(float64(count) / float64(limit)))
	return &common.ListResult[*report.Report]{Data: reportData, EndPage: pageNum}, tx.Error
}

func (r ReportRepository) GetAllByStatusApproved(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*report.Report], error) {
	reportData := make([]*report.Report, 0)
	limit := opts.Size
	offset := opts.Size * (opts.Page - 1)
	tx := r.db.Where("status = ?", "approved").Limit(limit).Offset(offset).Find(&reportData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	count := int64(0)
	tx = r.db.Model(&report.Report{}).Where("status = ?", "approved").Count(&count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	pageNum := int64(0)
	if count > 0 {
		pageNum = int64(math.Ceil(float64(count) / float64(limit)))
	}
	return &common.ListResult[*report.Report]{Data: reportData, EndPage: int(pageNum)}, tx.Error

}

func (r ReportRepository) GetAllByStatusPending(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*report.Report], error) {
	reportData := make([]*report.Report, 0)
	limit := opts.Size
	offset := opts.Size * (opts.Page - 1)
	tx := r.db.Where("status = ?", "pending").Limit(limit).Offset(offset).Find(&reportData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	count := int64(0)
	tx = r.db.Model(&report.Report{}).Where("status = ?", "pending").Count(&count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	pageNum := int64(0)
	if count > 0 {
		pageNum = int64(math.Ceil(float64(count) / float64(limit)))
	}
	return &common.ListResult[*report.Report]{Data: reportData, EndPage: int(pageNum)}, tx.Error
}

func (r ReportRepository) Update(ctx context.Context, report *report.Report, id string) error {
	tx := r.db.Where("id = ?", id).Save(report)
	return tx.Error
}

// DeleteReport delete report by id
func (r ReportRepository) Delete(ctx context.Context, id string) error {
	tx := r.db.Where("id = ?", id).Delete(&report.Report{})
	return tx.Error

}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	err := db.AutoMigrate(&report.Report{})
	if err != nil {
		panic(err)
	}
	return &ReportRepository{db: db}

}
