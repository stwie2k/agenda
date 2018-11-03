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
	"agenda/entity"
	"agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// querymeetingCmd represents the querymeeting command
var querymeetingCmd = &cobra.Command{
	Use:   "querymeeting",
	Short: "Query meeting",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("querymeeting called")

		starttime, _ := cmd.Flags().GetString("starttime")
		endtime, _ := cmd.Flags().GetString("endtime")
		if starttime == "" || endtime == "" {
			fmt.Println("Please input the start time and the end time")
			return
		}
		user, flag := service.GetCurUser()
		if flag == false {
			fmt.Println("Please login!")
		} else {
			allMeeting, flag2 := service.QueryMeeting(user.Name, starttime, endtime)
			if flag2 == false {
				fmt.Println("Wrong Date!")
			} else {
				for _, m := range allMeeting {
					fmt.Println("----------------")
					fmt.Println("Title: ", m.Title)
					st, _ := entity.DateToString(m.StartDate)
					fmt.Println("Start Time", st)
					et, _ := entity.DateToString(m.EndDate)
					fmt.Println("End Time", et)
					fmt.Printf("Participators: ")
					for _, p := range m.Participators {
						fmt.Printf(p, " ")
					}
					fmt.Printf("\n")
					fmt.Println("----------------")
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(querymeetingCmd)

	// Here you will define your flags and configuration settings.

	querymeetingCmd.Flags().StringP("starttime", "s", "", "the start time of the meeting")
	querymeetingCmd.Flags().StringP("endtime", "e", "", "the end time of the meeting")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// querymeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// querymeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
