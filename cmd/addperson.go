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
// test

package cmd

import (
	"agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// addpersonCmd represents the addperson command
var addpersonCmd = &cobra.Command{
	Use:   "addperson",
	Short: "Add participators to a meeting",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addperson called")

		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetStringSlice("participator")
		if title == "" {
			fmt.Println("Please input the title")
			return
		}
		if len(participators) == 0 {
			fmt.Println("Please input the participators")
			return
		}
		if user, flag := service.GetCurUser(); flag != true {
			fmt.Println("Please login!")
		} else {
			flag := service.AddMeetingParticipator(user.Name, title, participators)
			if flag == true {
				fmt.Println("Add participators success!")
			} else {
				fmt.Println("Error!")
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(addpersonCmd)

	// Here you will define your flags and configuration settings.
	addpersonCmd.Flags().StringSliceP("participator", "p", nil, "the participators of meeting, such as \"name1, name2\"")
	addpersonCmd.Flags().StringP("title", "t", "", "the title of meeting")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addpersonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addpersonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
