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
)

const (
	MaxTimerNum = 10
)

type Timerdb struct {
	Num        int    `orm:"column(id);pk"`
	Name       string `orm:"column(name);"`
	Enable     bool   `orm:"column(enable);"`
	OnEnable   bool   `orm:"column(on_enable);"`
	OffEnable  bool   `orm:"column(off_enable);"`
	Cascode    bool   `orm:"column(cascode);"`
	Week       int    `orm:"column(week);"`
	CascodeNum int    `orm:"column(cascode_num);"`
	OnTime     string `orm:"column(on_time);"`
	OffTime    string `orm:"column(off_time);"`
}

func (t *Timerdb) TableName() string {
	return "timer"
}

func (t *Timerdb) Add() error {
	o := orm.NewOrm()
	_, err := o.Insert(t)
	if err != nil {
		return err
	}
	return nil
}

func (t *Timerdb) Delete() error {
	o := orm.NewOrm()
	_, err := o.QueryTable(t).Filter("id", t.Num).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (t *Timerdb) Exist() bool {
	o := orm.NewOrm()
	return o.QueryTable(t).Filter("id", t.Num).Exist()
}

func (t *Timerdb) All() (*[]Timerdb, error) {
	d := []Timerdb{}
	o := orm.NewOrm()
	_, err := o.QueryTable(t).All(&d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (t *Timerdb) GetByID() (*[]Timerdb, error) {
	d := &[]Timerdb{}
	o := orm.NewOrm()
	_, err := o.QueryTable(t).Filter("id",t.Num).All(d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (t *Timerdb) UpdateByID() error{
	d := &[]Timerdb{}
	o := orm.NewOrm()
	count, err := o.QueryTable(t).Filter("id",t.Num).All(d)
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

//CheckTimerData 检查timer数据是否完整
func CheckTimerData() error {
	t := &Timerdb{}
	for i := 1; i <= MaxTimerNum; i++ {
		t.Num = i
		exist := t.Exist()
		if !exist {
			d := newDefaultTimer(i)
			err := d.Add()
			if err != nil {
				logs.Error("insert timer data into DB failed, err:%s", err.Error())
				return err
			}
			logs.Info("insert timer data into DB suncces, timer.Num=%d", d.Num)
		}
	}

	timers, err := t.All()
	if err != nil {
		logs.Info("query all timer data from DB suncces")
		return err
	}
	for _, timer := range *timers {
		if timer.Num < 1 || timer.Num > MaxTimerNum {
			err := timer.Delete()
			if err != nil {
				logs.Error("delete timer data from DB failed, timer.Num=%d err: %s", timer.Num, err.Error())
				return err
			}
			logs.Info("delete timer data from DB, timer.Num=%d", timer.Num)
		}
	}

	return nil
}

func newDefaultTimer(num int) *Timerdb {
	return &Timerdb{
		Num:        num,
		Name:       fmt.Sprintf("timer %d", num),
		Enable:     false,
		OnEnable:   false,
		OffEnable:  false,
		Cascode:    false,
		Week:       0,
		CascodeNum: num + 1,
		OnTime:     "00:00",
		OffTime:    "00:00",
	}
}





