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

type InfraredController struct {
	beego.Controller
}

type Infrared struct {
	Num      int    `json:"Num"`
	Name     string `json:"Name"`
	Enable   bool   `json:"Enable"`
	OnValue  string `json:"OnValue"`
	OffValue string `json:"OffValue"`
}

var infrared = []Infrared{}

func init() {
	for i := 0; i < 10; i++ {
		v := Infrared{}
		v.Num = i + 1
		v.Name = fmt.Sprintf("infrared %d", i+1)
		v.Enable = false
		v.OnValue = "0"
		v.OffValue = "0"

		infrared = append(infrared, v)
	}
}

func (c *InfraredController) GetInfrared() {
	infraredStr := c.Ctx.Input.Param(":infrared")
	respInfrared := []Infrared{}
	if infraredStr == "all" {
		respInfrared = infrared
	} else {
		num, err := strconv.Atoi(infraredStr)
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
		respInfrared = append(respInfrared, infrared[num-1])
	}

	data, err := json.Marshal(respInfrared)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(
			fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write(data)
}

func (c *InfraredController) UpdateInfrared() {
	respInfrared := []Infrared{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &respInfrared)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(
			fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}

	for _, v := range respInfrared {
		if v.Num < 0 || v.Num >= len(delay) {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			c.Ctx.ResponseWriter.Write([]byte(
				fmt.Sprintf(`{"result":"fail", "msg":"%s is num is less 0 or large than %d"}`, v.Name, len(infrared))))
			return
		}
		infrared[v.Num-1] = v
	}

	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(`{"result":"success", "msg":""}`))
}

func (c *InfraredController) GetInfraredValue() {
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(`{"Num":1, "Value":"FF02FD"}`))
}
