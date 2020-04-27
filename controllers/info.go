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
	"github.com/astaxie/beego"
)

type InfoController struct {
	beego.Controller
}

func (c *InfoController) GetInfo() {
	c.Ctx.Output.Body([]byte(`{"GitCommit":"0be1d90165cf509ee102df72ddbce5cddf8c5c43","BuildDate":"Nov 29 2019 20:32:5","SDKVersion":"2.0.0(e271380)","FlashMap":"4M","UserBin":"user2.bin","RunTime":2279864}`))
}
