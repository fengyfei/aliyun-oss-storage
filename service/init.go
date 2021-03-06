/*
* MIT License
*
* Copyright (c) 2017 SmartestEE Co.,Ltd..
*
* Permission is hereby granted, free of charge, to any person obtaining a copy of
* this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:
*
* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.
*
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
 */

/*
* Revision History
*     Initial: 2017-08-01          Yusan Kurban
 */

package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/fengyfei/aliyun-oss-storage/ali"
	"github.com/fengyfei/aliyun-oss-storage/general"
	"github.com/fengyfei/aliyun-oss-storage/router"
	"github.com/fengyfei/aliyun-oss-storage/bolt"
	"github.com/fengyfei/aliyun-oss-storage/handler"
)

func startService() {
	server := echo.New()

	server.Use(middleware.Recover())

	server.HTTPErrorHandler = general.EchoRestfulErrorHandler
	server.Validator = general.NewEchoValidator()

	router.InitRouter(server)

	server.Start(conf.Address)
}

func init() {
	readConfiguration()

	bolt.InitBolt(conf.BoltPath)
	bolt.UseProjectList(projectList)

	handler.InitReadConf(&projectList, readProjectList)

	ali.Connection(conf.EndPoint, conf.AccessKeyID, conf.AccessKeySecret)
	ali.GetBucket(conf.BucketName)
}
