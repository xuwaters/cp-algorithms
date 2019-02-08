package ternary

import "testing"

func TestTernary(t *testing.T) {
	for i := -20; i < 20; i++ {
		digits := Int2BalancedTernary(i)
		str := Digits2String(digits)
		val := String2Value(str)
		t.Logf("i = %3d, str = %s, ret = %+v", i, str, digits)
		if val != i {
			t.Fatalf("ERR: val = %d, str = %s", val, str)
		}
	}
}
