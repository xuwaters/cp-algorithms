package lyndon

import "testing"

func TestLyndonFactorization(t *testing.T) {
	dataList := []struct {
		ans []string
		s   string
	}{
		{
			ans: []string{"adc", "abd", "abc", "abacdaehello"},
			s:   "adcabdabcabacdaehello",
		},
	}

	for _, data := range dataList {
		got := LyndonFactorization(data.s)
		t.Logf("got = %+v, data = %+v", got, data)
	}
}
