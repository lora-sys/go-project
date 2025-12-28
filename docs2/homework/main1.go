package main
import "fmt"

type Person struct {
	Name string
	Age int
}

func main(){
	//  │ fmt.Println │ 直接打印字符串，不解析占位符 %p、%v、%s │
  //    │ fmt.Printf  │ 解析占位符，按格式输出
	p := &Person{Name:"张三",Age:25}
	fmt.Printf("Name: %p\n",p)
	fmt.Printf("Name: %v\n",*p)
	fmt.Printf("Name: %s\n",p.Name)
	fmt.Printf("Name memory: %s\n",(*p).Name)
	// 自动解引用
	p.Name= "李四"
	fmt.Println("Name:",p.Name)
}
