/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co.,ltd..
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
 *     Initial: 2017/08/05        Liu JiaChang
 */

package bolt

import (
	"time"

	"github.com/boltdb/bolt"

	"aliyun-oss-storage/general"
)

var BoltDb *bolt.DB

// InitBolt 初始化 bolt 数据库
func InitBolt(path string) error {
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: time.Second})
	if err != nil {
		panic(err)
	}

	BoltDb = db

	tx, err := BoltDb.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err = tx.CreateBucketIfNotExists([]byte(general.BoltProjectList)); err != nil {
		return err
	}

	if _, err = tx.CreateBucketIfNotExists([]byte(general.BoltUserData)); err != nil {
		return err
	}

	return tx.Commit()
}

// UseProjectList 用配置文件读取的 map 更新数据库
func UseProjectList(plist map[string]string) error {
	return BoltDb.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(general.BoltProjectList))
		cursor := bucket.Cursor()

		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			if secure, ok := plist[string(key)]; ok {
				if secure != string(value) {
					if err := bucket.Put(key, []byte(secure)); err != nil {
						return err
					}
				}
				delete(plist, string(key))
			} else {
				if err := bucket.Delete([]byte(key)); err != nil {
					return err
				}
			}
		}

		for name, secure := range plist {
			if err := bucket.Put([]byte(name), []byte(secure)); err != nil {
				return err
			}
		}

		return nil
	})
}

// GetProjectList 从数据库获取 map
func GetProjectList() map[string]string {
	projectList := make(map[string]string)

	BoltDb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(general.BoltProjectList))
		cursor := bucket.Cursor()

		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			projectList[string(key)] = string(value)
		}

		return nil
	})

	return projectList
}

func GetProjectSecure(pname string) string {
	var secure []byte

	BoltDb.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(general.BoltProjectList))
		secure = bucket.Get([]byte(pname))

		return nil
	})

	return string(secure)
}
