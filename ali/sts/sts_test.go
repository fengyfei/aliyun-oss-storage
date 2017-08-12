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
*     Initial: 2017/08/12          Sun Anxiang
 */

package sts

import (
	"log"
	"testing"
)

var (
	TempKey    = "LTAI24bPTHC8Tyfw"
	TempSecret = "J2Hy5FepMWDaO560yiT1Tty7LsD9Qj"
	RoleAcs    = "acs:ram::1042201469700617:role/anxiang"
)

func TestGenerateSignatureUrl(t *testing.T) {
	client := NewStsClient(TempKey, TempSecret, RoleAcs)

	url, err := client.GenerateSignatureUrl("client", "1800")
	if err != nil {
		t.Error(err)
	}

	data, err := client.GetStsResponse(url)
	if err != nil {
		t.Error(err)
	}

	log.Println("result:", string(data))
}
