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
	"reflect"
	"smartplug/models"

	"net/http"
)

type CloudplatformController struct {
	beego.Controller
}

func (c *CloudplatformController) GetCloudplatform() {
	cpf := models.Cloudplatform{ID:1}
	cloudPlatForm, err := cpf.All()
	if err != nil {
		msg := fmt.Sprintf("query all cloudplatforms DB failed. err:%s", err.Error())
		logs.Error(msg)
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, msg)))
	}
	data, err := json.Marshal(cloudPlatForm)
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
	cloudPlatFormMap := make(map[string]interface{})
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cloudPlatFormMap)
	if err != nil {
		msg := fmt.Sprintf("unmarshal failed. err:%s", err.Error())
		logs.Error(msg)
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"failed", "msg":"%s"}`, err.Error())))
		return
	}

	code, err := updateCloudPlatformData(cloudPlatFormMap)
	if err != nil {
		logs.Error(err.Error())
		c.Ctx.ResponseWriter.WriteHeader(code)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"failed", "msg":"%s"}`, err.Error())))
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(`{"result":"success", "msg":""}`))
}

func updateCloudPlatformData(cloudPlatFormMap map[string]interface{}) (int, error) {
	c := models.Cloudplatform{ID: 1}
	cpf, err := c.All()
	if err != nil {
		msg := fmt.Sprintf("query all cloudplatform failed, err.%s", err.Error())
		logs.Error(msg)
		return http.StatusInternalServerError, errors.New(msg)
	}

	for i := 0; i < reflect.TypeOf(cpf).Elem().NumField(); i++ {
		fieldName := reflect.TypeOf(cpf).Elem().Field(i).Name
		if v, ok := cloudPlatFormMap[fieldName]; ok {
			fieldValue := reflect.ValueOf(cpf).Elem().Field(i)
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
	err = cpf.Update()
	if err != nil {
		msg := fmt.Sprintf("update cloudplatform failed, err.%s", err.Error())
		logs.Error(msg)
		return http.StatusInternalServerError, errors.New(msg)
	}
	return http.StatusOK, nil
}
