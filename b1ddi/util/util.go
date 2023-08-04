package util

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ParseError = "unable to parse key '%s': %v"
)

func ToInt(d map[string]interface{}, key string) (int, bool, error) {
	var (
		i   int
		err error
		ok  bool
		val interface{}
	)
	if val, ok = d[key]; ok {
		i, err = strconv.Atoi(val.(string))
		if err != nil {
			err = errors.New(fmt.Sprintf(ParseError, key, err))
		}
	}
	return i, ok, err
}

func ToBool(d map[string]interface{}, key string) (bool, bool, error) {
	var (
		b   bool
		err error
		ok  bool
		val interface{}
	)
	if val, ok = d[key]; ok {
		b, err = strconv.ParseBool(val.(string))
		if err != nil {
			err = errors.New(fmt.Sprintf(ParseError, key, err))
		}
	}
	return b, ok, err
}
