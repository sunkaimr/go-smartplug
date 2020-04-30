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
	"net/http"
	"reflect"
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

	data, err := json.Marshal(meters)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write(data)
}

func (c *MeterController) SetMeter() {
	meterMap := make(map[string]interface{})
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &meterMap)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}

	code, err := updateMeterDB(meterMap)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(code)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(`{"result":"success", "msg":""}`))
}

func updateMeterDB(meterMap map[string]interface{}) (int, error) {
	m, err := (&models.Meter{ID: 1}).All()
	if err != nil {
		logs.Error("query meter failed, err:%s", err.Error())
		return http.StatusInternalServerError, err
	}

	for i := 0; i < reflect.TypeOf(m).Elem().NumField(); i++ {
		fieldName := reflect.TypeOf(m).Elem().Field(i).Name
		if v, ok := meterMap[fieldName]; ok {
			fieldValue := reflect.ValueOf(m).Elem().Field(i)
			switch v.(type) {
			case string, bool:
				fieldValue.Set(reflect.ValueOf(v))
			default:
				msg := fmt.Sprintf("can not conversion %s:%s to %s",
					fieldName, fieldValue.Kind().String(), reflect.TypeOf(v).String())
				logs.Error(msg)
				return http.StatusInternalServerError, errors.New(msg)
			}
		}
	}

	err = m.Update()
	if err != nil {
		logs.Error("update meter to DB failed, err:%s", err.Error())
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
