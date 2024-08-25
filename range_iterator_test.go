package test

import (
	"fmt"
	"sync"
	"testing"
)

// TL;DR go 1.23 for-range loop ใช้ได้กับ
// func(func() bool)
// func(func(K) bool)
// func(func(K, V) bool)

func TestFuncRange122(t *testing.T) {

	// จากข้างบน เขียนบน go 1.22 จะเขียนแบบนี้
	var m sync.Map

	m.Store("iter0", 11)
	m.Store("iter1", 12)
	m.Store("iter2", 13)

	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

	// ถ้าเป็น go 1.23
	for key, val := range m.Range {
		fmt.Println(key, val)
	}
}

// func(func() bool)
func iter0(yield func() bool) {
	for range 3 {
		if !yield() {
			return
		}
	}
}

func TestFuncRange0(t *testing.T) {
	for range iter0 {
		t.Log("iter0")
	}
}

// func(func(K) bool)
func iter1(yield func(i int) bool) {
	for i := range 4 {
		if !yield(i) {
			return
		}
	}
}

func TestFuncRange1(t *testing.T) {

	//  ใช้ len check ไม่ได้ แตก
	//	fmt.Println(len(iter1))

	// คล้าย ๆ เขียน for กับรอ <-chan แต่ต้องใช้ chan
	for i := range iter1 {
		t.Log("iter1", i)
	}

}

// func(func(K, V) bool)
func iter2(yield func(int, int) bool) {
	for i := range 3 {
		if !yield(i, i+1) {
			return
		}
	}
}

func TestFuncRange2(t *testing.T) {
	for i, e := range iter2 {
		t.Log("iter2", i, e)
	}
}
