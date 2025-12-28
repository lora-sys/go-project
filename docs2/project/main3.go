package main

import (
	"fmt"
	mathAlias "math"
)
// 导入写在 import 之后，单行或分组都可以；分组更便于管理第三方依赖和标准库。
// 可选用别名：import m "math"；只执行副作用用空白标识符：import _ "net/http/pprof"。
// 包可见性：标识符首字母大写即导出（对其他包可见），小写仅包内可用；与文件名、目录名无关。
// 作用域：package 级（同包文件共享）、file 级 init/var、代码块级（if/for/func 等）。避免变量遮蔽（重名导致外层变量被屏蔽）。
// internal 包：只能被其父目录及子目录导入，编译器强制限制，适合放私有实现，防止被外部依赖污染模块 API。

// //project/
// ├─ internal/auth  // 仅 project 下其他包可导入
// └─ cmd/app       // ok: import "project/internal/auth"
//    vendor/client // 不允许导入 internal/auth
var pkgLevel = "package scope"

func main() {
	msg := "block scope"
	fmt.Println(pkgLevel, msg, mathAlias.Pi)
}
