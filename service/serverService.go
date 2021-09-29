package service

import (
	"encoding/json"
	"github.com/modood/table"
	"strconv"

	//"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/yanchengyou/nacos-util/model"
	"github.com/yanchengyou/nacos-util/util"
)

var config model.Config

func GetAccessToken() string {
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	params := fmt.Sprintf("username=%s&password=%s", config.Username, config.Password)
	header := make(map[string]string)
	header["Content-Type"] = "application/x-www-form-urlencoded"
	response, err := util.HttpRequest("POST", util.UrlJoin(model.LoginUrl), params, header)
	if err != nil {
		panic(err)
	}

	rspByte, err := util.ParseRspToString(200, response)
	if err != nil {
		panic(err)
	}
	rspMap := make(map[string]interface{})
	err = json.Unmarshal(rspByte, &rspMap)
	if err != nil {
		panic(err)
	}
	token := rspMap["accessToken"].(string)
	return token
}

func GetServerList(pageNo int) model.ServerList {
	serverList := new(model.ServerList)
	params := fmt.Sprintf("pageNo=%d&pageSize=%d&accessToken=%s", pageNo, model.PageSize, GetAccessToken())
	url := util.UrlJoin(fmt.Sprintf("%s?%s", model.ServerListUrl, params))
	rsp, err := util.HttpRequest("GET", url, "", nil)
	if err != nil {
		panic(err)
	}
	serverListByte, err := util.ParseRspToString(200, rsp)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(serverListByte, serverList)
	if err != nil {
		panic(err)
	}
	return *serverList
}

func ServerListService(serverList model.ServerList, accessToken string) []model.ServerView {
	ServerViewSlice := make([]model.ServerView, 0, 0)
	for _, serverName := range serverList.Doms {
		serverDetail := getServerDetail(serverName, accessToken)
		for index, value := range serverDetail.List {
			serverView := model.NewServerView(serverName, value.Ip, value.Port, value.Enabled, value.Healthy, value.ClusterName, index+1, serverDetail.Count)
			ServerViewSlice = append(ServerViewSlice, serverView)
		}
	}
	table.Output(ServerViewSlice)
	return ServerViewSlice
}

func getServerViewConfig() string {
	clusterName := viper.GetString("clusterName")
	groupName := viper.GetString("groupName")
	pageSize := viper.GetInt("pageSize")
	namespaceId := viper.GetString("namespaceId")
	return fmt.Sprintf("clusterName=%s&pageSize=%d&groupName=%s&namespaceId=%s", clusterName, pageSize, groupName, namespaceId)
}

func getServerDetail(serverName, accessToken string) model.ServerDetail {
	params := fmt.Sprintf("accessToken=%s&pageNo=%d&serviceName=%s&%s", accessToken, 1, serverName, getServerViewConfig())
	url := util.UrlJoin(fmt.Sprintf("%s?%s", model.ServerDetailUrl, params))
	rsp, err := util.HttpRequest("GET", url, "", nil)
	if err != nil {
		panic(err)
	}
	a, err := util.ParseRspToString(200, rsp)
	if err != nil {
		panic(err)
	}
	serverDetail := new(model.ServerDetail)
	err = json.Unmarshal(a, serverDetail)
	if err != nil {
		panic(err)
	}
	return *serverDetail
}

func SetServiceStatusService(serviceName, accessToken string, ip string, port int, enable bool) {
	serverList := getServerDetail(serviceName, accessToken).List
	for _, server := range serverList {

		if ip != "" && port != 0 {
			if server.Ip == ip && server.Port == port && enable != server.Enabled {
				setServiceResult(serviceName, setServiceEnable(serviceName, ip, port, accessToken, enable), enable)
			}
		} else if ip != "" && port == 0 && enable != server.Enabled {
			if server.Ip == ip {
				setServiceResult(serviceName, setServiceEnable(serviceName, ip, server.Port, accessToken, enable), enable)
			}
		} else if ip == "" && port != 0 && enable != server.Enabled {
			if server.Port == port {
				setServiceResult(serviceName, setServiceEnable(serviceName, server.Ip, port, accessToken, enable), enable)
			}
		} else if enable != server.Enabled {
			setServiceResult(serviceName, setServiceEnable(serviceName, server.Ip, server.Port, accessToken, enable), enable)
		}
	}
}

func setServiceEnable(serviceName string, ip string, port int, token string, enable bool) string {
	url := fmt.Sprintf("http://%s%s?%s", viper.GetString("host"), model.InstanceUrl, token)
	params := fmt.Sprintf("serviceName=%s&ip=%s&port=%d&namespaceId=%s&enabled=%t", serviceName, ip, port, viper.GetString("namespaceId"), enable)
	header := make(map[string]string)
	header["Content-Type"] = "application/x-www-form-urlencoded"
	rsp, err := util.HttpRequest("POST", url, params, header)
	if err != nil {
		panic(err)
	}
	rspByte, err := util.ParseRspToString(200, rsp)
	if err != nil {
		panic(err)
	}
	return string(rspByte)
}

func setServiceResult(serverName, res string, enable bool) {
	if "ok" == res {
		fmt.Println(fmt.Sprintf("服务：%s，设置状态：%t，成功", serverName, enable))
	} else {
		fmt.Println(fmt.Sprintf("服务：%s，设置状态：%t，失败", serverName, enable))
	}
}

type serverArg struct {
	ip      string
	ports   []int
	isEmpty bool
}

// 可能不用这个了
func serverArgs(args ...string) serverArg {
	var ip string
	var ports []int
	isEmpty := true

	if len(args) == 1 {
		ip = args[0]
		isEmpty = false
	} else if len(args) > 1 {
		isEmpty = true
		ip = args[0]
		for _, portStr := range args[1:] {
			port, err := strconv.Atoi(portStr)
			if err != nil {
				panic(err)
			}
			ports = append(ports, port)
		}

	}
	return serverArg{ip: ip, ports: ports, isEmpty: isEmpty}
}
