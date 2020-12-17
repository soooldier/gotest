package stack

import (
	"reflect"
	"sync"
	"testing"
)

func TestStack(t *testing.T) {
	var cases = []struct {
		input    []interface{}
		expected []interface{}
	}{
		{[]interface{}{1, 2, 3, 4, 5, 6}, []interface{}{6, 5, 4, 3, 2, 1}},
		{[]interface{}{100, 100000}, []interface{}{100000, 100}},
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{"a", "b", "c"}, []interface{}{"c", "b", "a"}},
	}
	for _, c := range cases {
		s1 := new(Stack)
		for _, v := range c.input {
			s1.Push(v)
		}
		for _, v := range c.expected {
			e := s1.Pop()
			if reflect.ValueOf(e).Type() != reflect.ValueOf(v).Type() || e != v {
				t.Errorf("not match %#v, except %d", e, v)
			}
		}
	}
}

func TestStackGoroutine(t *testing.T) {
	s := new(Stack)
	var wg1 sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg1.Add(1)
		go func(i int) {
			defer wg1.Done()
			s.Push(i)
		}(i)
	}
	wg1.Wait()
	var wg2 sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			item := s.Pop()
			if item == nil {
				t.Error("push error")
			}
		}()
	}
	wg2.Wait()
	if s.Pop() != nil {
		t.Error("pop error")
	}
}

func BenchmarkStack(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	s := new(Stack)
	var wg1 sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg1.Add(1)
		go func(i int) {
			defer wg1.Done()
			s.Push(i)
			s.Pop()
		}(i)
	}
	wg1.Wait()
}
