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
	"smartplug/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/logs"
	"net/http"
	"strconv"
)

type TimerController struct {
	beego.Controller
}

func (c *TimerController) GetTimer() {
	timerStr := c.Ctx.Input.Param(":timer")
	timers := &[]models.Timer{}
	num := 0
	if timerStr == "all" {
		num = 0
	} else {
		num, err := strconv.Atoi(timerStr)
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

	timers, code, err := queryTimer(num)
	if err != nil {
		logs.Error("queryTimer failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(code)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}

	data, err := json.Marshal(*timers)
	if err != nil {
		logs.Error("Marshal failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write(data)
}

func (c *TimerController) UpdateTimer() {
	timers := &[]models.Timer{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, timers)
	if err != nil {
		logs.Error("Unmarshal failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}

	code, err := updateTimerDB(timers)
	if err != nil {
		logs.Error("update timerDB failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(code)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"failed", "msg":"%s"}`, err.Error())))
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(`{"result":"success", "msg":""}`))
}

func queryTimer(timerNum int) (*[]models.Timer, int, error) {
	t := &models.Timer{}
	if timerNum == 0 {
		timers, err := t.All()
		if err != nil {
			logs.Error("query all timer failed, err:%s", err.Error())
			return nil, http.StatusInternalServerError, err
		}
		return timers, http.StatusOK, nil
	}

	t.Num = timerNum
	timers, err := t.GetByID()
	if err != nil {
		logs.Error("query timer failed timer.Num=%d, err:%s", t.Num, err.Error())
		return nil, http.StatusInternalServerError, err
	}
	return timers, http.StatusOK, nil
}

func updateTimerDB(timers *[]models.Timer) (int, error) {
	for _, t := range *timers {
		err := t.UpdateByID()
		if err != nil {
			logs.Error("update timer filed timer.Num=%d, err:%s", t.Num, err.Error())
			return http.StatusInternalServerError, err
		}
	}
	return http.StatusOK, nil
}
