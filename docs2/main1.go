// code/02/main.go
package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "gopher", "who to greet")
//init() 用于做轻量初始化，每个文件可有多个，执行顺序
// 导入包的 init() → 当前包各文件自上而下的 init() → 最终执行
// 最终执行 main()。常见用途：注册驱动、设置日志、校验环境变量，也可以在这里提前解析命令行 flag
// 类似python argparse
func init() {
	flag.Parse()
	fmt.Println("init runs before main, name flag =", *name)
}

func main() {
	fmt.Println("hello,", *name)
}
