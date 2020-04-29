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

type SystemController struct {
	beego.Controller
}

type SystemSet struct {
	PlugName        string `json:"PlugName"`
	RelayPowerUp    int    `json:"RelayPowerUp"`
	WifiMode        int    `json:"WifiMode"`
	WifiSSID        string `json:"WifiSSID"`
	WifiPasswd      string `json:"WifiPasswd"`
	SmartConfigFlag bool   `json:"SmartConfigFlag"`
	RelayStatus     bool   `json:"RelayStatus"`
	IP              string `json:"IP"`
	GetWay          string `json:"GetWay"`
	NetMask         string `json:"NetMask"`
	Mac             string `json:"Mac"`
}

var systemSet = SystemSet{}

func init() {
	systemSet.PlugName = "smartplug"
	systemSet.RelayPowerUp = 0
	systemSet.WifiMode = 1
	systemSet.WifiSSID = "TPLINK"
	systemSet.WifiPasswd = "123456"
	systemSet.SmartConfigFlag = true
	systemSet.RelayStatus = true
	systemSet.IP = "192.168.1.104"
	systemSet.GetWay = "192.168.1.1"
	systemSet.NetMask = "255.255.255.0"
	systemSet.Mac = "ECFABC0D6308"
}

func (c *SystemController) GetSystem() {
	data, err := json.Marshal(systemSet)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(
			fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write(data)

}

func (c *SystemController) SetSystem() {
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &systemSet)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(
			fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(`{"result":"success", "msg":""}`))
}
