package mysql

import (
	"context"
	"time"

	"github.com/lethe/config"
)

type OrderInfo struct{
	Id			int 		`gorm:"primaryKey" json:"id"`
	Status 		int 		`gorm:"column:status" json:"status"`
	Count 		int 		`gorm:"column:count" json:"count"`
	Name		string 		`gorm:"column:name" json:"name"`
	Ctime 		*time.Time 	`gorm:"column:ctime;default:null" json:"ctime"`
	Mtime 		*time.Time 	`gorm:"column:mtime;default:null" json:"mtime"`
}

type OrderReq struct{
	Page 		int  `json:"page"`
	PageSize 	int  `json:"page_size"`
}

type OrderDetail struct{
	Id			int 		`gorm:"primaryKey" json:"id"`
	UserId		int			`gorm:"column:user_id" json:"user_id"`
	OrderId		int 		`gorm:"column:order_id" json:"order_id"`
	OrderNum	int 		`gorm:"column:order_num" json:"order_num"`
	Ctime 		*time.Time 	`gorm:"column:ctime;default:null" json:"ctime"`
	Mtime 		*time.Time 	`gorm:"column:mtime;default:null" json:"mtime"`
}

var OrderStatusOk = 1

func GetOrderList(ctx context.Context, req *OrderReq) []*OrderInfo {
	query := config.MysqlClient.WithContext(ctx).Model(&OrderInfo{})
	query = query.Where("status = ?", OrderStatusOk)
	query = query.Offset((req.Page-1)*req.PageSize).Limit(req.PageSize)
	var list []*OrderInfo
	query.Find(&list) 
	return list
}

func GetOrderById(ctx context.Context, id int) *OrderInfo {
	query := config.MysqlClient.WithContext(ctx).Model(&OrderInfo{})
	query = query.Where("id = ?", id)
	var order *OrderInfo
	query.Find(&order)
	return order
}

func UpdateOrderNum(ctx context.Context, id, num int) error {
	query := config.MysqlClient.WithContext(ctx).Model(&OrderInfo{})
	query = query.Where("id = ?", id)
	return query.Update("count",num).Error
}

func CreateOrder(ctx context.Context, order *OrderDetail) error {
	query := config.MysqlClient.WithContext(ctx).Model(&OrderDetail{})
	return query.Create(order).Error
}
