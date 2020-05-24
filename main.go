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
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"smartplug/config"

	"smartplug/models"
	_ "smartplug/routers"
)

func init() {
	config.InitConfig()
	initLog()
	initDB()
	initWeb()
}

func main() {
	err := models.CheckTimerData()
	if err != nil {
		logs.Error("CheckTimerData failed, err:%s", err.Error())
		return
	}

	err = models.CheckDelayData()
	if err != nil {
		logs.Error("CheckDelayData failed, err:%s", err.Error())
		return
	}

	err = models.CheckInfraredData()
	if err != nil {
		logs.Error("CheckInfraredData failed, err:%s", err.Error())
		return
	}

	err = models.CheckMeterData()
	if err != nil {
		logs.Error("CheckMeterData failed, err:%s", err.Error())
		return
	}

	err = models.CheckCloudplatformData()
	if err != nil {
		logs.Error("CheckCloudplatformData failed, err:%s", err.Error())
		return
	}

	err = models.CheckSystemData()
	if err != nil {
		logs.Error("CheckSystemData failed, err:%s", err.Error())
		return
	}

	beego.Run()
}

func initLog() {
	//logs.SetLogger(logs.AdapterFile,`{"filename":"smartplug.log"}`)
	// 日志打印到控制台
	logs.SetLogger(logs.AdapterConsole)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

	switch config.LogLevel {
	case "LevelInfo":
		logs.SetLevel(logs.LevelInfo)
	case "LevelTrace":
		logs.SetLevel(logs.LevelTrace)
	case "LevelWarn":
		logs.SetLevel(logs.LevelWarn)
	default:
		logs.SetLevel(logs.LevelInfo)
	}
}

func initDB() {
	if config.DBDriver != "mysql" {
		logs.Error("unsupport db driver %s", config.DBDriver)
		os.Exit(1)
	}

	err := orm.RegisterDriver(config.DBDriver, orm.DRMySQL)
	if err != nil {
		logs.Error("RegisterDriver fail, ", err.Error())
		return
	}
	err = orm.RegisterDataBase("default", config.DBDriver,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBuser, config.DBpasswd,
			config.DBaddr, config.DBport, config.Database))
	if err != nil {
		logs.Error("RegisterDataBase fail, ", err.Error())
		return
	}
	orm.RegisterModel(new(models.Timer), new(models.Delay), new(models.Infrared), new(models.Meter),
		new(models.Cloudplatform), new(models.System))
	orm.RunSyncdb("default", false, true)

	logs.Info("init DataBase succes")
}

func initWeb() {
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.SetStaticPath("/", "static")
}
