package main

import (
	"fmt"
	"strconv"
)

type Printer interface {
	Print() string
}

type User struct {
	Name string
	Age int
}

func (u User) Print() string {
	return fmt.Sprintf("User(name=%s, age=%s)", u.Name, strconv.Itoa(u.Age))
}
type Product struct {
	ID int
	Name string
	PriceCents int
}

func (p Product) Print() string {
	return fmt.Sprintf("Product(id=%d, name=%s, price=%s)", p.ID, p.Name, strconv.Itoa(p.PriceCents))
}

func LogAll(ps []Printer) {
	for _, p := range ps {
		fmt.Println(p.Print())
	}
}

type Box struct {
	val any // any=interface{}
}

func (b Box) AsString() string {
	switch v := b.val.(type) {
		case nil:
			return "<nil>"
		case fmt.Stringer:
			return v.String()
		case string:
		    return v
		case int:
			return strconv.Itoa(v)
		case float64:
			return strconv.FormatFloat(v, 'f', -1, 64)
		default:
			return fmt.Sprintf("%v", v)
	}
}
/**
 * 第一层（编译时）**：通过 `Printer` 接口定义规范，确保 `User` 和 `Product` 有共同的行为。
 2.  **第二层（内存层）**：展示了如何通过显式循环将具体对象提升为接口对象，解决类型系统的安全限制。
 3.  **第三层（运行时）**：通过 `Type Switch` 实现了对未知数据的“反向解析”，让代码具备了处理动态数据的能力
 */



func main() {
	users := []User{
		User{Name: "Alice", Age: 30},
		User{Name: "Bob", Age: 25},
	}
	products := []Product{
		Product{ID: 1, Name: "Apple", PriceCents: 100},
		Product{ID: 2, Name: "Banana", PriceCents: 50},
	}
    var items []Printer
    for _,u := range users {
    items = append(items, u)
    }
    for _,p := range products {
    items = append(items, p)
    }
    LogAll(items)
    box := []Box {
    {val: "Hello"},
    {val: 123},
    {val: 3.14},
    {val:users[0]},
    {val:nil},
    }

    for _,b := range box {
        fmt.Println(b.AsString())
    }
}
