/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co., Ltd..
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
 *     Initial: 2017/08/05        Yang Chenglong
 */

package handler

import (
	"net/http"
	"strconv"
	"time"
	"strings"

	"github.com/labstack/echo"

	"aliyun-oss-storage/log"
	"aliyun-oss-storage/model"
)

func UploadFileHandler(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		log.Logger.Error("UploadFileHandler FormFile %v:", err)
		return err
	}

	suffixs := strings.Split(file.Filename, ".")
	suffix := suffixs[len(suffixs)-1]

	project := c.FormValue("project")
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	fileName := project + "/" + timestamp + "." + suffix

	src, err := file.Open()
	if err != nil {
		log.Logger.Error("UploadFileHandler Open %v:", err)
		return err
	}
	defer src.Close()

	err = model.FileService.UploadFile(fileName, &src)
	if err != nil {
		log.Logger.Error("FileService UploadFile %v:", err)
		return err
	}

	return c.JSON(http.StatusOK, fileName)
}
