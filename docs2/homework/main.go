package main

import (
	"fmt"
	"flag"
	"time"
	"log"
	"os"
	"unicode/utf8"

)

var(
	timeFormat = flag.String("timefmt",time.RFC3339, "layout for print current time")
)

func init() {
	flag.Parse()
}

func main() {
	log.SetPrefix("[demo]")
	log.SetFlags(0)

	fmt.Println("current time:",time.Now().Format(*timeFormat))
	fmt.Println("args(raws)",os.Args)
	fmt.Println("args (after flag parsing):", flag.Args())

	text :="Go语言"
	fmt.Printf("text=%q bytes=%d runes=%d\n ",text,len(text),utf8.RuneCountInString(text))
	// idx是UTF-8 字节偏移， 不是字符序号i，中文占用3字节，所以出现偏移
	for idx,r := range text{
		fmt.Printf("idx=%d,rune=%c\n",idx,r)
	}
}
