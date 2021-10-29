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
	etcd "omc/cmd/etcd"

	"github.com/spf13/cobra"
)

func etcdStatusCommand() {
	etcdFolderPath := currentContextPath + "/etcd_info/"
	etcd.EndpointStatus(etcdFolderPath)
}

// etcdCmd represents the etcd command
var etcdStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Etcd status",
	Run: func(cmd *cobra.Command, args []string) {
		etcdStatusCommand()
	},
}

func init() {
	etcdCmd.AddCommand(etcdStatusCmd)
}
