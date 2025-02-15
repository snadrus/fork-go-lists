package lists

import (
	"sync"
	"testing"
)

func TestStack(t *testing.T) {
	var stack Stack[int]
	var v int
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if !stack.Pop(&v) {
		t.Fatal("x")
	} else if v != 3 {
		t.Fatal(v)
	}
	if !stack.Pop(&v) {
		t.Fatal("y")
	} else if v != 2 {
		t.Fatal(v)
	}
	if !stack.Pop(&v) {
		t.Fatal("z")
	} else if v != 1 {
		t.Fatal(v)
	}
	if stack.Pop(&v) {
		t.Fatal(v)
	}
}

func TestStackRace(t *testing.T) {
	var wg sync.WaitGroup
	var stack Stack[int]
	var v int
	wg.Add(100)
	for i := 0; i < 100; i++ {
		i := i
		go func() {
			defer wg.Done()
			stack.Push(i)
		}()
	}
	wg.Wait()
	var count int
	for {
		if stack.Pop(&v) {
			count += 1
		} else {
			break
		}
	}
}
