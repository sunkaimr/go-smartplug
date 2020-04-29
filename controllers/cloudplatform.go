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
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"smartplug/models"

	"net/http"
)

type CloudplatformController struct {
	beego.Controller
}

const (
	CloudPlatform  = "CloudPlatform"
	MqttProductKey = "MqttProductKey"
	MqttDevName    = "MqttDevName"
	MqttDevSecret  = "MqttDevSecret"
	DevType        = "DevType"
	BigiotDevId    = "BigiotDevId"
	BigiotApiKey   = "BigiotApiKey"
	SwitchId       = "SwitchId"
	TempId         = "TempId"
	HumidityId     = "HumidityId"
	VoltageId      = "VoltageId"
	CurrentId      = "CurrentId"
	PowerId        = "PowerId"
	ElectricityId  = "ElectricityId"
	BigiotDevName  = "BigiotDevName"
	ConnectSta     = "ConnectSta"
)

func (c *CloudplatformController) GetCloudplatform() {
	cpf := models.Cloudplatform{ID:1}
	cloudplatforms, err := cpf.All()
	if err != nil {
		msg := fmt.Sprintf("query all cloudplatforms DB failed. err:%s", err.Error())
		logs.Error(msg)
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, msg)))
	}
	data, err := json.Marshal((*cloudplatforms)[0])
	if err != nil {
		logs.Error(err.Error())
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write(data)

}

func (c *CloudplatformController) SetCloudplatform() {
	cloudplatform := make(map[string]interface{})
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cloudplatform)
	if err != nil {
		msg := fmt.Sprintf("unmarshal failed. err:%s", err.Error())
		logs.Error(msg)
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"failed", "msg":"%s"}`, err.Error())))
		return
	}

	code, err := updateCloudPlatformData(cloudplatform)
	if err != nil {
		logs.Error(err.Error())
		c.Ctx.ResponseWriter.WriteHeader(code)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"failed", "msg":"%s"}`, err.Error())))
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(`{"result":"success", "msg":""}`))
}

func updateCloudPlatformData(cloudplatform map[string]interface{}) (int, error) {
	c := models.Cloudplatform{ID: 1}
	cpfs, err := c.All()
	if err != nil {
		msg := fmt.Sprintf("query all cloudplatform failed, err.%s", err.Error())
		logs.Error(msg)
		return http.StatusInternalServerError, errors.New(msg)
	}
	cpf := (*cpfs)[0]

	for k, v := range cloudplatform {
		switch k {
		case CloudPlatform:
			cpf.CloudPlatform = v.(uint8)
		case MqttProductKey:
			cpf.MqttProductKey = v.(string)
		case MqttDevName:
			cpf.MqttDevName = v.(string)
		case MqttDevSecret:
			cpf.MqttDevSecret = v.(string)
		case DevType:
			cpf.DevType = v.(uint8)
		case BigiotDevId:
			cpf.BigiotDevId = v.(string)
		case BigiotApiKey:
			cpf.BigiotApiKey = v.(string)
		case SwitchId:
			cpf.SwitchId = v.(string)
		case TempId:
			cpf.TempId = v.(string)
		case HumidityId:
			cpf.HumidityId = v.(string)
		case VoltageId:
			cpf.VoltageId = v.(string)
		case CurrentId:
			cpf.CurrentId = v.(string)
		case PowerId:
			cpf.PowerId = v.(string)
		case ElectricityId:
			cpf.ElectricityId = v.(string)
		case BigiotDevName:
			cpf.BigiotDevName = v.(string)
		case ConnectSta:
			cpf.ConnectSta = v.(string)
		}
	}
	err = cpf.Update()
	if err != nil {
		msg := fmt.Sprintf("update cloudplatform failed, err.%s", err.Error())
		logs.Error(msg)
		return http.StatusInternalServerError, errors.New(msg)
	}
	return http.StatusOK, nil
}
