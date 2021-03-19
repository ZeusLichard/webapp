package app

import (
	"context"
	"webapp/appframework"
	"webapp/appframework/code"
	"webapp/subsys"
	"webapp/toolkit"

	"github.com/gin-gonic/gin"
	"github.com/zanlichard/beegoe/validation"
)

func MarkErrors(ctx context.Context, errors []*validation.Error) {
	for _, err := range errors {
		appframework.BusinessLogger.Error(ctx, err.Key, err.Message)
	}
	return
}

func JsonResponse(ctx *gin.Context, httpCode, retCode int, rspHead subsys.SubsysHeader, data interface{}) {
	dataInfo := &subsys.SubsysCommonRsp{
		Msg:     code.GetMsg(retCode),
		Data:    data,
		RetCode: toolkit.ConvertToString(retCode),
		Ret:     toolkit.ConvertToString(retCode),
	}
	ctx.JSON(httpCode, gin.H{
		"_head": rspHead,
		"_data": dataInfo,
	})
}

func JsonResponsev2(ctx *gin.Context, httpCode, retCode int, data interface{}) {
	ctx.JSON(httpCode, gin.H{
		"code": retCode,
		"msg":  code.GetMsg(retCode),
		"data": data,
	})
}

func ProtoBufResponse(ctx *gin.Context, httpCode int, data interface{}) {
	ctx.ProtoBuf(httpCode, data)
}
