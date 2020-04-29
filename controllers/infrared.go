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
	"strconv"
)

type InfraredController struct {
	beego.Controller
}


func (c *InfraredController) GetInfrared() {
	infraredStr := c.Ctx.Input.Param(":infrared")
	infrared := &[]models.Infrared{}
	num := 0
	if infraredStr == "all" {
		num = 0
	} else {
		num, err := strconv.Atoi(infraredStr)
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
			return
		} else if num > models.MaxTimerNum {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%d up to limit %d"}`, num, models.MaxTimerNum)))
			return
		}
	}

	infrared, code, err := queryInfrared(num)
	if err != nil {
		logs.Error("queryTimer failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(code)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}

	data, err := json.Marshal(*infrared)
	if err != nil {
		logs.Error("Marshal failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write(data)
}

func (c *InfraredController) UpdateInfrared() {
	infrareds := &[]models.Infrared{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, infrareds)
	if err != nil {
		logs.Error("Unmarshal failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}

	code, err := updateInfraredDB(infrareds)
	if err != nil {
		logs.Error("update Infrared to DB failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(code)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"failed", "msg":"%s"}`, err.Error())))
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(`{"result":"success", "msg":""}`))
}

func queryInfrared(num int) (*[]models.Infrared, int, error) {
	infrared := &models.Infrared{}
	if num == 0 {
		timers, err := infrared.All()
		if err != nil {
			logs.Error("query all infrared failed, err:%s", err.Error())
			return nil, http.StatusInternalServerError, err
		}
		return timers, http.StatusOK, nil
	}

	infrared.Num = num
	infrareds, err := infrared.GetByID()
	if err != nil {
		logs.Error("query infrared failed timer.Num=%d, err:%s", infrared.Num, err.Error())
		return nil, http.StatusInternalServerError, err
	}
	return infrareds, http.StatusOK, nil
}

func updateInfraredDB(infrareds *[]models.Infrared) (int, error) {
	for _, i := range *infrareds {
		err := i.UpdateByID()
		if err != nil {
			logs.Error("update timer infrared infrareds.Num=%d, err:%s", i.Num, err.Error())
			return http.StatusInternalServerError, err
		}
	}
	return http.StatusOK, nil
}

func (c *InfraredController) GetInfraredValue() {
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(`{"Num":1, "Value":"FF02FD"}`))
}
