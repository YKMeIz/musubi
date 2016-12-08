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

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Musubi's version information",
	Long:  `Musubi's version information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Musubi Message Server " + "\033[35m\033[1m" + version + "\033[0m\033[39m")
	},
}
