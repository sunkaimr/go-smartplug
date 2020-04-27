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
package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"

	"SmartPlug/models"
	_ "SmartPlug/routers"
)

func init(){
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLevel(logs.LevelDebug)
	logs.EnableFuncCallDepth(true)

	beego.BConfig.AppName = "SmartPlug"
	beego.BConfig.RunMode = "dev"
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.Listen.HTTPAddr = "localhost"
	beego.BConfig.Listen.HTTPPort = 80
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.SetStaticPath("/", "static")
}

func main() {
	err := models.CheckTimerData()
	if err != nil {
		logs.Error("CheckTimerData failed, err:%s", err.Error())
		return
	}

	logs.Info("CheckTimerData succesd")
	beego.Run()
}
