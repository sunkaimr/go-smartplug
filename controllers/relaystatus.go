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
	"github.com/astaxie/beego/logs"
	"net/http"
	"smartplug/models"
)

const (
	ON  = "on"
	OFF = "off"
)

type RelaystatusController struct {
	beego.Controller
}

type RelayStatus struct {
	Status string `json:"status"`
}

func (c *RelaystatusController) GetRelaystatus() {
	relayStatus, code, err := getRelayStatus()
	if err != nil {
		logs.Error("get relaystatus failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(code)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	body, err := json.Marshal(relayStatus)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.Output.Body(body)
}

func (c *RelaystatusController) SetRelaystatus() {
	relayStatus := &RelayStatus{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, relayStatus)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}
	code, err := updateRelayStatus(relayStatus)
	if err != nil {
		logs.Error("update relaystatus failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(code)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}

	body, err := json.Marshal(relayStatus)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.Output.Body(body)
}

func getRelayStatus() (*RelayStatus, int, error) {
	s := models.System{}
	system, err := s.All()
	if err != nil {
		logs.Error("query relayStatus data failed, err:%s", err.Error())
		return nil, http.StatusInternalServerError, err
	}
	relayStatus := &RelayStatus{}
	if system.RelayStatus {
		relayStatus.Status = ON
	} else {
		relayStatus.Status = OFF
	}

	return relayStatus, http.StatusOK, nil
}

func updateRelayStatus(relayStatus *RelayStatus) (int, error) {
	s := models.System{}
	system, err := s.All()
	if err != nil {
		logs.Error("query relayStatus data failed, err:%s", err.Error())
		return http.StatusInternalServerError, err
	}
	if relayStatus.Status == ON {
		system.RelayStatus = true
	} else {
		system.RelayStatus = false
	}

	err = system.Update()
	if err != nil {
		logs.Error("update relayStatus data failed, err:%s", err.Error())
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
