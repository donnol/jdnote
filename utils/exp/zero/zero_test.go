package zero

import (
	"fmt"
	"testing"
)

var (
	a int

	b float64

	c string

	e [2]int

	f  []int
	fi = []int{}

	g  map[int]int
	gi = map[int]int{}

	h chan int
	// 没有字面量，必须用make

	i  interface{}
	ii = interface{}(&a)

	fn  func()
	fni = func() {}

	s  struct{}
	si = struct{}{}

	sp  *struct{}
	spi = &struct{}{}
)

func TestZero(t *testing.T) {
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v, %v, %p, %v, %v, %v, %v, %v, %p, %v, %v\n",
		a,
		b,
		c,
		e,
		f,
		g,
		h,
		i,
		fn,
		s,
		sp,
		fi,
		gi,
		ii,
		fni,
		si,
		spi,
	)

	fmt.Printf("%p, %p, %p, %v, %p, %p, %p, %p, %p, %p, %p\n",
		// 0x0，也就是nil
		f,
		g,
		h,
		i,
		fn,
		sp,

		// 非0x0，也就是非nil
		fi,  // 0x64f9c0
		gi,  // 0xc0000681b0
		ii,  // 0x64fa40
		fni, // 0x509aa0
		spi, // 0x64f9c0
	)

	// From golang runtime malloc.go
	// // base address for all 0-byte allocations
	// var zerobase uintptr
	// 结合运行时里的growslice可以看出，上面的切片和struct指针字面量的值就是zerobase的值
	// 那么，zerobase的值在什么时候，哪行代码里初始化呢？
	// 其实并不需要赋值，用的其实是它的地址；跟context包里的cancelCtxKey一样用法。
}
