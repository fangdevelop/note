package note

import (
	"fmt"
	"gonote/util"
	"math/rand"
	"sort"
	"time"
)

//7.1 递归
var fibonacciRes []int

func fibonacci(n int) int {
	if n < 3 {
		return 1
	}
	if fibonacciRes[n] == 0 {
		fibonacciRes[n] = fibonacci(n-2) + fibonacci(n-1)
	}
	return fibonacciRes[n]
}
func Recursion() {
	n := 45
	fibonacciRes = make([]int, n+1)
	fmt.Printf("第%v位斐波那契数是%v\n", n, fibonacci(n))
}

//7.2 闭包
func closureFunc() func(int) int {
	i := 0
	return func(n int) int {
		fmt.Printf("本次调用接收到n=%v\n", n)
		i++
		fmt.Printf("匿名工具函数被第%v次调用\n", i)
		return i
	}
}

func Closure() {
	f := closureFunc()
	f(2)
	f(4)
	f = closureFunc()
	f(6)
}

//7.3 排序
//7.3.1 冒泡排序
func BubbleSort(s []int) {
	lastIndex := len(s) - 1
	for i := 0; i < lastIndex; i++ {
		for j := 0; j < lastIndex-i; j++ {
			if s[j] < s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

//7.3.2 选择排序
func SelectionSort(s []int) {
	lastIndex := len(s) - 1
	for i := 0; i < lastIndex; i++ {
		min := i
		for j := i + 1; j <= lastIndex; j++ {
			if s[min] > s[j] {
				min = j
			}
		}
		if min != i {
			s[i], s[min] = s[min], s[i]
		}
	}
}

//7.3.3 插入排序
func InsertionSort(s []int) {
	for i := 1; i < len(s); i++ {
		t := s[i]
		j := i - 1
		for ; j >= 0 && s[j] > t; j-- {
			s[j+1] = s[j]
		}
		if j != i-1 {
			s[j+1] = t
			//fmt.Println("s=", s)
		}
	}
}

//7.3.4 快速排序
func QuickSort(s []int, leftIndex, rightIndex int) {
	if leftIndex < rightIndex {
		pivot := s[rightIndex]
		var rs []int
		l := leftIndex
		for i := leftIndex; i < rightIndex; i++ {
			if s[i] > pivot {
				rs = append(rs, s[i])
			} else {
				s[l] = s[i]
				l++
			}
		}
		s[l] = pivot
		copy(s[l+1:], rs)
		if leftIndex < l-1 {
			QuickSort(s, leftIndex, l-1)
		}
		if l+1 < rightIndex {
			QuickSort(s, l+1, rightIndex)
		}
	}
}

func Sort() {
	n := 10
	s := make([]int, n)
	seedNum := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		rand.Seed(seedNum)
		s[i] = rand.Intn(10001)
		seedNum++
	}
	fmt.Println("排序前:", s)
	QuickSort(s, 0, len(s)-1)
	fmt.Println("排序后:", s)
}

//7.5.2 二分查找
func BinarySearch(s []int, key int) int {
	startIndex := 0
	endIndex := len(s) - 1
	midIndex := 0
	for startIndex <= endIndex {
		midIndex = startIndex + (endIndex-startIndex)/2
		if s[midIndex] < key {
			startIndex = midIndex + 1
		} else if s[midIndex] > key {
			endIndex = midIndex - 1
		} else {
			return midIndex
		}
	}
	return -1
}
func BinarySearchTest() {
	s := make([]int, util.RandInt(1000)+1)
	for i := 0; i < len(s); i++ {
		s[i] = util.RandInt(1000)
	}
	sort.Ints(s)
	i := BinarySearch(s, 555)
	if i == -1 {
		fmt.Println("没有找到555")
	} else {
		fmt.Println("555的下标是", i)
	}
}
