package v1

import (
	"webapp/appinterface"
	"webapp/errors"
	"webapp/globalconfig"
	"webapp/pkg/app"
	"webapp/pkg/code"
	"webapp/service"
	"webapp/stat"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"time"
)

const (
	StatGetAppVersion  = "GetAppVersion"
)

func CheckAppVersionApi(c *gin.Context)  {
	t1 := time.Now()
	var form appinterface.AppVersionCheckReq
	err := app.BindAndValid(c, &form)
	ipSrc := net.ParseIP(c.Request.RemoteAddr)
	payload := int(c.Request.ContentLength)
	if err != nil {
		app.JsonResponse(c,http.StatusBadRequest,code.INVALID_PARAMS,err.Error())
		go stat.PushStat(StatGetAppVersion, int(time.Now().Sub(t1).Seconds()*1000), ipSrc, payload, int(code.INVALID_PARAMS))
		return
	}

	result,retCode := service.GetAppVersion(c,&form)
	if retCode != errors.RetCode_SUCCESS {
		globalconfig.ErrorLogger.Errorf(c,"GetAppVersion form: %+v, err: %v",form,err)
		app.JsonResponse(c,http.StatusOK,int(retCode),nil)
		go stat.PushStat(StatGetAppVersion, int(time.Now().Sub(t1).Seconds()*1000), ipSrc, payload, int(retCode))
		return
	}

	app.JsonResponse(c,http.StatusOK,code.SUCCESS,result)
	go stat.PushStat(StatGetAppVersion, int(time.Now().Sub(t1).Seconds()*1000), ipSrc, payload, int(errors.RetCode_SUCCESS))
}
