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
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"

	"net/http"
)

type CloudplatformController struct {
	beego.Controller
}

type Cloudplatform struct {
	CloudPlatform  int    `json:"CloudPlatform"`
	MqttProductKey string `json:"MqttProductKey"`
	MqttDevName    string `json:"cloud_platform"`
	MqttDevSecret  string `json:"MqttDevName"`
	DevType        int    `json:"DevType"`
	BigiotDevId    string `json:"BigiotDevId"`
	BigiotApiKey   string `json:"BigiotApiKey"`
	SwitchId       string `json:"SwitchId"`
	TempId         string `json:"TempId"`
	HumidityId     string `json:"HumidityId"`
	VoltageId      string `json:"VoltageId"`
	CurrentId      string `json:"CurrentId"`
	PowerId        string `json:"PowerId"`
	ElectricityId  string `json:"ElectricityId"`
	BigiotDevName  string `json:"BigiotDevName"`
	ConnectSta     string `json:"ConnectSta"`
}

var cloudplatform = Cloudplatform{}

func init() {
	cloudplatform.CloudPlatform = 2
	cloudplatform.MqttProductKey = ""
	cloudplatform.MqttDevName = ""
	cloudplatform.MqttDevSecret = ""
	cloudplatform.DevType = 0
	cloudplatform.BigiotDevId = "12429"
	cloudplatform.BigiotApiKey = "bf7ab3c13"
	cloudplatform.SwitchId = "11380"
	cloudplatform.TempId = "11772"
	cloudplatform.HumidityId = ""
	cloudplatform.VoltageId = "13981"
	cloudplatform.CurrentId = "13982"
	cloudplatform.PowerId = "13983"
	cloudplatform.ElectricityId = "13984"
	cloudplatform.BigiotDevName = "电热水壶"
	cloudplatform.ConnectSta = "connect"
}

func (c *CloudplatformController) GetCloudplatform() {
	data, err := json.Marshal(cloudplatform)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(
			fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write(data)

}

func (c *CloudplatformController) SetCloudplatform() {
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cloudplatform)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(
			fmt.Sprintf(`{"result":"success", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(`{"result":"success", "msg":""}`))
}
