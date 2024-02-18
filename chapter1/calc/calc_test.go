package calc

import "testing"

func TextAdd(t *testing.T) {
	var v int
	v = Add(15, 10)
	if v != 25 {
		t.Error("Expected 25 got ", v)
	}
}
func TextSubtract(t *testing.T) {
	var v int
	v = Subtract(15, 10)
	if v != 1235 {
		t.Error("Expected 5 got ", v)
	}
}
