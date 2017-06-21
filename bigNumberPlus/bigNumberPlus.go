package main

import (
	"strconv"
	"bytes"
	"fmt"
	"github.com/gavinzhs/utils"
)

/*
	思路:
	大数相加,不可以用int,因为承载不住
	1. 用str进行对应位数相加
	2. 反转一下利于相加
	3. 长度不一的话用0补全
	4. 对应位相加并进位
	5. 最后反转回来
 */

func main() {
	print("bigNumberPlus")

	num1 := "111111111111111111111111111111111111111111111111111122222222"
	num2 := "111111111111111111111111111111111111111111111111999"

	reverseNum1 := utils.ReverseString(num1)
	reverseNum2 := utils.ReverseString(num2)
	fmt.Printf("reverseNum1: %s, reverseNum2: %s\n", reverseNum1, reverseNum2)
	var buf bytes.Buffer
	up := 0
	for i := 0; i < len(reverseNum1) || i < len(reverseNum2); i++ {
		n1 := 0
		if i < len(reverseNum1) {
			n1, _ = strconv.Atoi(string(reverseNum1[i]))
		}

		n2 := 0
		if i < len(reverseNum2) {
			n2, _ = strconv.Atoi(string(reverseNum2[i]))
		}

		remain := (n1 + n2 + up) % 10
		up = (n1 + n2 + up) / 10
		buf.WriteString(strconv.Itoa(remain))
	}

	if up != 0 {
		buf.WriteString(strconv.Itoa(up))
	}

	rlt := utils.ReverseString(buf.String())
	fmt.Printf("rlt: %s\n", rlt)
}