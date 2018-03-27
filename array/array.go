package array

import (
	"fmt"
	"reflect"
)

//IsInArray 数组中是否存在此元素，存在返回true,否则false
//目前只支持[]int,[]float64,[]string
func IsInArray(elem interface{}, array interface{}) bool {

	switch array.(type) {
	case []int:
		ve, ok := elem.(int)
		if !ok {
			return false
		}

		for _, v := range array.([]int) {
			if ve == v {
				return true
			}
		}
	case []float64:
		ve, ok := elem.(float64)
		if !ok {
			return false
		}

		for _, v := range array.([]float64) {
			if ve == v {
				return true
			}
		}
	case []string:
		ve, ok := elem.(string)
		if !ok {
			return false
		}

		for _, v := range array.([]string) {
			if ve == v {
				return true
			}
		}
	default:
		fmt.Println("in array 不支持的类型判断:", reflect.TypeOf(array))
	}
	return false
}
