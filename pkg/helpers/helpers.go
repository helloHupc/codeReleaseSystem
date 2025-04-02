package helpers

import (
	"math/rand"
	"reflect"
	"time"
)

// Empty 类似于 PHP 的 empty() 函数
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Slice, reflect.Array:
		return v.Len() == 0
	case reflect.Map:
		return v.Len() == 0
	default:
		return false
	}
}

// GenerateRandomNumber 生成指定长度的随机数字字符串
func GenerateRandomNumber(length int) string {
	if length <= 0 {
		return ""
	}

	// 使用独立的随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbers := "0123456789"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = numbers[r.Intn(len(numbers))]
	}

	return string(result)
}
