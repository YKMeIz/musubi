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
	"errors"
	"fmt"

	"github.com/YKMeIz/akari"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(revokeCmd)
}

var revokeCmd = &cobra.Command{
	Use:   "revoke [name|token]",
	Short: "Revoke an exist user",
	Long: `Revoke (musubi revoke) will delete an exist user from database.

Username or user's token is accepted.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 0:
			return errors.New("Missing an exist username or token.")
		case 1:
			return revoke(args[0])
		default:
			return errors.New("More than one argument is not accepted.")
		}
	},
}

func revoke(arg string) error {
	var err error
	u := &akari.User{Name: arg}
	t := &akari.User{Token: arg}
	if u.IsUser() {
		u.UserCompletion()
		err = u.RevokeUser()
		if err != nil {
			return err
		}
		fmt.Println("User " + u.Name + " and related token " + u.Token + " have been revoked.")
	} else if t.IsUser() {
		t.UserCompletion()
		err = t.RevokeUser()
		if err != nil {
			return err
		}
		fmt.Println("User " + t.Name + " and related token " + t.Token + " have been revoked.")
	} else {
		return errors.New("No such a user.")
	}
	return err
}
