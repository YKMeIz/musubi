// Copyright © 2016 nrechn <nrechn@gmail.com>
//
// This file is part of musubi.
//
// musubi is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// musubi is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with musubi. If not, see <http://www.gnu.org/licenses/>.
//

package cmd

import (
	"fmt"

	"github.com/YKMeIz/akari"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(registerCmd)
}

var registerCmd = &cobra.Command{
	Use:   "register [name]",
	Short: "Register a new user",
	Long:  `Register (musubi register) will create a new user to database.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			er("Missing name for new user.")
		case 1:
			register(args[0])
		default:
			er("More than one argument is not accepted.")
		}
	},
}

func register(name string) {
	u := akari.User{Name: name}
	username, token, err := u.RegisterUser()
	if err != nil {
		er("An error occurred during user register.")
	}
	fmt.Println("Create User: " + "\033[36m\033[1m" + username + "\033[0m\033[39m")
	fmt.Println(name + "'s token is: " + "\033[32m\033[1m" + token + "\033[0m\033[39m")
}
