package tool

import "reflect"

// 排序

/*
 arr 可以是各种类型
sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] > arr[j] //从大到小排列
	})

*/

// 查看interface类型
func GetInterfaceType(i interface{}) reflect.Type {
	val := reflect.ValueOf(i)
	return reflect.Indirect(val).Type()
}
