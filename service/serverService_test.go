package service

import (
	"github.com/spf13/viper"
	"testing"
)

func TestGetServerList(t *testing.T) {
	viper.Set("host", "10.10.20.90:8848")
	viper.Set("username", "nacos")
	viper.Set("password", "nacos")
	a := GetServerList(1)
	viper.Set("clusterName", "DEFAULT")
	viper.Set("pageSize", "10")
	toke := GetAccessToken()
	//ServerListService(a, toke)
	serName:="ycy-jenkins-test"
	SetServiceStatusService(serName,toke,"172.19.69.196",8899,true)
	ServerListService(a, toke)
}
