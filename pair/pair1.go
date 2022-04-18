/*
变量中都会有一个 pair
	pair<type: , value: >
	变量中对应的信息（类型和值）
变量在持续赋值中，他的 pair 是保持持续不变的
*/
package main

import "fmt"

func main() {
	var a string
	// pair<type:string, value:"GoLand">
	a = "GoLand"

	// pair<type:string, value:"GoLand">
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}
