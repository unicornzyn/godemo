package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	res := Add(2, 3)
	if res != 5 {
		t.Fatalf("Add(2, 3)执行错误 期望值 %v, 实际值 %v", 5, res)
	}
	t.Logf("Add(2, 3)执行正确")
}

func TestSub(t *testing.T) {
	res := Sub(5, 3)
	if res != 2 {
		t.Fatalf("Sub(5, 3)执行错误 期望值 %v, 实际值 %v", 2, res)
	}
	t.Logf("Sub(5, 3)执行正确")
}
