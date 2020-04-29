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
package routers

import (
	"smartplug/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/timer", &controllers.TimerController{}, "post:UpdateTimer")
	beego.Router("/timer/:timer", &controllers.TimerController{}, "get:GetTimer")
	beego.Router("/delay", &controllers.DelayController{}, "post:UpdateDelay")
	beego.Router("/delay/:delay", &controllers.DelayController{}, "get:GetDelay")
	beego.Router("/infrared", &controllers.InfraredController{}, "post:UpdateInfrared")
	beego.Router("/infrared/:infrared", &controllers.InfraredController{}, "get:GetInfrared")
	beego.Router("/infrared/:infrared/switch/:switch", &controllers.InfraredController{}, "get:GetInfraredValue")

	beego.Router("/webset", &controllers.WebSetController{}, "get:GetWebSet;post:UpdateWebSet")
	beego.Router("/relaystatus", &controllers.RelaystatusController{}, "get:GetRelaystatus;post:SetRelaystatus")
	beego.Router("/meter", &controllers.MeterController{}, "get:GetMeter;post:SetMeter")
	beego.Router("/info", &controllers.InfoController{}, "get:GetInfo")
	beego.Router("/date", &controllers.DateController{}, "get:GetDate;post:SetDate")
	beego.Router("/temperature", &controllers.TemperatureController{}, "get:GetTemperature")
	beego.Router("/cloudplatform", &controllers.CloudplatformController{}, "get:GetCloudplatform;post:SetCloudplatform")
	beego.Router("/system", &controllers.SystemController{}, "get:GetSystem;post:SetSystem")
	beego.Router("/control", &controllers.DeviceController{}, "post:DeviceControl")
}
