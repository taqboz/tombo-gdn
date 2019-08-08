package check

import (
	"fmt"
	"github.com/taqboz/tombo_gdn/cli/config"
	"os"
	"reflect"
	"testing"
)

func TestCommon(T *testing.T) {
	input := []struct {
			s string
			i int
			c config.CheckLength
	}{
		{"あいうえお",5,config.CheckLength{"all", 2,6}},
		{"かきくけこ",5,config.CheckLength{"all", 4,5}},
		{"さしすせそ",5,config.CheckLength{"all", 5,6}},
		{"たちつ",3,config.CheckLength{"all", 5,6}},
		{"てと",2,config.CheckLength{"all", 0,1}},
		{"なにぬねの",3,config.CheckLength{"all", 5,7}},
	}

	output := []*NumIncorrect{
		nil,
		nil,
		nil,
		{"たちつ",3},
		{"てと", 2},
		{"なにぬねの", 3},
	}

	for k, v := range input {
		result := common(v.s, v.i, v.c)
		switch result {
		case nil:
			if output[k] != nil {
				fmt.Println("error", k)
				os.Exit(1)
			}
		default:
			if !reflect.DeepEqual(result, output[k]) {
				fmt.Print("r:", *result, "o:", *output[k])
				fmt.Println("error", k)
				os.Exit(1)
			}
		}
	}
}

func TestLength(t *testing.T) {
	input := []struct{
			s string
			c config.CheckLength
	}{
		{"abc",config.CheckLength{"all",0,2}},
		{"あいうえお",config.CheckLength{"all",0,6}},
	}

	output := []*NumIncorrect{
		{"abc", 3},
		nil,
	}

	for k, v := range input {
		result := Length(v.s, v.c)
		switch result {
		case nil:
			if output[k] != nil {
				fmt.Println("error", k)
				os.Exit(1)
			}
		default:
			if !reflect.DeepEqual(result, output[k]) {
				fmt.Print("r:", *result, "o:", *output[k])
				fmt.Println("error", k)
				os.Exit(1)
			}
		}
	}
}

func TestUseKws(t *testing.T) {
	kws := []string{"go", "a"}

	input := []struct{
		s string
		c config.CheckLength
	}{
		{"abc",config.CheckLength{"all",2,4}},
		{"goa",config.CheckLength{"all",1,2}},
	}

	output := []*NumIncorrectList{
		{"abc", []*NumIncorrect{{"go",0},{"a", 1}}},
		{"goa",[]*NumIncorrect{}},
	}

	for k, v := range input {
		result := UseKws(v.s, v.c, kws)
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

func TestMultipleCount(t *testing.T) {
	input := []struct{
		s string
		c config.CheckLength
	}{
		{"a,b,c",config.CheckLength{"all",0,2}},
		{"あい,うえお",config.CheckLength{"all",0,6}},
	}

	output := []*NumIncorrect{
		{"a,b,c", 3},
		nil,
	}

	for k, v := range input {
		result := MultipleCount(v.s, v.c,",")
		switch result {
		case nil:
			if output[k] != nil {
				fmt.Println("error", k)
				os.Exit(1)
			}
		default:
			if !reflect.DeepEqual(result, output[k]) {
				fmt.Print("r:", *result, "o:", *output[k])
				fmt.Println("error", k)
				os.Exit(1)
			}
		}
	}
}
