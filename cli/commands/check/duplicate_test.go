package check

import (
	"fmt"
	"github.com/taqboz/tombo_gdn/cli/config"
	"os"
	"reflect"
	"testing"
)

func TestMultipleDuplicate(t *testing.T) {
	input := []struct {
		t string
		c config.CheckLength
		split string
	}{
		{"a,b,a,a",config.CheckLength{"all",0,2}, ","},
	}

	output := []*NumIncorrectList{
		{"a,b,a,a",[]*NumIncorrect{{"a", 3}}},
	}

	for k, v := range input {
		result := MultipleDuplicate(v.t, v.c, v.split)
		switch result {
		case nil:
			if output[k] != nil {
				fmt.Println("error", k)
				os.Exit(1)
			}

		default:
			if !reflect.DeepEqual(result, output[k]) {
				fmt.Println(result.Incorrect)
				fmt.Println("error", k)
				os.Exit(1)
			}
		}
	}

}

func TestDuplicateNum(t *testing.T) {
	input := []struct {
		t string
		s []string
	}{
		{"a",[]string{"a","b","a"}},
	}

	output := []int{
		2,
	}

	for k, v := range input {
		result := DuplicateNum(v.t, v.s)
		if result != output[k] {
			fmt.Println("Test Failed")
			os.Exit(1)
		}
	}
}
