package goAinB
//package main

import(
	"strings"
)
//判断数组是否包含某个值
func AinBArrayOneItem(item string, items []string) bool{
	for _, TheItem := range items {
		if TheItem == item {
			return true
		}
	}
	return false
}

//判断数组是否包含另一个数组
func AinBArrayItems(itemA []string, itemB []string) bool{
	countItem := 0
	for _, TheItem := range itemA {
		if AinBArrayOneItem(TheItem, itemB) {
			countItem = countItem + 1
		}
	}
	if countItem == len(itemA){
		return true
	}else{
		return false
	}
}

//判断字符串A是否被字符串B包含
func AinBStr(strA string, strB string) bool{
	i := strings.Index(strB, strA)
	if i== -1 {
		return false
	}else{
		return true
	}
}