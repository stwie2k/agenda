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

var (
	username *string
	password *string
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")

		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		if username == "" {
			fmt.Println("Please input username")
			return
		}
		if password == "" {
			fmt.Println("Please input password")
			return
		}
		if _, flag := service.GetCurUser(); flag == true {
			fmt.Println("Please logout!")
			return
		}
		flag2 := service.UserLogin(username, password)
		if flag2 == true {
			fmt.Println("Login Successfully. Username: ", username)
		} else {
			fmt.Println("Wrong username or password!")
		}
		return
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
	username = loginCmd.Flags().StringP("username", "u", "", "Your username")
	password = loginCmd.Flags().StringP("password", "p", "", "Your password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
