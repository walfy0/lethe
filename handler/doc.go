package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lethe/common"
	"github.com/lethe/dao/mysql"
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
		c.JSON(200, common.ErrorResp(common.ParamsError, nil))
	}
	ctx := c.Request.Context()
	if req.Id == 0 {
		mysql.CreateDoc(ctx, req)
	} else {
		mysql.UpdateDoc(ctx, req)
	}
	c.JSON(http.StatusOK,common.SuccessResp(nil))
}