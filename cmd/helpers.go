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
)

const (
	configFilePath = "/etc/musubi/config.yml"
	configDirPath  = "/etc/musubi/"
	version        = "v0.1"
)

var (
	pushbulletToken string
)

func checkFile(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	return nil
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(-1)
}
