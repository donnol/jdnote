package worker

import (
	"log"
	"testing"
)

func TestWorker(t *testing.T) {
	w := New(10)
	w.Start()

	w.Push(func() error {
		panic("terriable")
		return nil
	})
	w.Push(func() error {
		for i := 0; i < 10; i++ {
			log.Printf("i: %d\n", i)
		}
		return nil
	})
	for i := 10; i < 1000; i++ {
		tmp := i
		w.Push(func() error {
			log.Printf("i: %d\n", tmp)
			return nil
		})
	}

	w.Stop()
}
