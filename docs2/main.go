
package main
// var name string;
// var name string = "John Doe"
// var age int;
// var count int = 3;
// var a,b=1,"hi"；
// var  gemini := "Gemini";
// gemini:="hello";

import (
	"fmt";
	"strconv";
)

// 常量：const name = value，值必须在编译期可确定（数字、字符串、布尔、rune）；不可使用 :=；常见配合 iota 自增生成枚举值。
const Pi = 3.1415
// 与变量区别，变量可用:=,可以在后续被修改，可以是运行期间

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)
const (
	KB = 1<<(10*iota) // 1<<0 =1
	MB = 1<<(10*iota) // 1<<10 =1024
	GB = 1<<(10*iota) // 1<<20 =1048576
)
//  错误:  const name :="HO"
// const time = time.Now(); // 运行时候无法确定



// //   iota 特点：
//  - 每行自动递增
//  - 从 0 开始
//  - 同一个 const 块内共享


// 类型推断
//
var x = 10
var y = "hello"
var z = true
const (
	StausOK=20
	StatusNotFound=404
)
var(
	port int  // 0
	host string // ""
)

// 短变量声明(:=)
// 函数内使用 语法 name := expr，自动推断类型
// 至少有一个新变量才允许：x := 1; x, y := 2, 3（y 为新变量）；否则编译报错。
// 常见搭配：v, err := someCall()

// 基本类型与转换
//
var i int = 10
var f float64 = float64(i)
var s string = fmt.Sprintf("%d", i)
// 数值：int/uint（与架构位数相关）、int8/16/32/64、float32/64、complex64/128。
// 布尔：bool 只能是 true/false，不允许与整数混用。
// 字符：byte 是 uint8 的别名，rune 是 int32 的别名，常用于区分按字节或按 Unicode 码点处理。
// 强制转换：Go 不做隐式类型转换，不同整数/浮点类型需显式转换；字符串与数字需配合 strconv 或 fmt。


// for i:=0;i<10;i++{
// 	fmt.Println("%c",str[i])//输出乱码
// }
//

func main() {
   str:="go 语言"
	fmt.Println(len(str))
	fmt.Println([]byte(str))
	fmt.Println(len([]rune(str)))
	// fmt.Println(string([]byte(str)))
	// fmt.Println(string([]rune(str)))
//6
// [228 184 150 231 149 140]
// 世界
// 世界
// %d->%c 10 19990
// %d->%c 10 30028


// 5
// [104 101 108 108 111]
// hello
// hello
// %d->%c 10 104
// %d->%c 10 101
// %d->%c 10 108
// %d->%c 10 108
// %d->%c 10 111
var i int32 = 42
f := float64(i)
s := fmt.Sprintf("%d", i) // 数字转字符串
n, err := strconv.Atoi("123")
fmt.Println(s)
fmt.Println(f)
fmt.Println(n)
fmt.Println(err) //nli -> null
	// for _, r:=range str{
 // fmt.Println("%d->%c",i,r)
 // }
}
