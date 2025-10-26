package helper

import (
	"strconv"
)

func StringToUintPtr(s string) (*uint, error) {
	if s == "" {
		return nil, nil
	}
	val, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return nil, err
	}
	u := uint(val)
	return &u, nil
}
