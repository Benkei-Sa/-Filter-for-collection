package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filter(predicate func(int) bool, iterable []int) []int {
	var src []int
	for _, n := range iterable {
		b := predicate(n)
		if b == true {
			src = append(src, n)
		}
	}
	return src
	// отфильтруйте `iterable` с помощью `predicate`
	// и верните отфильтрованный срез
}

func main() {
	src := readInput()
	result := func(n int) (b bool) {
		if n%2 == 0 {
			b = true
		} else {
			b = false
		}
		return b
	}

	res := filter(result, src)

	// отфильтруйте `src` так, чтобы остались только четные числа
	// и запишите результат в `res`
	// res := filter(...)
	fmt.Println(res)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
