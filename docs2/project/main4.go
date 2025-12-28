package main

import (
	_ "net/http/pprof" // 匿名导入，执行init 函数， 注册 pprof handler 到默认http.DefaultServeMux
	"log"
	"net/http"
)

func main() {
	log.Println("pprof ready on http://localhost:6060/debug/pprof")
	_ = http.ListenAndServe("localhost:6060", nil)
}

// fmt/print 与调试
// fmt.Print/Printf/Println 输出到标准输出；常用占位：%v 值、%+v 展示字段名、%#v Go 语法表示、%T 类型、%q 打印带引号字符串/rune。
// fmt.Sprintf 返回字符串便于日志或拼接；fmt.Errorf 搭配 %w 包装错误。
/**
 * Go设计的简洁性
 包管理为何这样设计：沿用“路径即标识”的思路，直接从远端仓库拉源码，避开单一中心；搭配 MVS + go.sum 保证构建可重复且解析规则简单，团队容易理解。
 大小写导出规则：首字母大写即公开，省去 public/private/protected 关键字，降低语言表面积；调用方读到名称即可判断可见性，统一风格减少认知切换。
 包名/标识符约定：包名短小、与目录一致、小写无下划线，标识符用驼峰，减少样板与噪音；强调“写起来少、读起来顺”，符合 Go “少即是多”的设计理念
 */
