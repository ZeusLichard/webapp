package appinterface

import (
	"webapp/frame/appframework"
	"webapp/frame/subsys"
	"webapp/toolkit"

	"github.com/zanlichard/beegoe/validation"
)

//通用返回
type HttpCommonRsp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

//检查APP版本请求定义
type AppVersionCheckReq struct {
	ClientType     int8   `valid:"Required" json:"client_type"` //当前版本
	CurrentVersion string `valid:"Required" json:"current_ver"` //客户端端类型(1:ios,2:android,3:web)
}

type GetImageReq struct {
	FileKey  string `valid:"Required" json:"file_key"`  //文件key
	FileMd5  string `valid:"Required" json:"file_md5"`  //文件md5
	FileSize int32  `valid:"Required" json:"file_size"` //文件大小
}

type GetImageRsp struct {
	UserID  int32  `json:"uid"`
	FileURL string `json:"file_url"`
	AppID   int32  `json:"app_id"`
}

//检查APP版本响应定义
type AppVersionCheckRsp struct {
	BuildCode   string `json:"build_code"`   // 构建的代码
	DownloadUrl string `json:"download_url"` // 下载的url
	ForceUpdate uint8  `json:"force_update"` // 0否，1是
	VersionName string `json:"version_name"` // 版本名称
	Title       string `json:"title"`        // 标题
	Content     string `json:"content"`      // 内容
	Remark      string `json:"remark"`       // 备注
}

//基本请求体定义
type AppVerCheckMsg struct {
	Head  subsys.SubsysHeader `json:"_head"`
	Param AppVersionCheckReq  `json:"_param"` //上层应用定义
}

//基本请求体定义
type GetImageMsg struct {
	Head  subsys.SubsysHeader `json:"_head"`
	Param GetImageReq         `json:"_param"` //上层应用定义
}

func (t *AppVersionCheckReq) Valid(v *validation.Validation) {
	if t.ClientType != 1 && t.ClientType != 2 {
		v.SetError("ClientType", "ClientType有效期取值只能为1,2")
	}
	if len(t.CurrentVersion) != 6 {
		v.SetError("current_ver", "长度不合法")
	}
}

func (t *GetImageReq) Valid(v *validation.Validation) {
	if t.FileSize <= 0 {
		v.SetError("FileSize", "文件大小不能小于等于0")
	}

}

//基本配置管理接口定义(header)
type BasicCfgGetReq struct {
	CfgType string `valid:"Required" json:"cfg_type"` //rabbitmq,mysql,redis,mongo as the key
}

func (t *BasicCfgGetReq) Valid(v *validation.Validation) {
	supportedTypes := []string{"rabbitmq", "mysql", "redis", "mongo"}
	if !toolkit.ArrayCheckIn(t.CfgType, supportedTypes) {
		v.SetError("cfg_type", "不支持的类型")
	}
}

type BasicCfgGetRsp struct {
	UserName     string   `json:"user_name"`
	Passwd       string   `json:"pass_word"`
	Database     string   `json:"database_name"`
	Hosts        []string `json:"host_names"`
	MaxOpenConns int      `json:"max_open_conns"`
	MaxIdleConns int      `json:"max_idle_conns"`
	IdleTimeout  int      `json:"idle_timeout"`
	MaxActive    int      `json:"max_active"` //redis
	Other        string   `json:"extends"`    //rabbitmq-vhost-queue
}

//依赖配置管理接口定义(header)
type DepCfgGetReq struct {
	IsServicesAll bool   `form:"is_services_all" valid:"Required"` //是否全部读取,false,则需要指定service_name
	ServiceName   string `form:"service_name"`                     //服务名
}

func (t *DepCfgGetReq) Valid(v *validation.Validation) {
	if !t.IsServicesAll {
		if t.ServiceName == "" {
			v.SetError("service_name", "参数不全")
		}

	}
}

type DepCfgGetRsp struct {
	Services []appframework.AclDependentItem `json:"services"`
}

type LocalAclRsp struct {
	LocalCfg appframework.LocalAcl `json:"local_config"`
}
