package main
import (
	"fmt"
	"errors"
	"log"
	"runtime/debug"
)

// sentinel error
var ErrNotFound = errors.New("not found")

type ValidationError struct{
	Field string
	Reason string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: field %s, reason %s", e.Field, e.Reason)
}

func load1(id int) (string,error) {
	if id ==0 {
		return "",fmt.Errorf("load1:%w",ErrNotFound)
	}
	if id <0 {
		return "",&ValidationError{Field:"id",Reason:"must be positive"}
	}
	return fmt.Sprintf("item-%d",id),nil
}

//load 模拟底层加载

func load(id int) (string ,error) {
	if id == 0 {
		return "",ErrNotFound
	}
	if id < 0 {
		return "",&ValidationError{Field:"id",Reason:"must be negative"}
	}
	return fmt.Sprintf("item-%d",id),nil
}

// WrappLoad 调用load 包装错误
func WrappLoad(id int) (error) {

    _,err := load(id)
    if err != nil {
        return fmt.Errorf("WrappLoad %d:%w",id,err)
    }
    return nil
}
// safeCall 调用fn 并在当前gorotine 中rover panic
func safeCall(fn func()){
	defer func() {
		if r:=recover();r!=nil {
			log.Printf("panic:%v\n%s",r,debug.Stack())
		}
	}()
	fn()
}

func demoWrapLoad(){
	for _,id := range []int{1,0,-1} {
	 if err := WrappLoad(id); err != nil {
			switch {
				case errors.Is(err,ErrNotFound):
					fmt.Println("not found")
				default:
					var vErr *ValidationError
					if errors.As(err,&vErr) {
						fmt.Printf("validation error: field %s, reason %s\n",vErr.Field,vErr.Reason)
					}else{
						fmt.Println("unknown error")
					}
			}
		} else {
			fmt.Println("loaded",id)
		}
	}

}

func main() {
	demoWrapLoad()
	safeCall(func() {
		panic("boom")
	})
}

//错误用返回值处理（Error is value），异常用 recover 兜底（Panic is exception）**。

/**
 * errors.Is`**：用于判断错误是否为某个**特定实例**（如：是不是那条特定的“不通过”消息）。
 2.  **`errors.As`**：用于判断错误是否为某个**特定类型**（如：是不是“验证类错误”这种类型，以便读取里面的参数）。
 3.  **`debug.Stack()`**：是程序的“黑匣子”，在发生意外（Panic）时通过 `defer` 机制将其导出，用于排查崩溃原因
 */

 /**
  * graph TD
      %% 第一部分：错误处理流程
      subgraph Error_Handling_Flow [错误处理演示: demoWrapLoad]
          A[开始循环 ID: 1, 0, -1] --> B{调用 WrappLoad}

          B -->|内部调用| C[load 函数]

          C -->|ID=1| C1[返回 成功结果]
          C -->|ID=0| C2[返回 ErrNotFound]
          C -->|ID=-1| C3[返回 ValidationError 结构体]

          C2 -.-> D[fmt.Errorf 包装: %w]
          C3 -.-> D

          D --> E{errors.Is 检查?}
          E -->|是 ErrNotFound| F[打印: not found]
          E -->|不是| G{errors.As 转换?}

          G -->|匹配 ValidationError| H[提取字段并打印: field, reason]
          G -->|不匹配| I[打印: unknown error]
      end

      %% 第二部分：Panic 与 Recover 流程
      subgraph Panic_Recover_Flow [异常捕获演示: safeCall]
          J[调用 safeCall] --> K[注册 defer 函数]
          K --> L[执行匿名函数 fn]
          L --> M[触发 panic 'boom']

          M --> N{defer 执行 recover}
          N -->|捕获到 panic| O[调用 debug.Stack]
          O --> P[打印堆栈日志并保持程序运行]
          N -->|无 panic| Q[正常结束]
      end

      %% 关联
      Error_Handling_Flow -->|完成后| Panic_Recover_Flow

  */
