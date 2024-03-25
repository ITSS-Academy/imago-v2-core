package repo

import (
	"context"
	"math"

	"github.com/itss-academy/imago/core/common"
	"github.com/itss-academy/imago/core/domain/Report"
	"gorm.io/gorm"
)

type ReportRepository struct {
	db *gorm.DB
}

func (r ReportRepository) Create(ctx context.Context, report *Report.Report) error {
	tx := r.db.Create(report)
	return tx.Error
}

func (r ReportRepository) GetById(ctx context.Context, id string) (*Report.Report, error) {
	reportData := &Report.Report{}
	tx := r.db.Where("id = ?", id).First(reportData)
	return reportData, tx.Error
}

func (r ReportRepository) Get(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Report.Report], error) {
	reportData := make([]*Report.Report, 0)
	limit := opts.Size
	offset := opts.Size * (opts.Page - 1)
	tx := r.db.Limit(limit).Offset(offset).Find(&reportData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	count := int64(0)
	tx = r.db.Model(&Report.Report{}).Count(&count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	pageNum := int(math.Ceil(float64(count) / float64(limit)))
	return &common.ListResult[*Report.Report]{Data: reportData, EndPage: pageNum}, tx.Error
}

func (r ReportRepository) GetAllByStatusCompleted(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Report.Report], error) {
	reportData := make([]*Report.Report, 0)
	limit := opts.Size
	offset := opts.Size * (opts.Page - 1)
	tx := r.db.Where("status = ?", "completed").Limit(limit).Offset(offset).Find(&reportData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	count := int64(0)
	tx = r.db.Model(&Report.Report{}).Where("status = ?", "completed").Count(&count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	pageNum := int64(0)
	if count > 0 {
		pageNum = int64(math.Ceil(float64(count) / float64(limit)))
	}
	return &common.ListResult[*Report.Report]{Data: reportData, EndPage: int(pageNum)}, tx.Error

}

func (r ReportRepository) GetAllByStatusPending(ctx context.Context, opts *common.QueryOpts) (*common.ListResult[*Report.Report], error) {
	reportData := make([]*Report.Report, 0)
	limit := opts.Size
	offset := opts.Size * (opts.Page - 1)
	tx := r.db.Where("status = ?", "pending").Limit(limit).Offset(offset).Find(&reportData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	count := int64(0)
	tx = r.db.Model(&Report.Report{}).Where("status = ?", "pending").Count(&count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	pageNum := int64(0)
	if count > 0 {
		pageNum = int64(math.Ceil(float64(count) / float64(limit)))
	}
	return &common.ListResult[*Report.Report]{Data: reportData, EndPage: int(pageNum)}, tx.Error
}

func (r ReportRepository) Update(ctx context.Context, report *Report.Report, id string) error {
	tx := r.db.Model(&Report.Report{}).Where("id = ?", id).Updates(report)
	return tx.Error
}

// DeleteReport delete report by id
func (r ReportRepository) Delete(ctx context.Context, id string) error {
	tx := r.db.Where("id = ?", id).Delete(&Report.Report{})
	return tx.Error

}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	err := db.AutoMigrate(&Report.Report{})
	if err != nil {
		panic(err)
	}
	return &ReportRepository{db: db}

}
