package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lethe/common"
	"github.com/lethe/common/util"
	"github.com/lethe/dao/kv"
	"github.com/lethe/dao/mysql"
	"github.com/sirupsen/logrus"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Login(c *gin.Context) {
	req := mysql.UserInfo{}
	c.BindJSON(&req)
	if req.Name == "" {
		c.JSON(200, "username is null")
		return
	}
	user := mysql.GetUserByName(c.Request.Context(), req.Name)
	if user.Password != req.Password {
		c.JSON(http.StatusOK, common.ErrorResp(common.PasswordError, nil))
		return
	}
	c.SetCookie(common.UserId, strconv.Itoa(user.Id), 0, "/", "0.0.0.0", false, true)
	c.JSON(http.StatusOK, common.SuccessResp(user.Id))
}

func Logout(c *gin.Context) {
	c.SetCookie(common.UserId, "1", -1, "/", "0.0.0.0", false, true)
	c.JSON(http.StatusOK, common.SuccessResp(nil))
}

func Register(c *gin.Context) {
	req := struct{
		Name 	 string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email 	 string `json:"email" binding:"required"`
		Code 	 string `json:"code"`
	}{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(200, common.ErrorResp(common.ParamsError,nil))
		return 
	}
	// code, err := kv.RedisClient.Get(req.Email).Result()
	// if err == redis.Nil {
	// 	c.JSON(200, common.ErrorResp(common.RegisterTimeout, nil))
	// 	return
	// } else if err != nil {
	// 	c.JSON(200, common.ErrorResp(common.ServiceError, nil))
	// 	return 
	// } else if code != req.Code {
	// 	c.JSON(200, common.ErrorResp(common.RegisterCodeError, nil))
	// 	return
	// }
	mysql.CreateUserInfo(c.Request.Context(), mysql.UserInfo{
		Name: req.Name,
		Password: req.Password,
		Email: req.Email,
		Status: mysql.UserStatusOk,
	})
	c.JSON(200, common.SuccessResp(nil))
}

func SendEmail(c *gin.Context) {
	req := mysql.UserInfo{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(200, common.ErrorResp(common.ParamsError,nil))
		return 
	}
	verificationCode := util.Md5(util.RandString(10))
	body := fmt.Sprintf("your verification code is %s, please register in one hour!", verificationCode)
	err := util.SendEmail(req.Email, body)
	if err != nil {
		logrus.Errorf("send mail: %v", err)
		c.JSON(200, common.ErrorResp(common.ServiceError, nil))
		return 
	}
	kv.RedisClient.SetNX(req.Email, verificationCode, time.Hour)
	c.JSON(200, common.SuccessResp(nil))
}

func ChangeInfo(c *gin.Context){
	type ChangeInfoReq struct{
		UserId 		int 	`json:"user_id"`
		OldPassword string 	`json:"old_password"`
		NewPassword string 	`json:"new_password"`
	}
	req := new(ChangeInfoReq)
	if err := c.BindJSON(&req); err != nil {
		c.JSON(200, common.ErrorResp(common.ParamsError,nil))
		return 
	}
	ctx := c.Request.Context()
	user := mysql.GetUserById(ctx, req.UserId)
	if user.Password != req.OldPassword {
		c.JSON(200, common.ErrorResp(common.PasswordError, nil))
		return 
	}
	mysql.UpdatePasswordById(ctx, req.UserId, req.NewPassword)
	c.JSON(200, common.SuccessResp(nil))
}
