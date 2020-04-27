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
	"strconv"
)

type DelayController struct {
	beego.Controller
}

type Delay struct {
	Num           int    `json:"Num"`
	Name          string `json:"Name"`
	Enable        bool   `json:"Enable"`
	OnEnable      bool   `json:"OnEnable"`
	OffEnable     bool   `json:"OffEnable"`
	CycleTimes    int    `json:"CycleTimes"`
	TmpCycleTimes int    `json:"TmpCycleTimes"`
	SwFlag        int    `json:"SwFlag"`
	Cascode       bool   `json:"Cascode"`
	CascodeNum    int    `json:"CascodeNum"`
	OnInterval    string `json:"OnInterval"`
	OffInterval   string `json:"OffInterval"`
	TimePoint     string `json:"TimePoint"`
}

var delay = []Delay{}

func init() {
	for i := 0; i < 10; i++ {
		d := Delay{}
		d.Num = i + 1
		d.Name = fmt.Sprintf("delay %d", i+1)
		d.Enable = false
		d.OnEnable = true
		d.OffEnable = true
		d.CycleTimes = 1
		d.TmpCycleTimes = 1
		d.SwFlag = 0
		d.Cascode = true
		d.CascodeNum = i + 1
		d.OnInterval = "00:15"
		d.OffInterval = "00:01"
		d.TimePoint = "19:03"

		delay = append(delay, d)
	}
}

func (c *DelayController) GetDelay() {
	delayStr := c.Ctx.Input.Param(":delay")
	respDelay := []Delay{}
	if delayStr == "all" {
		respDelay = delay
	} else {
		num, err := strconv.Atoi(delayStr)
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			c.Ctx.ResponseWriter.Write([]byte(
				fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
			return
		} else if num > len(infrared) {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			c.Ctx.ResponseWriter.Write([]byte(
				fmt.Sprintf(`{"result":"fail", "msg":"%d up to limit %d"}`, num, len(infrared)-1)))
			return
		}
		respDelay = append(respDelay, delay[num-1])
	}

	data, err := json.Marshal(respDelay)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(
			fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write(data)
}

func (c *DelayController) UpdateDelay() {
	respDelay := []Delay{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &respDelay)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(
			fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}

	for _, v := range respDelay {
		if v.Num < 0 || v.Num >= len(delay) {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			c.Ctx.ResponseWriter.Write([]byte(
				fmt.Sprintf(`{"result":"fail", "msg":"%s is num is less 0 or large than %d"}`, v.Name, len(infrared))))
			return
		}
		delay[v.Num-1] = v
	}

	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(`{"result":"success", "msg":""}`))
}
