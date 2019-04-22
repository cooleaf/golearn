package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFun(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("string " + s)
	if i, err := strconv.Atoi(s); err == nil {
		t.Log(10 + i)
	}
}
