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
 *     Initial: 2017/08/01        Liu JiaChang
 */

package main

import (
	"github.com/spf13/viper"

	"aliyun-oss-storage/log"
)

type workServerConfig struct {
	Address         string
	IsDebug         bool
	EndPoint        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
	BoltPath        string
	ProjectList     map[string]string
}

var (
	// Conf is a config
	conf        *workServerConfig
	projList    map[string]string
)

// ReadConfiguration initial read config
func readConfiguration() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf = &workServerConfig{
		Address:         viper.GetString("server.address"),
		IsDebug:         viper.GetBool("server.debug"),
		EndPoint:        viper.GetString("ali.endpoint"),
		AccessKeyID:     viper.GetString("ali.access.id"),
		AccessKeySecret: viper.GetString("ali.access.secret"),
		BucketName:      viper.GetString("ali.bucket"),
		BoltPath:        viper.GetString("bolt.path"),
		ProjectList:     viper.GetStringMapString("project"),
	}
}

func readProjectList() {
	if err := viper.ReadInConfig(); err != nil {
		log.Logger.Error("Read config error:", err)

		return
	}

	projList = viper.GetStringMapString("project")
}
