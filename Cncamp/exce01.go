package main

import "fmt"

/*
	给定一个字符串数组
	["I","am","stupid","and","weak"]
	用 for 循环遍历该数组并修改为
	["I","am","smart","and","strong"]
*/

func main() {
	// 定义一个字符串数组
	words := []string{"I", "am", "stupid", "and", "weak"}

	for i := range words {
		if words[i] == "stupid" {
			words[i] = "smart"
		} else if words[i] == "weak" {
			words[i] = "strong"
		}
	}

	// 输出修改后的数组
	fmt.Println(words)
}
