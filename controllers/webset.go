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
	"github.com/astaxie/beego"
	"net/http"
)

type WebSetController struct {
	beego.Controller
}

var webSet map[string]string

func init() {
	webSet = make(map[string]string)
	webSet["MeterRefresh"] = "0"
	webSet["ModelTab"] = "timer"
}

func (c *WebSetController) GetWebSet() {
	data, err := json.Marshal(webSet)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write(data)
}

func (c *WebSetController) UpdateWebSet() {
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &webSet)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	data, err := json.Marshal(webSet)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write(data)
}
