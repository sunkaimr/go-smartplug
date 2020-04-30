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

type Cloudplatform struct {
	ID             int       `orm:"column(id);pk"`
	CloudPlatform  int       `json:"CloudPlatform" orm:"column(cloud_platform);"`
	MqttProductKey string    `json:"MqttProductKey" orm:"column(mqtt_product_key)"`
	MqttDevName    string    `json:"cloud_platform" orm:"column(mqtt_devName)"`
	MqttDevSecret  string    `json:"MqttDevName" orm:"column(mqtt_dev_secret)"`
	DevType        int       `json:"DevType" orm:"column(dev_type);"`
	BigiotDevId    string    `json:"BigiotDevId" orm:"column(bigiot_dev_id)"`
	BigiotApiKey   string    `json:"BigiotApiKey" orm:"column(bigiot_api_key)"`
	SwitchId       string    `json:"SwitchId" orm:"column(switch_id)"`
	TempId         string    `json:"TempId" orm:"column(temp_id)"`
	HumidityId     string    `json:"HumidityId" orm:"column(humidity_id)"`
	VoltageId      string    `json:"VoltageId" orm:"column(voltage_id)"`
	CurrentId      string    `json:"CurrentId" orm:"column(current_id)"`
	PowerId        string    `json:"PowerId" orm:"column(power_id)"`
	ElectricityId  string    `json:"ElectricityId" orm:"column(electricity_id)"`
	BigiotDevName  string    `json:"BigiotDevName" orm:"column(bigiot_dev_name)"`
	ConnectSta     string    `json:"ConnectSta" orm:"column(connect_sta)"`
	CreateAt       time.Time `json:"CreateAt" orm:"column(create_at);auto_now"`
	UpdateAt       time.Time `json:"UpdateAt" orm:"column(update_at);auto_now_add"`
}

func (m *Cloudplatform) TableName() string {
	return "cloudplatform"
}

func (m *Cloudplatform) All() (*Cloudplatform, error) {
	d := &Cloudplatform{}
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

func (m *Cloudplatform) Add() error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	if err != nil {
		return err
	}
	return nil
}

func (m *Cloudplatform) Update() error {
	d := &[]Cloudplatform{}
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

func (m *Cloudplatform) Exist() bool {
	o := orm.NewOrm()
	return o.QueryTable(m).Exist()
}

func newDefaultCloudplatform() *Cloudplatform {
	return &Cloudplatform{
		ID:             1,
		CloudPlatform:  2,
		MqttProductKey: "",
		MqttDevName:    "",
		MqttDevSecret:  "",
		DevType:        0,
		BigiotDevId:    "12429",
		BigiotApiKey:   "bf7ab3c13",
		SwitchId:       "11380",
		TempId:         "11772",
		HumidityId:     "",
		VoltageId:      "13981",
		CurrentId:      "13982",
		PowerId:        "13983",
		ElectricityId:  "13984",
		BigiotDevName:  "电热水壶",
		ConnectSta:     "connect",
	}
}

func CheckCloudplatformData() error {
	t := &Cloudplatform{ID: 1}
	if !t.Exist() {
		return newDefaultCloudplatform().Add()
	}
	return nil
}
