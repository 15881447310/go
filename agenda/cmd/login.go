/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"os"
	"fmt"
	"zz/agenda/models"
	"zz/agenda/entity"
	"github.com/spf13/cobra"
)

var loginUser models.User

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "This command can login user",
	Long: `You can use agenda login to login one user.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("login called")
		users := entity.ReadUserInfoFromFile()
		models.Logger.SetPrefix("[agenda login]")

		// 判断是否已经登陆过
		isLoggedIn, user := entity.IsLoggedIn()
		if isLoggedIn == true {
			// 已经登陆
			fmt.Println(user.Username + " has already in")
			os.Exit(0)
		}

		for _, userInfo := range users {
			if userInfo.Username == loginUser.Username && userInfo.Password == loginUser.Password {
				// users[i].Login = true
				// entity.WriteUserInfoToFile(users)
				entity.SaveCurUserInfo(userInfo)
				models.Logger.Println("Login", loginUser.Username, "successfully!")
				fmt.Println("Login successfully")
				os.Exit(0)
			} else {
				models.Logger.Println("Login", loginUser.Username, "error!")
				fmt.Println("Username or Password error, please check your input")
				os.Exit(0)
			}
		}
		models.Logger.Println("Login", loginUser.Username, "no such an user!")
		fmt.Println("No such an user")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&loginUser.Username, "username", "u", "", "The User's Username")
	loginCmd.Flags().StringVarP(&loginUser.Password, "password", "p", "", "The User's Password")
	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("password")
}
