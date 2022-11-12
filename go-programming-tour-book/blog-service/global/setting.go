package global

import "blog-service/pkg/setting"

// 全局变量,设置一个指向各元素对应的指针
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)
