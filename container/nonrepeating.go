package cal

import (
	"fmt"
)

/**
  耗时

1.map的操作
2.rune的decode成utf-8


 */
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make([]int,0xffff)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI:=lastOccurred[ch]; lastI > start {
			start = lastI
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i+1
	}

	return maxLength
}


////用int数组 代替map 发现耗时增加 内存占用增加 因为for循环声明一直执行 声明数组空间


func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(""))
	fmt.Println(
		lengthOfNonRepeatingSubStr("b"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("这里是慕课网"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("一二三二一"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(
			"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
