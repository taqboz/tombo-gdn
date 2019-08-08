package check

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestNotPermit(T *testing.T) {
	input := []struct{
		s string
		c []string
	}{
		{"bad",[]string{"bad","not"}},
	}

	output := []*StrIncorrect{
		{"bad",[]string{"bad"}},
	}

	for k, v := range input {
		result := NotPermit(v.s, v.c)
		switch result {
		case nil:
			if output[k] != nil {
				fmt.Println("error", k)
				os.Exit(1)
			}

		default:
			if !reflect.DeepEqual(result, output[k]) {
				fmt.Println("error", k)
				os.Exit(1)
			}
		}
	}
}

func TestNotInclude(T *testing.T) {
	input := []struct {
		s string
		c []string
	}{
		{"bad", []string{"ba"}},
	}

	output := []*NumIncorrectList{
		{"bad", []*NumIncorrect{{"ba",1}}},
	}

	for k, v := range input {
		result := NotInclude(v.s, v.c)
		switch result {
		case nil:
			if output[k] != nil {
				fmt.Println("error", k)
				os.Exit(1)
			}

		default:
			if !reflect.DeepEqual(result, output[k]) {
				fmt.Println("error", k)
				os.Exit(1)
			}
		}
	}
}
