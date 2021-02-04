package main

import (
	"bytes"
	"sort"
	"strconv"
)

func FindMaximumXORMap(nums []int) int {
	/*
	 * 对数组切片排序，获取最大值的二进制长度
	 * strconv.FormatInt将int64转换为其他进制的字符串
	 */
	sort.Ints(nums)
	s := strconv.FormatInt(int64(nums[len(nums)-1]),2)
	l := len(s)
	maxXor := 0

	/*
	 * 循环获取num左侧bit位
	 * 例如第二次循环，将每个数的左侧两位作为prefixes的key
	 * 然后遍历prefixes，查看左侧两位组成的数中是否存在最大的异或值
	 */
	for i := l - 1; i >= 0; i-- {
		/*
		 * 每次遍历都需要清空map，go语言中最高效的方法就是重建一个map
		 */
		prefixes := make(map[int]int)
		/*
		 * 每次都要找到nums数组中是否存在p1，p2，使得p1^p2 等于currXor
		 * currXor在每轮遍历中最右侧位为1
		 */
		maxXor = maxXor << 1
		currXor := maxXor | 1
		for _, num := range nums {
			prefixes[num >> i] = 0
		}
		for p := range prefixes {
			if _,ok := prefixes[currXor^p]; ok {
				maxXor = currXor
				break
			}
		}
	}
	return maxXor
}

func Int2BinaryStr(i int, l int) string {
	i2 := strconv.FormatInt(int64(i),2)
	if l <= len(i2) {
		return i2
	}

	buffer := bytes.NewBufferString("0")
	for i := l - len(i2) - 1; i > 0; i-- {
		buffer.WriteString("0")
	}
	buffer.WriteString(i2)
	return buffer.String()
}

type Trie struct {
	next [int32(2)]*Trie
	isEnd bool
}

func Construct() Trie {
	return Trie{next: *new([int32(2)]*Trie)}
}

func (t *Trie) Insert(value string) {
	node := t
	for _, c := range value {
		//fmt.Printf("%c\t%d\t%#v\n", c,len(node.next), node.next)
		if node.next[c-'0'] == nil {
			node.next[c-'0'] = &Trie{next: *new([int32(2)]*Trie)}
		}
		node = node.next[c-'0']
	}
	node.isEnd = true
}

func (t *Trie) FindMaxXor(value string) int {
	//fmt.Printf("%s\n", value)
	node := t
	maxXor := 0
	for _, c := range value {
		//fmt.Printf("key: %c\t%#v\n", key, node.next)
		if node.next[1^(c-'0')] == nil {
			maxXor = maxXor << 1
			node = node.next[0^(c-'0')]
		} else {
			maxXor = maxXor << 1
			maxXor = maxXor | 1
			node = node.next[1^(c-'0')]
		}
	}
	if node.isEnd {
		return maxXor
	}
	return 0
}

func FindMaximumXORTrie(nums []int) int {
	sort.Ints(nums)
	s := strconv.FormatInt(int64(nums[len(nums)-1]),2)
	l := len(s)
	obj := Construct()
	for _, num := range nums {
		obj.Insert(Int2BinaryStr(num, l))
	}
	maxXor := 0
	for _, num := range nums {
		num2BinaryStr := Int2BinaryStr(num, l)
		currXor := obj.FindMaxXor(num2BinaryStr)
		if currXor > maxXor {
			maxXor = currXor
		}
	}
	return maxXor
}