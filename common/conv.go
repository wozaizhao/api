package common

import (
	"strconv"
)

// ParseFloat 转换为数值类型
func ParseFloat(str string) (float64, error) {
	num, err := strconv.ParseFloat(str, 64)
	return num, err
}

// ParseInt 转换为整型数值类型
func ParseInt(str string) (int64, error) {
	num, err := strconv.ParseInt(str, 10, 64)
	return num, err
}
