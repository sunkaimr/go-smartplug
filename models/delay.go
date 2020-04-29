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
package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	MaxDelayNum = 10
)

type Delay struct {
	Num           int    `json:"Num" orm:"column(id);pk"`
	Name          string `json:"Name" orm:"column(name);"`
	Enable        bool   `json:"Enable" orm:"column(enable);"`
	OnEnable      bool   `json:"OnEnable" orm:"column(on_enable);"`
	OffEnable     bool   `json:"OffEnable" orm:"column(off_enable);"`
	CycleTimes    int    `json:"CycleTimes" orm:"column(cycle);"`
	TmpCycleTimes int    `json:"TmpCycleTimes" orm:"column(tmp_cycle);"`
	SwFlag        int    `json:"SwFlag" orm:"column(sw_flag);"`
	Cascode       bool   `json:"Cascode" orm:"column(cascode);"`
	CascodeNum    int    `json:"CascodeNum" orm:"column(cascode_num);"`
	OnInterval    string `json:"OnInterval" orm:"column(on_interval);"`
	OffInterval   string `json:"OffInterval" orm:"column(off_interval);"`
	TimePoint     string `json:"TimePoint" orm:"column(time_point);"`
	CreateAt   time.Time `json:"CreateAt" orm:"column(create_at);auto_now_add"`
	UpdateAt   time.Time `json:"UpdateAt" orm:"column(update_at);auto_now"`

}

func (t *Delay) TableName() string {
	return "delay"
}

func (t *Delay) Add() error {
	o := orm.NewOrm()
	_, err := o.Insert(t)
	if err != nil {
		return err
	}
	return nil
}

func (t *Delay) Delete() error {
	o := orm.NewOrm()
	_, err := o.QueryTable(t).Filter("id", t.Num).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (t *Delay) Exist() bool {
	o := orm.NewOrm()
	return o.QueryTable(t).Filter("id", t.Num).Exist()
}

func (t *Delay) All() (*[]Delay, error) {
	d := []Delay{}
	o := orm.NewOrm()
	_, err := o.QueryTable(t).All(&d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (t *Delay) GetByID() (*[]Delay, error) {
	d := &[]Delay{}
	o := orm.NewOrm()
	_, err := o.QueryTable(t).Filter("id", t.Num).All(d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (t *Delay) UpdateByID() error {
	d := &[]Delay{}
	o := orm.NewOrm()
	count, err := o.QueryTable(t).Filter("id", t.Num).All(d)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("unable to query record")
	}
	if count != 1 {
		return errors.New("multiple data")
	}
	_, err = o.Update(t)
	if err != nil {
		return errors.New(fmt.Sprintf("update data failed, err:%s", err.Error()))
	}
	return nil
}

//CheckDelayData 检查数据是否完整
func CheckDelayData() error {
	t := &Delay{}
	for i := 1; i <= MaxDelayNum; i++ {
		t.Num = i
		exist := t.Exist()
		if !exist {
			d := newDelay(i)
			err := d.Add()
			if err != nil {
				logs.Error("insert delay data into DB failed, err:%s", err.Error())
				return err
			}
			logs.Info("insert delay data into DB suncces, delay.Num=%d", d.Num)
		}
	}

	delays, err := t.All()
	if err != nil {
		logs.Info("query all delay data from DB suncces")
		return err
	}
	for _, timer := range *delays {
		if timer.Num < 1 || timer.Num > MaxDelayNum {
			err := timer.Delete()
			if err != nil {
				logs.Error("delete delay data from DB failed, timer.Num=%d err: %s", timer.Num, err.Error())
				return err
			}
			logs.Info("delete delay data from DB, timer.Num=%d", timer.Num)
		}
	}
	return nil
}

func newDelay(num int) *Delay {
	return &Delay{
		Num : num,
		Name : fmt.Sprintf("delay %d",num),
		Enable : false,
		OnEnable : true,
		OffEnable : true,
		CycleTimes : 1,
		TmpCycleTimes : 1,
		SwFlag : 0,
		Cascode : true,
		CascodeNum : num + 1,
		OnInterval : "00:00",
		OffInterval : "00:00",
		TimePoint : "00:00",
	}
}
