nums:=[]int{1,2,3,4,5}
sub:=nums[1:3] // len=2 cap=3 ,cap =cap(s)-a ，cap（s） 代表原切片容量
sub[0]=20
safe:=nums[1:3:3] //// cap=2，后续 append 不影响 nums
/**
 * 与 Python 不同：
 * Go 不支持负索引或步长，
 * 也不支持 b 超过 len(s)（full slice 时不超过 cap）。省略 a/b 时分别默认为 0 和 len(s)。
 */


 func mutate(s[]int){
 s[0]=99
 }
