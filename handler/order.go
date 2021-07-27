package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/satori/go.uuid"
	"github.com/gin-gonic/gin"
	"github.com/lethe/common"
	"github.com/lethe/dao/kv"
	"github.com/lethe/dao/mysql"
	"github.com/sirupsen/logrus"
)

func OrderList(c *gin.Context){
	req := new(mysql.OrderReq)
	if err := c.BindJSON(req); err !=nil {
		logrus.Info("err: ",err)
		c.JSON(200, common.ErrorResp(common.ParamsError, nil))
		return
	}
	if req.Page == 0{
		req.Page = 1
		req.PageSize = 20
	}
	data := mysql.GetOrderList(c.Request.Context(), req)
	c.JSON(http.StatusOK, common.SuccessResp(data))
}

func CreateOrder(c *gin.Context){
	userId, _ := c.Cookie(common.UserId)
	Id, _ := strconv.Atoi(userId)
	req := new(mysql.OrderDetail)
	if err := c.BindJSON(req); err !=nil || req.OrderId == 0 || req.OrderNum == 0{
		logrus.Info("err: ",err)
		c.JSON(200, common.ErrorResp(common.ParamsError, nil))
		return
	}
	req.UserId = Id
	ctx := c.Request.Context()
	value := uuid.NewV4().String()
	kvKey := "create_order"+strconv.Itoa(req.OrderId)
	kv.RedisClient.GetLock(kvKey, value, time.Second*5)
	defer kv.RedisClient.ReleaseLock(kvKey, value)
	order := mysql.GetOrderById(ctx, req.OrderId)
	if order.Count >= req.OrderNum {
		err := mysql.CreateOrder(ctx, req)
		if err != nil {
			logrus.Info("create err: ",err)
			c.JSON(http.StatusOK, common.ErrorResp(common.DataBaseError,nil))
			return
		}
		err = mysql.UpdateOrderNum(ctx, order.Id, order.Count-req.OrderNum)
		if err != nil {
			logrus.Info("update err: ",err)
			c.JSON(http.StatusOK, common.ErrorResp(common.DataBaseError,nil))
			return
		}
		c.JSON(http.StatusOK, common.SuccessResp(nil))
	}else{
		logrus.Info("err: order not enough")
		c.JSON(http.StatusOK, common.ErrorResp(common.OrderNotEnoughError,nil))
	}
}
