package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"sort"
	"strconv"
)

// MD5 字符串加密
func MD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func StringIn(target string, array []string) bool {
	sort.Strings(array)
	index := sort.SearchStrings(array, target)
	// index的取值：[0,len(str_array)]
	// 需要注意此处的判断，先判断 &&左侧的条件，如果不满足则结束此处判断，不会再进行右侧的判断
	if index < len(array) && array[index] == target {
		return true
	}
	return false
}

func StringCovertInt64(strArr []string) []int64 {
	res := make([]int64, len(strArr))
	for index, val := range strArr {
		res[index], _ = strconv.ParseInt(val, 10, 64)
	}

	return res
}

// ArrayUniqueValue 数组唯一值
func ArrayUniqueValue[T any](arr []T) []T {
	size := len(arr)
	result := make([]T, 0, size)
	temp := map[any]struct{}{}
	for i := 0; i < size; i++ {
		if _, ok := temp[arr[i]]; ok != true {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}
	return result
}

// ArrayContainValue 数组是否包含值
func ArrayContainValue(arr []int64, search int64) bool {
	for _, v := range arr {
		if v == search {
			return true
		}
	}
	return false
}

// Intersect 切片插入
func Intersect(slice1 []int64, slice2 []int64) []int64 {
	m := make(map[int64]int64)
	n := make([]int64, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			n = append(n, v)
		}
	}
	return n
}

// Difference 辨别切片
func Difference(slice1 []int64, slice2 []int64) []int64 {
	m := make(map[int64]int)
	n := make([]int64, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}
	for _, v := range slice1 {
		times, _ := m[v]
		if times == 0 {
			n = append(n, v)
		}
	}
	return n
}
