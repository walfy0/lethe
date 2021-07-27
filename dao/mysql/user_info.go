package mysql

import (
	"context"
	"time"

	"github.com/lethe/config"
)

type UserInfo struct {
	Id 			int 		`gorm:"primaryKey" json:"id"`
	Status 		int			`gorm:"column:status" json:"status"`
	Name		string		`gorm:"column:name" json:"name"`
	Password 	string 		`gorm:"column:password" json:"password"`
	Email 		string 		`gorm:"column:email" json:"email"`
	Ctime 		*time.Time 	`gorm:"column:ctime;default:null" json:"ctime"`
	Mtime 		*time.Time 	`gorm:"column:mtime;default:null" json:"mtime"`
}

var (
	UserStatusOk  = 1 // using
	UserStatusBan = 2 // register
)

func (UserInfo)GetName() string {
	return "user_info"
}

func CreateUserInfo(ctx context.Context,user UserInfo) error{
	query := config.MysqlClient.WithContext(ctx).Create(&user)
	return query.Error
}

func GetUserByName(ctx context.Context, name string) UserInfo{
	query := config.MysqlClient.WithContext(ctx)
	query = query.Model(&UserInfo{}).Where("name = ?", name)
	query = query.Where("status = ?", UserStatusOk)
	var user UserInfo
	query.Find(&user)
	return user
}

func GetUserById(ctx context.Context, id int) UserInfo{
	query := config.MysqlClient.WithContext(ctx)
	query = query.Model(&UserInfo{}).Where("id = ?", id)
	var user UserInfo
	query.Find(&user)
	return user
}

func UpdatePasswordById(ctx context.Context, id int, password string) error {
	query := config.MysqlClient.WithContext(ctx)
	query = query.Model(&UserInfo{}).Where("id = ?", id)
	return query.Update("password", password).Error
}
