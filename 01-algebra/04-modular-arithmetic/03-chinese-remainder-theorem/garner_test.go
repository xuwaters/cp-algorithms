package chinese_remainder

import "testing"

func TestGarnerBigInt(t *testing.T) {
	dataList := []struct {
		a int
		b int
	}{
		{a: 38102, b: 8203},
		{a: 82331, b: 8213},
		{a: 893231, b: 71342},
		{a: 893231, b: 1},
		{a: 0, b: 81033},
	}

	for _, data := range dataList {
		a := NewBigInt(data.a)
		b := NewBigInt(data.b)
		ret := a.Multiply(b).ToValue()
		expect := data.a * data.b
		if ret == expect {
			t.Logf(" OK: %d = %d * %d", ret, data.a, data.b)
		} else {
			t.Fatalf("ERR: %d != %d = %d * %d", ret, expect, data.a, data.b)
		}
	}
}
