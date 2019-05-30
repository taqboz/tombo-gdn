package check

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/taqboz/tombo-gdn/cli/config"
)

func TestLength(t *testing.T) {
	t.Helper()
	testConfig := []config.Tag{
		{Min:2,Max:5},
		{Min:2},
		{Max:5},
		{},
	}

	tests := []struct{
		in string
		out []*ErrCont
	} {
		{"こんにちは", []*ErrCont{{},{},{},{},}},
		{"Hello",[]*ErrCont{{},{},{},{},}},
		{"はい", []*ErrCont{{},{},{},{},}},
		{"no", []*ErrCont{{},{},{},{},}},
		{"あ", []*ErrCont{{"あ", 1},{"あ", 1},{},{},}},
		{"あかさたなはまやらわ", []*ErrCont{{"あかさたなはまやらわ", 10},{},{"あかさたなはまやらわ", 10},{},}},
		{"吾輩は猫である。名前はまだない。どこで生れたか頓と見当がつかぬ。何でも薄暗いじめじめした所でニャーニャー泣いていた事だけは記憶している。",
		[]*ErrCont{{"吾輩は猫である。名前はまだない。どこで生れたか頓と見当がつかぬ。何でも薄暗いじめじめした所でニャーニ...", 68},{},
				{"吾輩は猫である。名前はまだない。どこで生れたか頓と見当がつかぬ。何でも薄暗いじめじめした所でニャーニ...", 68},{},}},
	}

	for _, v := range tests {
		for k2, v2 := range testConfig {
			testTag := ErrTag{}
			length(&testTag, &v2, v.in)
			cont := (*v.out[k2]).Content
			num := (*v.out[k2]).Num

			if len(testTag.Length) == 0 {
				if cont != "" || num != 0 {
					t.Fatal("failed test1")
				}
			} else if r := testTag.Length[0]; r.Content != cont || r.Num != num {
				fmt.Println(r.Content, r.Num, k2)
				t.Fatal("failed test2")

			} else if testTag.IsErr {
					t.Fatal("failed test3")
			}
		}
	}

}

func TestKwCheck(t *testing.T)  {
	t.Helper()
	testConfig := []config.Tag{
		{KwMin:1,KwMax:3},
		{KwMin:1},
		{KwMax:3},
		{},
	}

	testKws := []string{"go","test"}

	tests := []struct{
		in string
		out [][]*ErrUseKw
	}{
		{"gotestgo", [][]*ErrUseKw{
			{}, {}, {}, {},
		}},

		{"goゴー", [][]*ErrUseKw{
			{{"goゴー", "test", 0}},
			{{"goゴー", "test", 0}}, {}, {},
		}},

		{"gotestgogogo", [][]*ErrUseKw{
			{{"gotestgogogo", "go", 4}}, {},
			{{"gotestgogogo", "go", 4}}, {},
		}},

		{"abc", [][]*ErrUseKw{
			{{"abc", "go", 0}, {"abc", "test", 0}},
			{{"abc", "go", 0}, {"abc", "test", 0}}, {}, {},
		}},

		{"", [][]*ErrUseKw{
			{}, {}, {}, {},
		}},
	}

	for _, v := range tests {
		for k2, v2 := range testConfig {
			testTag := ErrTag{}
			kwCheck(&testTag, &v2, v.in, testKws)

			for k3, v3 := range testTag.UseKw {
				if len(testTag.UseKw) == 0 {
					if len(v.out[k2]) != 0 {
						t.Fatal("failed test1")
					}

				} else if v3.Content != (v.out[k2][k3]).Content ||
					v3.Kw != (v.out[k2][k3]).Kw ||
					v3.Num != (v.out[k2][k3]).Num {
					t.Fatal("failed test2")

				} else if testTag.IsErr {
					t.Fatal("failed test3")
				}
			}
		}
	}
}

func TestOmission(t *testing.T) {
	t.Helper()
	tests := []struct{
		in string
		out string
	} {
		{"abc", "abc"},
		{"あいうえお","あいうえお"},
		{"吾輩は猫である。名前はまだない。どこで生れたか頓と見当がつかぬ。何でも薄暗いじめじめした所でニャーニャー泣いていた事だけは記憶している。",
			"吾輩は猫である。名前はまだない。どこで生れたか頓と見当がつかぬ。何でも薄暗いじめじめした所でニャーニ..."},
		{"Go, also known as Golang, is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.",
			"Go, also known as Golang, is a statically typed, c..."},
	}

	for _, v := range tests {
		if v.out != omission(v.in) {
			t.Fatal("failed test")
		}
	}
}

func TestMatch(t *testing.T) {
	t.Helper()

	testWords := []string{"go","テスト"}

	tests := []struct{
		in string
		out map[string]int

	} {
		{"go", map[string]int{"go":1}},
		{"テスト", map[string]int{"テスト":1}},
		{"goテスト", map[string]int{}},
		{"goユニットテスト", map[string]int{}},
		{"", map[string]int{}},
	}

	for _, v := range tests {
		testTag := ErrTag{}
		match(&testTag, testWords, v.in)
		for k2, v2 := range testTag.Match {
			if v.out[k2] != v2 {
				t.Fatal("failed test")
			}
		}

		if len(testTag.UseKw) == 0 && testTag.IsErr {
			t.Fatal("failed test")
		}
	}
}

func TestInclude(t *testing.T) {
	t.Helper()

	testWords := []string{"go","テスト"}

	tests := []struct{
		in string
		out map[string]int

	} {
		{"go", map[string]int{"go":1}},
		{"テスト", map[string]int{"テスト":1}},
		{"goテスト", map[string]int{"go":1,"テスト":1}},
		{"goユニットテスト", map[string]int{"go":1,"テスト":1}},
		{"ポケモンgoアップデート", map[string]int{"go":1}},
		{"goでテストを行う。テストはgoで書く", map[string]int{"go":2,"テスト":2}},
		{"ゴーラング", nil},
		{"", nil},
	}

	for k, v := range tests {
		testTag := ErrTag{}
		include(&testTag, testWords, v.in)
		if !reflect.DeepEqual(testTag.Include, v.out) {
			t.Fatal("failed test1", k)
		}

		if len(testTag.Include) == 0 && testTag.IsErr {
			t.Fatal("failed test2", k)
		}
	}
}

func TestMultipleContent(t *testing.T) {
	m := []*config.MultipleContent{
		{DuplicateInContent:true, Split:",", Min:2, Max:4},
		{DuplicateInContent:true, Split:",", Min:2},
		{DuplicateInContent:true, Split:",", Max:4},
		{DuplicateInContent:true, Split:","},
		{DuplicateInContent:false, Split:",", Min:2, Max:4},
	}

	tests := []struct {
		in string
		out []*NumMultiple
	}{
		{"go,golang", []*NumMultiple{{},{},{},{},nil}},
		{"go,golang,テスト,テスティング,test",
			[]*NumMultiple{{true, 5},{},{true, 5},{},nil}},
		{"go",
			[]*NumMultiple{{true, 1},{true, 1},{},{},nil}},
		{"",[]*NumMultiple{{true,0},{true,0},{},{},nil}},
	}

	for k, v := range tests {
		for k2, v2 := range m {
			testTag := ErrTag{}
			multipleContent(&testTag, v2, v.in)
			if !reflect.DeepEqual(testTag.NumMultiple, v.out[k2]) {
				t.Fatal("failed test",k,k2)
			}
		}


	}
}
