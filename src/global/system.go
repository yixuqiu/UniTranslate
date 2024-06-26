package global

import (
	"errors"
	"uniTranslate/src/types"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/xgd16/gf-x-tool/xstorage"
)

// SystemConfig 系统配置信息
var SystemConfig *gjson.Json

var CacheMode = "mem"
var CachePlatform = false

// InitSystemConfig 初始化系统配置信息
func InitSystemConfig() {
	cfg, err := g.Cfg().Data(gctx.New())
	if err != nil {
		panic("初始化系统配置错误: " + err.Error())
	}
	SystemConfig = gjson.New(cfg, true)
	// 初始化配置的缓存模式
	CacheMode = SystemConfig.Get("server.cacheMode", "mem").String()
	CachePlatform = SystemConfig.Get("server.cachePlatform", false).Bool()
	CacheRefreshOnStartup = SystemConfig.Get("server.cacheRefreshOnStartup", false).Bool()
	ServiceKey = SystemConfig.Get("server.key").String()
	KeyMode = SystemConfig.Get("server.keyMode", 1).Int()
	ConfigDeviceMode = SystemConfig.Get("server.configDeviceMode", "xdb").String()
	ConfigDeviceMySqlDb = SystemConfig.Get("server.configDeviceMySqlDb", "default").String()
	CacheWriteToStorage = SystemConfig.Get("server.cacheWriteToStorage", false).Bool()
	RequestRecordKeepDays = SystemConfig.Get("server.requestRecordKeepDays", 7).Int()
}

// XDB 文件式存储
var XDB = xstorage.CreateXDB()

// ConfigDeviceMode 配置驱动模式
var ConfigDeviceMode = "xdb"

// ConfigDeviceMySqlDb 配置驱动模式 MySQL 驱动db设置
var ConfigDeviceMySqlDb = "default"

// ConfigDevice 配置驱动
var ConfigDevice types.ConfigDeviceInterface

// GetConfigDevice 获取驱动配置
func GetConfigDevice() (device types.ConfigDeviceInterface, err error) {
	if ConfigDevice == nil {
		err = errors.New("配置获取驱动尚未初始化")
		return
	}
	device = ConfigDevice
	return
}

// MustGetConfigDevice 忽略错误获取驱动配置
func MustGetConfigDevice() (device types.ConfigDeviceInterface) {
	device, err := GetConfigDevice()
	if err != nil {
		panic(err)
	}
	return
}

// CacheRefreshOnStartup 启动时是否从数据库刷新缓存 (会先清除缓存里所有的 缓存 在从数据库逐条初始化 数据 慎用!!!)
var CacheRefreshOnStartup = false

// ServiceKey 服务 key
var ServiceKey string

// KeyMode 密钥验证模式
var KeyMode int

// GfCache 全局缓存
var GfCache *gcache.Cache

var StatisticalProcess types.StatisticsInterface = new(types.MySqlStatistics)



// 是否将缓存写入存储
var CacheWriteToStorage = false

// 保留几天的请求记录
var RequestRecordKeepDays = 7
