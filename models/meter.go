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
	"github.com/astaxie/beego/orm"
	"time"
)

type Meter struct {
	ID                 int       `orm:"column(id);pk"`
	Voltage            string    `json:"Voltage" orm:"column(voltage);"`
	Current            string    `json:"Current" orm:"column(current);"`
	Power              string    `json:"Power" orm:"column(power);"`
	ApparentPower      string    `json:"ApparentPower" orm:"column(apparent_power);"`
	PowerFactor        string    `json:"PowerFactor" orm:"column(power_factor);"`
	Electricity        string    `json:"Electricity" orm:"column(electricity);"`
	RunTime            string    `json:"RunTime" orm:"column(runTime);"`
	UnderVoltage       string    `json:"UnderVoltage" orm:"column(under_voltage);"`
	OverVoltage        string    `json:"OverVoltage" orm:"column(over_voltage);"`
	OverCurrent        string    `json:"OverCurrent" orm:"column(over_current);"`
	OverPower          string    `json:"OverPower" orm:"column(over_power);"`
	UnderPower         string    `json:"UnderPower" orm:"column(under_power);"`
	UnderVoltageEnable bool      `json:"UnderVoltageEnable" orm:"column(under_voltage_enable);"`
	OverVoltageEnable  bool      `json:"OverVoltageEnable" orm:"column(over_voltage_enable);"`
	OverCurrentEnable  bool      `json:"OverCurrentEnable" orm:"column(over_current_enable);"`
	OverPowerEnable    bool      `json:"OverPowerEnable" orm:"column(over_power_enable);"`
	UnderPowerEnable   bool      `json:"UnderPowerEnable" orm:"column(under_power_enable);"`
	CreateAt           time.Time `json:"CreateAt" orm:"column(create_at);auto_now_add"`
	UpdateAt           time.Time `json:"UpdateAt" orm:"column(update_at);auto_now"`
}

func (m *Meter) TableName() string {
	return "meter"
}

func (m *Meter) All() (*Meter, error) {
	d := &Meter{}
	o := orm.NewOrm()
	n, err := o.QueryTable(m).Filter("id", m.ID).All(d)
	if err != nil {
		return nil, err
	}
	if n < 1 {
		return nil, nil
	}
	return d, nil
}

func (m *Meter) Add() error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	if err != nil {
		return err
	}
	return nil
}

func (m *Meter) Update() error {
	d := &[]Meter{}
	o := orm.NewOrm()
	count, err := o.QueryTable(m).Filter("id", m.ID).All(d)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("unable to query record")
	}
	if count != 1 {
		return errors.New("multiple data")
	}
	_, err = o.Update(m)
	if err != nil {
		return errors.New(fmt.Sprintf("update data failed, err:%s", err.Error()))
	}
	return nil
}

func (m *Meter) Exist() bool {
	o := orm.NewOrm()
	return o.QueryTable(m).Exist()
}

func newDefaultMeter() *Meter {
	return &Meter{
		ID:                 1,
		Voltage:            "226.6",
		Current:            "0.1",
		Power:              "0.0",
		ApparentPower:      "0.0",
		PowerFactor:        "0.00",
		Electricity:        "9887.6",
		RunTime:            "3480.9",
		UnderVoltage:       "180",
		OverVoltage:        "250",
		OverCurrent:        "10",
		OverPower:          "2200",
		UnderPower:         "0.5",
		UnderVoltageEnable: true,
		OverVoltageEnable:  true,
		OverCurrentEnable:  true,
		OverPowerEnable:    true,
		UnderPowerEnable:   false,
	}
}

func CheckMeterData() error {
	t := &Meter{}
	if !t.Exist() {
		return newDefaultMeter().Add()
	}
	return nil
}
