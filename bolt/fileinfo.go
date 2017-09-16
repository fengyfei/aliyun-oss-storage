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

package bolt

import (
	"encoding/json"
	"errors"

	"github.com/fengyfei/aliyun-oss-storage/log"
)

var ErrNotFound = errors.New("Database Not Found")

type FileInfoServiceProvider struct{}

var (
	FileInfoService *FileInfoServiceProvider = &FileInfoServiceProvider{}
)

func (asp *FileInfoServiceProvider) CreateInfo(hash *string, project *string) error {
	tx, err := BoltDb.Begin(true)
	if err != nil {
		log.Logger.Error("[create] begin txn error: %v", err)
		return err
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte("userdata"))

	if ab, err := json.Marshal(*project); err != nil {
		log.Logger.Error("marshal error: %v", err)
		return err
	} else if err := bucket.Put([]byte(*hash), ab); err != nil {
		log.Logger.Error("put error: %v", err)
		return err
	}

	return tx.Commit()
}

func (asp *FileInfoServiceProvider) GetInfo(hash *string) (string, error) {
	var info string

	tx, err := BoltDb.Begin(false)
	if err != nil {
		log.Logger.Error("[get] begin txn error: %v", err)
		return info, err
	}
	defer tx.Rollback()

	if v := tx.Bucket([]byte("userdata")).Get([]byte(*hash)); v == nil {
		log.Logger.Debug("don't have record")
		return info, ErrNotFound
	} else if err := json.Unmarshal(v, &info); err != nil {
		log.Logger.Error("unmarshal error: %v", err)
		return info, err
	}

	return info, err
}
