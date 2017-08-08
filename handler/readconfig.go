/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co.,Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
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
 * Revision History:
 *     Initial: 2017/08/06        Liu JiaChang
 */

package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"aliyun-oss-storage/bolt"
	"aliyun-oss-storage/log"
)

var (
	projectList *map[string]string
	readConf    func() error
)

func InitReadConf(projlist *map[string]string, f func() error) {
	projectList = projlist
	readConf = f
}

func UpdateConf(c echo.Context) error {
	var err error

	if err = readConf(); err != nil {
		log.Logger.Error("Read config error:", err)

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Logger.Debug("Project list....:", *projectList)

	if err = bolt.UseProjectList(*projectList); err != nil {
		log.Logger.Error("Using project list error:", err)

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func ReadBoltDb(c echo.Context) error {
	projList := bolt.GetProjectList()

	return c.JSON(http.StatusOK, projList)
}
