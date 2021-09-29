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
	"github.com/yanchengyou/nacos-util/app"
)

// setUsernameCmd represents the setUsername command
var setUsernameCmd = &cobra.Command{
	Use:   "set-username",
	Short: "配置nacos username",
	Long:  `配置nacos username`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("请输入username")
		} else {
			err := app.SetValue("username",args[0])
			if err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	configCmd.AddCommand(setUsernameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setUsernameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setUsernameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
