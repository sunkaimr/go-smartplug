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
	"SmartPlug/models"
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

type timerResp struct {
	Num        int    `json:"Num"`
	Name       string `json:"Name"`
	Enable     bool   `json:"Enable"`
	OnEnable   bool   `json:"OnEnable"`
	OffEnable  bool   `json:"OffEnable"`
	Cascode    bool   `json:"Cascode"`
	Week       int    `json:"Week"`
	CascodeNum int    `json:"CascodeNum"`
	OnTime     string `json:"OnTime"`
	OffTime    string `json:"OffTime"`
}

func (c *TimerController) GetTimer() {
	timerStr := c.Ctx.Input.Param(":timer")
	respTimer := &[]timerResp{}
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

	respTimer, code, err := queryTimer(num)
	if err != nil {
		logs.Error("queryTimer failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(code)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}

	data, err := json.Marshal(*respTimer)
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
	respTimer := &[]timerResp{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, respTimer)
	if err != nil {
		logs.Error("Unmarshal failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"fail", "msg":"%s"}`, err.Error())))
		return
	}

	code, err := updateTimerDB(respTimer)
	if err != nil {
		logs.Error("update timerDB failed, err:%s", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(code)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"result":"failed", "msg":"%s"}`, err.Error())))
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Write([]byte(`{"result":"success", "msg":""}`))
}

func assemblyTimerResp(db *[]models.Timerdb) *[]timerResp {
	timer := &[]timerResp{}
	for _, t := range *db {
		tr := timerResp{}
		tr.Num = t.Num
		tr.Name = t.Name
		tr.Enable = t.Enable
		tr.OnEnable = t.OnEnable
		tr.OffEnable = t.OffEnable
		tr.Cascode = t.Cascode
		tr.Week = t.Week
		tr.CascodeNum = t.CascodeNum
		tr.OnTime = t.OnTime
		tr.OffTime = t.OffTime

		*timer = append(*timer, tr)
	}
	return timer
}

func queryTimer(timerNum int) (*[]timerResp, int, error) {
	t := &models.Timerdb{}
	if timerNum == 0 {

		timers, err := t.All()
		if err != nil {
			logs.Error("query all timer failed, err:%s", err.Error())
			return nil, http.StatusInternalServerError, err
		}
		return assemblyTimerResp(timers), http.StatusOK, nil
	}

	t.Num = timerNum
	timers, err := t.GetByID()
	if err != nil {
		logs.Error("query timer failed timer.Num=%d, err:%s", t.Num, err.Error())
		return nil, http.StatusInternalServerError, err
	}
	return assemblyTimerResp(timers), http.StatusOK, nil
}

func updateTimerDB(timerRes *[]timerResp) (int, error) {
	for _, t := range *timerRes {
		db := models.Timerdb{}
		db.Num = t.Num
		db.Name = t.Name
		db.Enable = t.Enable
		db.OnEnable = t.OnEnable
		db.OffEnable = t.OffEnable
		db.Cascode = t.Cascode
		db.Week = t.Week
		db.CascodeNum = t.CascodeNum
		db.OnTime = t.OnTime
		db.OffTime = t.OffTime

		err := db.UpdateByID()
		if err != nil {
			logs.Error("update timer filed timer.Num=%d, err:%s", db.Num, err.Error())
			return http.StatusInternalServerError, err
		}
	}
	return http.StatusOK, nil
}
