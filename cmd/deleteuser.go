// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// deleteuserCmd represents the deleteuser command
var deleteuserCmd = &cobra.Command{
	Use:   "deleteuser",
	Short: "Delete current user",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteuser called")

		user, flag := service.GetCurUser()
		if flag == false {
			fmt.Println("Please login!")
		} else {
			flag2 := service.DeleteUser(user.Name)
			if flag2 == true {
				fmt.Println("Successfully Delete")
			} else {
				fmt.Println("Error!")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteuserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
