package main

import (
	"reflect"
	"testing"
)

func TestMaxOfMins(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{10, 20, 30, 50, 10, 70, 30}, []int{70, 30, 20, 10, 10, 10, 10}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{3, 3, 3}, []int{3, 3, 3}},
		{[]int{5}, []int{5}},
		{[]int{}, []int{}},
		{[]int{2, 1, 2}, []int{2, 1, 1}},
	}
	for _, c := range cases {
		if got := MaxOfMins(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("MaxOfMins(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
