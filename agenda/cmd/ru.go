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
	"zz/agenda/models"
	"zz/agenda/entity"
	"github.com/spf13/cobra"
)

// ruCmd represents the ru command
var ruCmd = &cobra.Command{
	Use:   "ru",
	Short: "This command is used to clear the account for the user who has logged in.",
	Long: `You can use agenda ru to Clear your account information [use with caution]`,
	Run: func(cmd *cobra.Command, args []string) {
		models.Logger.SetPrefix("[agenda remove user]")
		isLoggedIn, user := entity.IsLoggedIn()
		if isLoggedIn == true {
			// delete login info
			entity.ClearCurUserInfo()

			// delete user info
			entity.RemoveUser(user.Username)

			models.Logger.Println(user.Username, "clear account")
			fmt.Println("Remove user ["+ user.Username + "] successfully")
		} else {
			fmt.Println("Please login first")
		}
	},
}

func init() {
	rootCmd.AddCommand(ruCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ruCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ruCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
