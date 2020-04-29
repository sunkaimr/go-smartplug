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

type DelayController struct {
	beego.Controller
}

func (c *DelayController) GetDelay() {
	delayStr := c.Ctx.Input.Param(":delay")
	delays := &[]models.Delay{}
	num := 0
	if delayStr == "all" {
		num = 0
	} else {
		num, err := strconv.Atoi(delayStr)
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

	delays, code, err := queryDelay(num)
	if err != nil {
		logs.Error("queryDelay failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(code)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}

	data, err := json.Marshal(*delays)
	if err != nil {
		logs.Error("Marshal failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write(data)
}

func (c *DelayController) UpdateDelay() {
	delays := &[]models.Delay{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, delays)
	if err != nil {
		logs.Error("Unmarshal failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}

	code, err := updateDelayDB(delays)
	if err != nil {
		logs.Error("update timerDB failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(code)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"failed", "msg":"%s"}`, err.Error())))
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(`{"result":"success", "msg":""}`))
}

func queryDelay(num int) (*[]models.Delay, int, error) {
	d := &models.Delay{}
	if num == 0 {
		delays, err := d.All()
		if err != nil {
			logs.Error("query all delay failed, err:%s", err.Error())
			return nil, http.StatusInternalServerError, err
		}
		return delays, http.StatusOK, nil
	}

	d.Num = num
	delays, err := d.GetByID()
	if err != nil {
		logs.Error("query delay failed delay.Num=%d, err:%s", d.Num, err.Error())
		return nil, http.StatusInternalServerError, err
	}
	return delays, http.StatusOK, nil
}

func updateDelayDB(delays *[]models.Delay) (int, error) {
	for _, d := range *delays {
		err := d.UpdateByID()
		if err != nil {
			logs.Error("update delay filed delay.Num=%d, err:%s", d.Num, err.Error())
			return http.StatusInternalServerError, err
		}
	}
	return http.StatusOK, nil
}
