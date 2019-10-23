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
	"fmt"
	"os"
	"zz/agenda/models"
	"zz/agenda/entity"
	"github.com/spf13/cobra"
)

var registerUser models.User

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "This command can register user",
	Long: `You can use agenda register to sign up one user.`,
	Run: func(cmd *cobra.Command, args []string) {
		models.Logger.SetPrefix("[agenda register]")
		users := entity.ReadUserInfoFromFile()
		for _, user := range users {
			if user.Username == registerUser.Username {
				models.Logger.Println(registerUser.Username, "has been registered!")
				fmt.Println(registerUser.Username, "has been registered!")
				os.Exit(0)
			}
		}
		users = append(users, registerUser)
		entity.WriteUserInfoToFile(users)
		models.Logger.Println("Register", registerUser.Username, "successfully!")
		fmt.Println("Register", registerUser.Username, "successfully!")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&registerUser.Username, "username", "u", "", "The User's Username")
	registerCmd.Flags().StringVarP(&registerUser.Password, "password", "p", "", "The User's Password")
	registerCmd.Flags().StringVarP(&registerUser.Email, "email", "e", "", "The User's Email")
	registerCmd.Flags().StringVarP(&registerUser.Telephone, "telephone", "t", "", "The User's telephone")
	registerCmd.MarkFlagRequired("username")
	registerCmd.MarkFlagRequired("password")
}
