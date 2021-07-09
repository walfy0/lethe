package mysql

import (
	"context"
	"time"

	"github.com/lethe/config"
)

type DocInfo struct{
	Id			int 		`gorm:"primaryKey" json:"id"`
	Status 		int 		`gorm:"column:status" json:"status"`
	Header 		string 		`gorm:"column:header" json:"header"`
	Content		string 		`gorm:"column:content" json:"content"`
	Creator 	string 		`gorm:"column:creator" json:"creator"`
	Ctime 		*time.Time 	`gorm:"column:ctime;default:null" json:"ctime"`
	Mtime 		*time.Time 	`gorm:"column:mtime;default:null" json:"mtime"`
}

type DocReq struct{
	Page 		int  `json:"page"`
	PageSize 	int  `json:"page_size"`
	Id 			*int `json:"id"`
}

func (DocInfo)GetName() string {
	return "doc_info"
}

func GetDocList(ctx context.Context,req DocReq) []*DocInfo {
	query := config.MysqlClient.WithContext(ctx)
	if req.Id != nil {
		query.Where("id = ?", req.Id)
	}
	query.Offset((req.Page-1)*req.PageSize).Limit(req.PageSize)
	var list []*DocInfo
	query.Find(&list)
	return list
}

func CreateDoc(ctx context.Context,doc DocInfo) error {
	query := config.MysqlClient.WithContext(ctx)
	return query.Create(&doc).Error
}

func UpdateDoc(ctx context.Context, doc DocInfo) error {
	query := config.MysqlClient.WithContext(ctx)
	return query.Where("id = ?", doc.Id).Updates(doc).Error
}
