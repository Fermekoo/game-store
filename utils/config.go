package utils

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	IsProduction      bool   `mapstructure:"IS_PRODUCTION"`
	VIPBaseURL        string `mapstructure:"VIP_BASE_URL"`
	VIPApiID          string `mapstructure:"VIP_API_ID"`
	VIPApiKey         string `mapstructure:"VIP_API_KEY"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
	ServiceFee        uint   `mapstructure:"SERVICE_FEE"`
	MidtransServerKey string `mapstructure:"MIDTRANS_SERVER_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
	})

	err = viper.Unmarshal(&config)
	return
}
