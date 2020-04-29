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
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Println("RegisterDriver fail, ", err.Error())
		return
	}
	err = orm.RegisterDataBase("default", "mysql", "root:123456@tcp(192.168.1.107:3306)/smartplug")
	if err != nil {
		fmt.Println("RegisterDataBase fail, ", err.Error())
		return
	}
	orm.RegisterModel(new(Timer), new(Delay), new(Infrared), new(Meter), new(Cloudplatform))
	orm.RunSyncdb("default", false, true)
}
