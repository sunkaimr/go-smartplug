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

type SystemController struct {
	beego.Controller
}

func (c *SystemController) GetSystem() {
	s := models.System{}
	system, err := s.All()
	if err != nil {
		logs.Error("query system data failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	data, err := json.Marshal(system)
	if err != nil {
		logs.Error("marshal failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write(data)
}

func (c *SystemController) SetSystem() {
	systemMap := make(map[string]interface{})
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &systemMap)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}

	code , err := updateSystemDB(systemMap)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(code)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(`{"result":"success", "msg":""}`))
}

func updateSystemDB( systemMap map[string]interface{})(int, error){
	s := models.System{}
	system, err := s.All()
	if err != nil {
		logs.Error("query system data failed, err:%s", err.Error())
		return http.StatusInternalServerError, err
	}

	for i := 0; i < reflect.TypeOf(system).Elem().NumField(); i++ {
		fieldName := reflect.TypeOf(system).Elem().Field(i).Name
		if v, ok := systemMap[fieldName]; ok {
			fieldValue := reflect.ValueOf(system).Elem().Field(i)
			switch v.(type){
			case float64:
				v1 := (int)(v.(float64))
				fieldValue.Set(reflect.ValueOf(v1))
			case string, bool :
				fieldValue.Set(reflect.ValueOf(v))
			default:
				msg := fmt.Sprintf("can not conversion %s:%s to %s",
					fieldName, fieldValue.Kind().String(), reflect.TypeOf(v).String())
				logs.Error(msg)
				return http.StatusInternalServerError, errors.New(msg)
			}
		}
	}
	err = system.Update()
	if err != nil {
		logs.Error("update system data failed, err:%s", err.Error())
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
