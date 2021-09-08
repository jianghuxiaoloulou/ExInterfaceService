package global

import (
	"WowjoyProject/ExInterfaceService/pkg/logger"
	"WowjoyProject/ExInterfaceService/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	GeneralSetting  *setting.GeneralSettingS
	ReadCardSetting *setting.ReadCardSettingS
	Logger          *logger.Logger
)
