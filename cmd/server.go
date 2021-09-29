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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "nacos server",
	Long:  `nacos server [list|online|offlie]`,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")
	serverCmd.PersistentFlags().StringP("clusterName", "", "DEFAULT", "clusterName ")
	serverCmd.PersistentFlags().StringP("groupName", "", "DEFAULT_GROUP", "groupName ")
	serverCmd.PersistentFlags().StringP("namespaceId", "", "", "namespaceId ")
	if err := viper.BindPFlag("clusterName", serverCmd.PersistentFlags().Lookup("clusterName")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("groupName", serverCmd.PersistentFlags().Lookup("groupName")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("namespaceId", serverCmd.PersistentFlags().Lookup("namespaceId")); err != nil {
		panic(err)
	}
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
