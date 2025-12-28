package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// 打印当前时间
	fmt.Println("当前时间:", time.Now().Format("2006-01-02 15:04:05"))

	// 解析命令行参数
	flag.Parse()
	fmt.Println("命令行参数:", flag.Args())

	// 读取字符串变量
	var str string
	fmt.Print("请输入一个字符串: ")
	fmt.Scanln(&str)

	// 输出字节长度和 rune 数
	fmt.Printf("字节长度: %d\n", len(str))
	fmt.Printf("rune 数: %d\n", len([]rune(str)))

	// 用 range 打印每个字符
	fmt.Println("每个字符:")
	for i, r := range str {
		fmt.Printf("索引: %d, 字符: %c, Unicode: %U\n", i, r, r)
	}
}
