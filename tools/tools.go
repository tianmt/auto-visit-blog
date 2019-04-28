package tools

import (
	"errors"
	"math/rand"
	"time"
)

// 初始化随机数种子
func init() {
	rand.Seed(time.Now().Unix())
}

// 查错
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// 获取指定范围随机 Int 类型
func GetRandInt(start int, end int) (int, error) {
	if start >= end {
		return 0, errors.New("rand range error")
	}

	return rand.Intn(end-start) + start, nil
}

// 返回指定范围随机 time.Duration 类型
func GetRandDurationInt(start int, end int) (time.Duration, error) {
	tmp_int, err := GetRandInt(start, end)

	if err != nil {
		return 0, err
	}

	return time.Duration(tmp_int), nil
}
