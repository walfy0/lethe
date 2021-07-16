package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lethe/common"
	"github.com/lethe/dao/mysql"
	"github.com/sirupsen/logrus"
)

func DocList(c *gin.Context){
	req := mysql.DocReq{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(200, common.ErrorResp(common.ParamsError, nil))
	}
	if req.Page == 0{
		req.Page = 1
		req.PageSize = 20
	}
	list := mysql.GetDocList(c.Request.Context(), req)
	c.JSON(http.StatusOK,common.SuccessResp(list))
}

func DocUpdate(c *gin.Context){
	req := mysql.DocInfo{}
	if err := c.BindJSON(&req); err != nil {
		logrus.Info("err: %v",err)
		c.JSON(200, common.ErrorResp(common.ParamsError, nil))
		return
	}
	ctx := c.Request.Context()
	var err error
	if req.Id == 0 {
		err = mysql.CreateDoc(ctx, req)
	} else {
		err = mysql.UpdateDoc(ctx, req)
	}
	if err != nil {
		logrus.Info("err: %v",err)
		c.JSON(http.StatusOK, common.ErrorResp(common.DataBaseError,nil))
		return
	}
	c.JSON(http.StatusOK,common.SuccessResp(nil))
}