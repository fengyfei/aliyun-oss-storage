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
*     Initial: 2017/08/04          Yusan Kurban
 */

package middleware

import (
	"errors"

	"github.com/labstack/echo"
	jwt "github.com/dgrijalva/jwt-go"

	"aliyun-oss-storage/log"
	"fmt"
)

const (
	reqTime = "requesttime"
	proName = "project"
	header = "authorization"
	schema = "bearer"
	stand   = 5
	algorithm = "HS256"
)

type (
	tokenHeader struct {
	Alg 	string
	Typ 	string
}
	payload struct {
	Project 	string		`json:"project"`
	RequestTime	int64		`json:"requesttime"`
}
	jwtExtra func(echo.Context) (string, error)

	jwtConfig struct {
		keyFunc 	jwt.Keyfunc
	}

)

var (
	ErrUnauthorized = errors.New("Authorization failed")
	ErrInvalidJWT	= errors.New("Missing or invalid jwt header	")
	ErrNotFound		= errors.New("Project information not found")
)

// CheckJWT check project token and if it timeout .
func ParseJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		pro := projectFromHeader(proName)
		name, err := pro(c)
		if err != nil {
			log.Logger.Error("Query header crash with:", err)
			return err
		}

		secret := getSecret(name)
		if secret == "" {
			log.Logger.Error("Get project secret error:", ErrNotFound)
			return ErrNotFound
		}

		extractor := jwtFromHeader(header, schema)
		token, err := extractor(c)
		if err != nil {
			log.Logger.Error("Query header crash with:", err)
			return err
		}

		err = decodingJWT(token, secret)
		if err != nil {
			log.Logger.Error("Parser jwt crash with:", err)
			return err
		}

		c.Set(proName, name)

		return next(c)
	}
}

func decodingJWT(token, secret string) error {
	var conf jwtConfig

	conf.keyFunc = func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != algorithm {
			return nil, fmt.Errorf("Unexpected algorithm method=%v", t.Header["alg"])
		}
		return secret, nil
	}

	tt := new(jwt.Token)
	tt, err := jwt.Parse(token, conf.keyFunc)
	if err == nil && tt.Valid {
		return nil
	}

	return ErrUnauthorized
}

func getSecret(project string) string {
	return project
}

func generateJWT(key, project string, req int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims[proName] = project
	claims[reqTime] = req

	t, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return t, nil
}

func jwtFromHeader(header, authSchema string) jwtExtra {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		l := len(authSchema)
		if len(auth) > l + 1 && auth[:l] == authSchema {
			return auth[l + 1:], nil
		}
		return "", ErrInvalidJWT
	}
}

func projectFromHeader(project string) jwtExtra {
	return func (c echo.Context) (string, error) {
		pro := c.Request().Header.Get(project)
		if len(pro) > 3 {
			return pro, nil
		}
		return "", ErrInvalidJWT
	}
}
