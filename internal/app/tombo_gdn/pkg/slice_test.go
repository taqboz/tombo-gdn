package pkg

import (
	"fmt"
	"strconv"
	"testing"
)

func TestIsContain(t *testing.T) {
	tests := []struct{
		in []string
		in2 string
		out bool
	}{
		// http(s)のURLのみ抽出できるか確認
		{[]string{"a","b","c"}, "a", true},

		{[]string{"a","b","c"}, "e", false},
		{[]string{"a","b","c"},"ｂ", false},
		{[]string{}, "a", false},
	}

	for i, test := range tests {
		result := IsContain(test.in, test.in2)

		if result != test.out {
			fmt.Println(result)
			t.Fatal("failed test ", strconv.Itoa(i))
		}
	}
}
