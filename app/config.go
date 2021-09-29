package app

import (
	"fmt"
	"github.com/modood/table"
	"github.com/spf13/viper"
	"github.com/yanchengyou/nacos-util/model"
	"github.com/yanchengyou/nacos-util/util"
)


var configjson model.Config

func SetValue(key, value string) error {
	err := util.CreateFile(viper.GetString("config"))
	if err != nil {
		return err
	}
	viper.Set(key, value)
	err = viper.WriteConfig()
	return err
}

func GetValue(key string) string {
	return viper.GetString(key)
}

func View() {
	if err := viper.ReadInConfig(); err == nil {
		err := viper.Unmarshal(&configjson)
		if err != nil {
			panic(err)
		}
		a := make([]model.Config, 1, 1)
		a[0]=configjson
		table.Output(a)
	} else {
		fmt.Println(err)
	}
}

func Connect() () {
	err := viper.Unmarshal(&configjson)
	if err != nil {
		panic(err)
	}
	params := fmt.Sprintf("username=%s&password=%s", configjson.Username, configjson.Password)
	url := fmt.Sprintf("http://%s%s", configjson.Host, model.LoginUrl)
	header := make(map[string]string)
	header["Content-Type"] = "application/x-www-form-urlencoded"
	response, err := util.HttpRequest("POST", url, params, header)
	if err != nil {
		panic(err)
	}
	if response.StatusCode == 200 {
		fmt.Println("测试连接 nacos 成功。")
	} else {
		fmt.Println("测试连接 nacos 失败。")
	}
}

