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

const ()

type System struct {
	ID              int       `orm:"column(id);pk"`
	PlugName        string    `json:"PlugName" orm:"column(plug_name);"`
	RelayPowerUp    int       `json:"RelayPowerUp"  orm:"column(relay_power_up);"`
	WifiMode        int       `json:"WifiMode"  orm:"column(wifi_mode);"`
	WifiSSID        string    `json:"WifiSSID"  orm:"column(wifi_ssid);"`
	WifiPasswd      string    `json:"WifiPasswd"  orm:"column(wifi_passwd);"`
	SmartConfigFlag bool      `json:"SmartConfigFlag"  orm:"column(smart_config_flag);"`
	RelayStatus     bool      `json:"RelayStatus"  orm:"column(relay_status);"`
	IP              string    `json:"IP"  orm:"column(ip);"`
	GetWay          string    `json:"GetWay"  orm:"column(getway);"`
	NetMask         string    `json:"NetMask"  orm:"column(netmask);"`
	Mac             string    `json:"Mac"  orm:"column(mac);"`
	CreateAt        time.Time `json:"CreateAt" orm:"column(create_at);auto_now"`
	UpdateAt        time.Time `json:"UpdateAt" orm:"column(update_at);auto_now_add"`
}

func (m *System) TableName() string {
	return "system"
}

func (m *System) All() (*System, error) {
	d := &System{}
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

func (m *System) Add() error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	if err != nil {
		return err
	}
	return nil
}

func (m *System) Update() error {
	d := &System{}
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

func (m *System) Exist() bool {
	o := orm.NewOrm()
	return o.QueryTable(m).Exist()
}

func newDefaultSystem() *System {
	return &System{
		PlugName:        "smartplug",
		RelayPowerUp:    0,
		WifiMode:        1,
		WifiSSID:        "TPLINK",
		WifiPasswd:      "123456",
		SmartConfigFlag: true,
		RelayStatus:     true,
		IP:              "192.168.1.104",
		GetWay:          "192.168.1.1",
		NetMask:         "255.255.255.0",
		Mac:             "ECFABC0D6308",
	}
}

func CheckSystemData() error {
	t := &System{ID:1}
	if !t.Exist() {
		return newDefaultSystem().Add()
	}
	return nil
}
