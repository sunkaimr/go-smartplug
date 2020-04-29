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

type MeterController struct {
	beego.Controller
}

func (c *MeterController) GetMeter() {
	meter := &models.Meter{ID: 1}
	meters, err := meter.All()
	if err != nil {
		logs.Error("query all meter failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
	}

	data, err := json.Marshal((*meters)[0])
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write(data)
}

func (c *MeterController) SetMeter() {
	meter := make(map[string]interface{})
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &meter)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}

	code, err := updateMeterDB(meter)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(code)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(`{"result":"success", "msg":""}`))
}

func updateMeterDB(meter map[string]interface{}) (int, error) {
	m, err := (&models.Meter{ID: 1}).All()
	if err != nil {
		logs.Error("query meter failed, err:%s", err.Error())
		return http.StatusInternalServerError, err
	}

	for k, v := range meter {
		switch k {
		case models.Electricity:
			(*m)[0].Electricity = v.(string)
		case models.RunTime:
			(*m)[0].RunTime = v.(string)
		case models.OverCurrent:
			(*m)[0].OverCurrent = v.(string)
		case models.OverCurrentEnable:
			(*m)[0].OverCurrentEnable = v.(bool)
		case models.OverPower:
			(*m)[0].OverPower = v.(string)
		case models.OverPowerEnable:
			(*m)[0].OverPowerEnable = v.(bool)
		case models.OverVoltage:
			(*m)[0].OverVoltage = v.(string)
		case models.OverVoltageEnable:
			(*m)[0].OverVoltageEnable = v.(bool)
		case models.UnderPower:
			(*m)[0].UnderPower = v.(string)
		case models.UnderPowerEnable:
			(*m)[0].UnderPowerEnable = v.(bool)
		case models.UnderVoltage:
			(*m)[0].UnderVoltage = v.(string)
		case models.UnderVoltageEnable:
			(*m)[0].UnderVoltageEnable = v.(bool)
		}
	}
	err = (&(*m)[0]).Update()
	if err != nil {
		logs.Error("update meter to DB failed, err:%s", err.Error())
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
