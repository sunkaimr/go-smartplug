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

type MeterController struct {
	beego.Controller
}

type Meter struct {
	Voltage            string `json:"Voltage"`
	Current            string `json:"Current"`
	Power              string `json:"Power"`
	ApparentPower      string `json:"ApparentPower"`
	PowerFactor        string `json:"PowerFactor"`
	Electricity        string `json:"Electricity"`
	RunTime            string `json:"RunTime"`
	UnderVoltage       string `json:"UnderVoltage"`
	OverVoltage        string `json:"OverVoltage"`
	OverCurrent        string `json:"OverCurrent"`
	OverPower          string `json:"OverPower"`
	UnderPower         string `json:"UnderPower"`
	UnderVoltageEnable bool   `json:"UnderVoltageEnable"`
	OverVoltageEnable  bool   `json:"OverVoltageEnable"`
	OverCurrentEnable  bool   `json:"OverCurrentEnable"`
	OverPowerEnable    bool   `json:"OverPowerEnable"`
	UnderPowerEnable   bool   `json:"UnderPowerEnable"`
}

var meter = Meter{}

func init() {
	meter.Voltage = "226.6"
	meter.Current = "0.0"
	meter.Power = "0.0"
	meter.ApparentPower = "0.0"
	meter.PowerFactor = "0.00"
	meter.Electricity = "9887.6"
	meter.RunTime = "3480.9"
	meter.UnderVoltage = "180"
	meter.OverVoltage = "250"
	meter.OverCurrent = " 10"
	meter.OverPower = "2200"
	meter.UnderPower = "0.5"
	meter.UnderVoltageEnable = true
	meter.OverVoltageEnable = true
	meter.OverCurrentEnable = true
	meter.OverPowerEnable = true
	meter.UnderPowerEnable = false
}

func (c *MeterController) GetMeter() {
	data, err := json.Marshal(meter)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(
			fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write(data)

}

func (c *MeterController) SetMeter() {
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &meter)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(
			fmt.Sprintf(`{"result":"success", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(`{"result":"success", "msg":""}`))
}
