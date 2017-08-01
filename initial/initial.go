/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co,ltd..
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

package initial

import (
	"Winston/config"
	"Winston/orm/cockroach"
)

var (
	// DbConnPool connect pool
	DbConnPool *cockroach.Pool
)

// InitConf intiacl config
func InitConf() {
	config.ReadConfiguration()
}

// InitCockroachPool initial cockroach pool
func InitCockroachPool() {
	DbConnPool = cockroach.NewPool(config.Conf.DbUrl, config.Conf.PoolSize)

	if DbConnPool == nil {
		panic("Cockroach DB connection error")
	}
}
