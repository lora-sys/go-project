package main

import (
	"fmt"
	"flag"
	"time"
)



func main() {
	fmt.Println("当前时间:",time.Now().Format("2006-01-02 15:04:05"))
	flag.Parse()
	fmt.Println("命令行参数:", flag.Args())
	var str string
	fmt.Println("请输入一个字符串")
	fmt.Scanln(&str)
	fmt.Println("您输入的字符串是:", str)

	fmt.Printf("字节长度:%d\n",len(str));
	fmt.Printf("rune 数：%d\n",len([]rune(str)))

	for i,r := range str{
		fmt.Printf("索引:%d\n字符:%c\nUnicode:%U\n",i,r,r)
	}

}
