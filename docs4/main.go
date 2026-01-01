package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"go","rust","python","go","rust"}
	fmt.Println("Before sorting:", words)
	fmt.Println("After sorting:", dedup(words))

   fmt.Println("\nCount\n")
   counts:= countWords(words)
   printCounts(counts)
}
// 去重，保持第一次出现的顺序
func dedup(in []string) []string{
	seen:= make(map[string]bool,len(in))
	out:= make([]string,0,len(in))
	for _,v := range in {
		if seen[v] {
			continue;
		}
		seen[v] = true
		out = append(out,v)
	}
	return out
}

func countWords(words []string) map[string]int{
	m:= make(map[string]int,len(words))
	for _,w := range words {
		m[w]++
	}
	return m
}

// 打印map，键按字典序排序
func  printCounts(m map[string]int) {
	keys := make([]string,0,len(m))
	for k:= range m{
		keys = append(keys,k)
	}
	sort.Strings(keys)
	for _,k := range keys{
		fmt.Printf("%s->%d\n",k,m[k])
	}
}
