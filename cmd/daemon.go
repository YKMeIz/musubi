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

	"github.com/YKMeIz/akari"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(daemonCmd)
}

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run musubi in daemon mode",
	Long:  `daemon (musubi daemon) will run musubi service in daemon mode.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return daemon()
		}
		return errors.New("more than one argument is not accepted.")
	},
}

func daemon() error {
	a := akari.New()
	readConfig(a)
	a.Run()

	return nil
}

func readConfig(a *akari.Core) {
	a.DatabasePath = viper.GetString("databasePath")

	if s := viper.GetString("domainName"); s != "" {
		a.Domain = s
	}
	if s := viper.GetString("portNumber"); s != "" {
		a.Port = s
	}
	if s := viper.GetString("certChain"); s != "" {
		a.CertChain = s
	}
	if s := viper.GetString("certKey"); s != "" {
		a.CertKey = s
	}
	if s := viper.GetString("messageRelativePath"); s != "" {
		a.MessageRelativePath = s
	}
	if s := viper.GetString("websocketRelativePath"); s != "" {
		a.WebsocketRelativePath = s
	}
	if s := viper.GetString("pushbullet.token"); s != "" {
		pushbulletToken = s
	}
}
