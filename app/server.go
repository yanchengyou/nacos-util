package app

import (
	"github.com/spf13/viper"
	"github.com/yanchengyou/nacos-util/service"
)

func ServerList() {
	//token:=service.GetAccessToken()
	// 1 需要改成参数的
	svl := service.GetServerList(1)
	token := service.GetAccessToken()
	service.ServerListService(svl, token)
}

func ServerOffline() {
	serviceName := viper.GetString("service")
	//file := viper.GetString("file")
	host := viper.GetString("host")
	port := viper.GetInt("port")
	token := service.GetAccessToken()
	if serviceName==""{
		return
	}
	service.SetServiceStatusService(serviceName, token, host, port, false)

}
