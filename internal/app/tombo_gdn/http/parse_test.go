package http

import (
	"fmt"
	"os"
	"testing"
)

func TestResolveURL(t *testing.T) {
	type url struct {
		host string
		path string
	}

	tests := []struct {
		in url
		out string
	}{
		{url{"http://example.com/", ""}, "http://example.com/"},
		{url{"http://example.com/abc", ""}, "http://example.com/abc"},
	}

	for k, v := range tests {
		result, err := ResolveURL(v.in.host, v.in.path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if result != v.out {
			err := fmt.Errorf("Input: %v |test %d",v.in, k)
			fmt.Printf("Test.out: %s\n", v.out)
			fmt.Printf("Result: %s\n", result)
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
