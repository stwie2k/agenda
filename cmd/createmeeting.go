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

// createmeetingCmd represents the createmeeting command
var createmeetingCmd = &cobra.Command{
	Use:   "createmeeting",
	Short: "Create a meeting",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createmeeting called")

		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetStringSlice("participator")
		starttime, _ := cmd.Flags().GetString("starttime")
		endtime, _ := cmd.Flags().GetString("endtime")
		if title == "" {
			fmt.Println("Please input the title")
			return
		}
		if len(participators) == 0 {
			fmt.Println("Please input the participator")
			return
		}
		if starttime == "" || endtime == "" {
			fmt.Println("Please input the date")
			return
		}
		user, flag := service.GetCurUser()
		if flag == false {
			fmt.Println("Please login!")
		} else {
			flag := service.CreateMeeting(user.Name, title, starttime, endtime, participators)
			if flag == true {
				fmt.Println("Create meeting successfully!")
			} else {
				fmt.Println("Error!")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(createmeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
