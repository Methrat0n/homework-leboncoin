package FizzBuzzHandler

import (
	"errors"
	"leboncoin/model"
	"strconv"

	routing "github.com/qiangxue/fasthttp-routing"
)

type FizzBuzzHandler interface {
	GetFizzBuzz(request *routing.Context) error
}

type defaultFizzBuzzHandler struct {
	defaultLimit int64
}

func New(defaultLimit int64) FizzBuzzHandler {
	return defaultFizzBuzzHandler{defaultLimit}
}

func (handler defaultFizzBuzzHandler) GetFizzBuzz(request *routing.Context) error {
	int1, int2, limit, str1, str2, err := parseQueryString(request, handler.defaultLimit)
	if err != nil {
		return err
	}
	res := model.FizzBuzz(int1, int2, limit, str1, str2)
	request.SetStatusCode(200)
	request.Response.SetBody([]byte(res))
	return nil
}

func parseQueryString(request *routing.Context, defaultLimit int64) (int1 int64, int2 int64, limit int64, str1 string, str2 string, err error) {
	request.QueryArgs().VisitAll(func(key, value []byte) {
		stringKey := string(key)

		limit = defaultLimit

		if stringKey == "int1" {
			if int1, err = strconv.ParseInt(string(value), 10, 0); err != nil {
				return
			}
		}
		if stringKey == "int2" {
			if int2, err = strconv.ParseInt(string(value), 10, 0); err != nil {
				return
			}
		}
		if stringKey == "limit" {
			if limit, err = strconv.ParseInt(string(value), 10, 0); err != nil {
				return
			}
		}
		if stringKey == "str1" {
			str1 = string(value)
		}
		if stringKey == "str2" {
			str2 = string(value)
		}
		return
	})
	if err == nil && (int1 == 0 || int2 == 0 || str1 == "" || str2 == "") {
		err = errors.New("Missing mandatory query string parameter")
	}
	return
}
