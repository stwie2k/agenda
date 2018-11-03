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

// quitmeetingCmd represents the quitmeeting command
var quitmeetingCmd = &cobra.Command{
	Use:   "quitmeeting",
	Short: "Quit a meeting",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("quitmeeting called")

		title, _ := cmd.Flags().GetString("title")
		if title == "" {
			fmt.Println("Please input the title")
			return
		}
		if user, flag := service.GetCurUser(); flag != true {
			fmt.Println("Please login!")
		} else {
			flag := service.QuitMeeting(user.Name, title)
			if flag == true {
				fmt.Println("Quit Successfully")
			} else {
				fmt.Println("Error! You're not a participator of the meeting")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(quitmeetingCmd)

	// Here you will define your flags and configuration settings.

	quitmeetingCmd.Flags().StringP("title", "t", "", "the title of meeting")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quitmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quitmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
