package admin

import (
	"net"
	"net/http"
	"time"
	"webapp/appframework"
	"webapp/appframework/app"
	"webapp/appframework/code"
	"webapp/appinterface"
	"webapp/apptoml"
	"webapp/errors"
	"webapp/stat"

	"github.com/gin-gonic/gin"
)

const (
	StatGetBasicCfg     = "GetBasicCfg"
	StatGetDependentCfg = "GetDepCfg"
)

func GetBasicConfig(c *gin.Context) {
	appframework.BusinessLogger.Infof(c, "content-type:%s", c.Request.Header.Get("Content-Type"))
	t1 := time.Now()
	var form appinterface.BasicCfgGetReq
	ipSrc := net.ParseIP(c.Request.RemoteAddr)
	payload := int(c.Request.ContentLength)
	err := app.BindAndValid(c, &form)
	appframework.BusinessLogger.Infof(c, "req body:%+v", form)
	if err != nil {
		app.JsonResponse(c, http.StatusBadRequest, code.INVALID_PARAMS, err.Error())
		appframework.ErrorLogger.Errorf(c, "GetBasicCfg form: %+v, err: %+v", form, err)
		go stat.PushStat(StatGetBasicCfg, int(time.Now().Sub(t1).Seconds()*1000), ipSrc, payload, int(code.INVALID_PARAMS))
		return
	}
	var result appinterface.BasicCfgGetRsp
	switch form.CfgType {
	case "mysql":
		result.UserName = apptoml.Config.Database.Mysql.User
		result.Passwd = apptoml.Config.Database.Mysql.Passwd
		result.Database = apptoml.Config.Database.Mysql.Database
		result.Hosts = []string{apptoml.Config.Database.Mysql.ServerAddr}
		result.MaxOpenConns = apptoml.Config.Database.Mysql.MaxOpenConns
		result.MaxIdleConns = apptoml.Config.Database.Mysql.MaxIdleConns
		result.IdleTimeout = apptoml.Config.Database.Mysql.IdleTimeout

	case "redis":
		result.Passwd = apptoml.Config.Redisinfo.Passwd
		result.Hosts = apptoml.Config.Redisinfo.ServerList
		result.MaxOpenConns = apptoml.Config.Redisinfo.MaxIdle
		result.IdleTimeout = apptoml.Config.Redisinfo.IdleTimeout
		result.MaxActive = apptoml.Config.Redisinfo.MaxActive

	case "rabbitmq":
		result.UserName = apptoml.Config.RabbitMq.Username
		result.Passwd = apptoml.Config.RabbitMq.Password
		result.Hosts = []string{apptoml.Config.RabbitMq.ServerAddr}
		result.Other = "{\"queuename\":apptoml.Config.RabbitMq.Queuename,\"vhost\":apptoml.Config.RabbitMq.Vhost}"

	default:
		result.UserName = ""
		result.Passwd = ""
		result.Database = ""
		result.MaxOpenConns = 0
		result.MaxIdleConns = 0
		result.IdleTimeout = 0
	}
	app.JsonResponse(c, http.StatusOK, code.SUCCESS, result)
	go stat.PushStat(StatGetBasicCfg, int(time.Now().Sub(t1).Seconds()*1000), ipSrc, payload, int(errors.RetCode_SUCCESS))
}

func GetDependentConfig(c *gin.Context) {
	appframework.BusinessLogger.Infof(c, "content-type:%s", c.Request.Header.Get("Content-Type"))
	t1 := time.Now()
	var form appinterface.DepCfgGetReq
	ipSrc := net.ParseIP(c.Request.RemoteAddr)
	payload := int(c.Request.ContentLength)
	err := app.BindAndValid(c, &form)
	appframework.BusinessLogger.Infof(c, "req body:%+v", form)
	if err != nil {
		app.JsonResponse(c, http.StatusBadRequest, code.INVALID_PARAMS, err.Error())
		appframework.ErrorLogger.Errorf(c, "GetDepCfg form: %+v, err: %+v", form, err)
		go stat.PushStat(StatGetDependentCfg, int(time.Now().Sub(t1).Seconds()*1000), ipSrc, payload, int(code.INVALID_PARAMS))
		return
	}
	var result appinterface.DepCfgGetRsp
	item := appinterface.ServiceItem{
		apptoml.Config.Server.ServiceName,
		"2160034",
		"jjjjfdsafdafdasf dsafds",
	}
	result.Services = append(result.Services, item)
	app.JsonResponse(c, http.StatusOK, code.SUCCESS, result)
	go stat.PushStat(StatGetDependentCfg, int(time.Now().Sub(t1).Seconds()*1000), ipSrc, payload, int(errors.RetCode_SUCCESS))

}

func SetBasicConfig(c *gin.Context) {

}