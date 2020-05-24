/*
 * Copyright sunkai
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package config

import (
	"fmt"
	"github.com/astaxie/beego"
)

var(
	LogLevel = ""
	DBDriver = ""
	Database = ""
	DBaddr = ""
	DBport = ""
	DBuser = ""
	DBpasswd = ""
)

func InitConfig() error {
	LogLevel = beego.AppConfig.DefaultString("LogLevel", "LevelInfo")

	if DBDriver = beego.AppConfig.String("dbdriver"); DBDriver == "" {
		fmt.Printf("dbdriver is none")
		return fmt.Errorf("dbdriver is none")
	}

	if Database = beego.AppConfig.String("database"); Database == "" {
		fmt.Printf("database is none")
		return fmt.Errorf("database is none")
	}
	if DBaddr = beego.AppConfig.String("dbaddr"); DBaddr == ""{
		fmt.Printf("dbaddr is none")
		return fmt.Errorf("dbaddr is none")
	}

	if DBport = beego.AppConfig.String("dbport"); DBport == ""{
		fmt.Printf("dbport is none")
		return fmt.Errorf("dbport is none")
	}

	if DBuser = beego.AppConfig.String("dbuser"); DBuser == ""{
		fmt.Printf("dbuser is none")
		return fmt.Errorf("dbuser is none")
	}

	if DBpasswd = beego.AppConfig.String("dbpasswd"); DBpasswd == ""{
		fmt.Printf("dbpasswd is none")
		return fmt.Errorf("dbpasswd is none")
	}

	return nil
}
