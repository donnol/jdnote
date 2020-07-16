package main

import "fmt"

func fibonacci() func() int {
	var i int
	return func() int {
		r := fi(i)
		i++
		return r
	}
}

var m = make(map[int]int)

func fi(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	if v, ok := m[i]; ok {
		return v
	}
	return fi(i-1) + fi(i-2)
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
