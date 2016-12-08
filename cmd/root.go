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
	"os"

	"github.com/nrechn/akari"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "musubi",
	Short: "A message server for IoT communication and notification push",
	Long:  `Musubi is a message server designed for IoT communication and notification push from *nix side to any device.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initDatabase)
}

// initConfig reads the config file.
func initConfig() {
	if err := checkFile(configFilePath); err != nil {
		fmt.Println("Error: Musubi config file is not found.")
		os.Exit(1)
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(configDirPath)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initDatabase initializes the database.
func initDatabase() {
	if !viper.IsSet("databasePath") {
		fmt.Println("Error: SQLite database file's path is not set.")
		os.Exit(1)
	}
	if err := checkFile(viper.GetString("databasePath")); err != nil {
		fmt.Println("\033[31m\033[1mWarn: SQLite database file is not found.\033[0m\033[39m")
		akari.InitDatabase(viper.GetString("databasePath"))
		fmt.Println("\033[33m\033[1mWarn: A new SQLite database file is created by Musubi.\033[0m\033[39m")
	} else {
		c := akari.Core{DatabasePath: viper.GetString("databasePath")}
		c.OpenDatabase()
	}
}
