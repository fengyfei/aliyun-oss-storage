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
*     Initial: 2017/08/02          Yusan Kurban
*     Modify : 2017/08/05          Yang Chenglong
 */

package model

import (
	"bufio"
	"mime/multipart"

	"github.com/fengyfei/aliyun-oss-storage/ali"
)

// FileProviderService is an receiver .
type FileProviderService struct{}

// FileService is encapsulation of provider.
var FileService *FileProviderService

const (
	// LocalURL is server url.
	LocalURL = "https://exmaple.com/"
	// LocalPath is where object storage in server.
	LocalPath = "static/"
)

// GetFile is get a file by key.
func (fs *FileProviderService) GetFile(objectKey, project string) (string, error) {
	var err error

	path := LocalPath + "/" + project + "/" + objectKey
	err = ali.Bucket.GetObjectToFile(objectKey, path)
	if err != nil {
		return "", err
	}

	return LocalPath + path, nil
}

func (fs *FileProviderService) UploadFile(objectKey string, file *multipart.File) error {
	err := ali.Bucket.PutObject(objectKey, bufio.NewReader(*file))
	if err != nil {
		return err
	}

	return nil
}
