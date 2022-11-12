package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

// 设置配置文件
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/") // 设置相对路径
	vp.SetConfigType("yaml")     // 配置类型为yaml
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
