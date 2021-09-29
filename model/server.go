package model

import "strings"

type ServerList struct {
	Count int      `json:"count"`
	Doms  []string `json:"doms"`
}

func (s ServerList) String() string {
	return strings.Join(s.Doms, "|")
}

type ServerDetail struct {
	List  []SubServerDetail `json:"list"`
	Count int               `json:"count"`
}

type SubServerDetail struct {
	InstanceId  string `json:"instanceId"`
	Ip          string `json:"ip"`
	Port        int    `json:"port"`
	Enabled     bool   `json:"enabled"`
	Healthy     bool   `json:"healthy"`
	ClusterName string `json:"clusterName"`
}

type ServerView struct {
	ServerName  string `json:"serverName"`
	Ip          string `json:"ip"`
	Port        int    `json:"port"`
	Enabled     bool   `json:"enabled"`
	Healthy     bool   `json:"healthy"`
	ClusterName string `json:"clusterName"`
	Index       int    `json:"index"`
	Count       int    `json:"count"`
}

func NewServerView(serverName string, ip string, port int, enable bool, healthy bool, clusterName string, index int, count int) ServerView {
	return ServerView{
		ServerName:  serverName,
		Ip:          ip,
		Port:        port,
		Enabled:     enable,
		Healthy:     healthy,
		ClusterName: clusterName,
		Index:       index,
		Count:       count,
	}
}
