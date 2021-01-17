package basictips

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// slice
func stack() {
	// create stack
	stack := make([]int, 0)
	// push
	stack = append(stack, 10)
	// pop
	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Printf("%v\n", v)
	// check is empty
	if len(stack) == 0 {
		fmt.Print("true")
	}
}

func queue() {
	// create queue
	queue := make([]int, 0)
	// enqueue
	queue = append(queue, 10)
	// dequeue
	v := queue[0]
	fmt.Println(v)
	queue = queue[1:]
	// check is empty
	if len(queue) == 0 {
		fmt.Print("true")
	}
}

// dict
func dict() {
	// create
	dict := make(map[string]int)
	// set k/v
	dict["a"] = 1
	// delete k/v
	delete(dict, "a")
	// for range
	for k, v := range dict {
		fmt.Println(k, v)
	}
}

// standard lib
// sort: https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter03/03.1.html
func gsort() {
	sort.Ints([]int{})       // int sort
	sort.Strings([]string{}) // string sort
	s := []int{3, 5, 7, 9}
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] }) // custom sort
}

// math
func gmath() {
	MaxInt32 := math.MaxInt32 // actual: 1<<31-1
	MinInt32 := math.MinInt32 // actual: -1<<31
	MaxInt64 := math.MaxInt64 // actual: 1<<63-1
	MinInt64 := math.MinInt64 // actual: -1<<63
	fmt.Println(MaxInt32, MinInt32, MaxInt64, MinInt64)
}

// copy
func gcopy() {
	a := []int{1, 2, 3, 4, 5}
	i := 2
	copy(a[i:], a[i+1:]) // 删除a[i]，可以用 copy 将i+1到末尾的值覆盖到i,然后末尾-1
	a = a[:len(a)-1]

	// make创建长度，则通过索引赋值
	n, x := 10, 7
	a = make([]int, n)
	a[n] = x
	// make长度为0，则通过append()赋值
	a = make([]int, 0)
	a = append(a, x)
}

// tips
func typeTranform() {
	// byte转数字
	s := "12345"                        // s[0] 类型是byte
	num := int(s[0] - '0')              // 1
	str := string(s[0])                 // "1"
	b := byte(num + '0')                // '1'
	fmt.Printf("%d%s%c\n", num, str, b) // 111

	// 字符串转数字
	num, _ = strconv.Atoi("12")
	str = strconv.Itoa(2)
}
