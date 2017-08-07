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
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"aliyun-oss-storage/bolt"
	"aliyun-oss-storage/log"
)

const (
	reqTime   = "requesttime"
	ProName   = "project"
	header    = "authorization"
	schema    = "bearer"
	stand     = 5
	algorithm = "HS256"
)

type (
	jwtExtra func(echo.Context) (string, error)

	jwtConfig struct {
		keyFunc jwt.Keyfunc
	}
)

var (
	ErrUnauthorized = errors.New("Authorization failed")
	ErrInvalidJWT   = errors.New("Missing or invalid jwt header	")
	ErrNotFound     = errors.New("Project information not found")
)

// ParseJWT CheckJWT check project token and if it timeout .
func ParseJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		pro := projectFromHeader(ProName)
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

		err = decodingJWT(token, secret, name)
		if err != nil {
			log.Logger.Error("Parser jwt crash with:", err)
			return err
		}

		c.Set(ProName, name)

		return next(c)
	}
}

func decodingJWT(tokenString, secret, name string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != algorithm {
			return nil, fmt.Errorf("Unexpected algorithm method=%v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims[ProName].(string) == name && float64(time.Now().Unix())-claims[reqTime].(float64) < stand {
			return nil
		}
	}

	return ErrUnauthorized
}

func getSecret(project string) string {
	return bolt.GetProjectSecure(project)
}

func jwtFromHeader(header, authSchema string) jwtExtra {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		l := len(authSchema)
		if len(auth) > l+1 && auth[:l] == authSchema {
			return auth[l+1:], nil
		}
		return "", ErrInvalidJWT
	}
}

func projectFromHeader(project string) jwtExtra {
	return func(c echo.Context) (string, error) {
		pro := c.Request().Header.Get(project)
		if len(pro) > 3 {
			return pro, nil
		}
		return "", ErrInvalidJWT
	}
}

// GenerateJWT return jwt and nil if success.
func GenerateJWT(project, secret string) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims[ProName] = project
	claims[reqTime] = time.Now().Unix()

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Logger.Error("generate jwt failed with error:", err)
	}

	log.Logger.Debug("i got token now:", t)

	err = decodingJWT(t, secret, project)
	if err != nil {
		log.Logger.Error("error", err)
	}
}
