package map_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m1 := map[int]func(op int) int{}
	m1[1] = func(op int) int { return op }
	m1[2] = func(op int) int { return op * op }
	m1[3] = func(op int) int { return op ^ 3 }

	a := m1[2](3)
	t.Log(a)
}
