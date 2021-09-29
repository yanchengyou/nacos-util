/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// offlineCmd represents the offline command
var offlineCmd = &cobra.Command{
	Use:   "offline",
	Short: "服务下线操作",
	Long:  `对服务进行下线处理，可以通过 -s 指定单个服务，也可以通过 -f 指定服务列表文件，-h 指定ip地址`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Flags())
		//app.ServerOffline()
	},
}

func init() {
	serverCmd.AddCommand(offlineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// offlineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// offlineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//offlineCmd.Flags().BoolP("all", "a", false, "是否下线所有服务")
	offlineCmd.Flags().StringP("service", "s", "", "需要下线的服务")
	offlineCmd.Flags().StringP("file", "f", "", "需要下线的服务文件列表，文件列表格式为serviceName:Host:port,如果port为空则默认下线host上所有服务。")
	offlineCmd.Flags().StringP("host", "H", "", "服务的服务器地址")
	offlineCmd.Flags().Int64P("port", "p", 0, "服务的端口")
	if err := viper.BindPFlags(offlineCmd.Flags()); err != nil {
		panic(nil)
	}

}
