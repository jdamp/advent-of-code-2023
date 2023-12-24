package main

import (
	_ "embed"
	"reflect"
	"testing"
)

func TestRange(t *testing.T) {
	t.Run("Check range conversion", func(t *testing.T) {
		rangeTests := []struct{
			r Range;
			input int;
			want int;
		}{
			{Range{Target: 50, Source: 98, Length: 2}, 98, 50,},
			{Range{Target: 52, Source: 50, Length: 48},51, 53,},
		}
		for _, rTest := range rangeTests {
			got, _ := rTest.r.PassThrough(rTest.input)
			if got !=rTest.want {
				t.Errorf("Wrong result, got %d, want %d", got, rTest.want)
			}
		}		
	})
	t.Run("Check out-of-range error", func(t *testing.T) {
		r := Range{Target: 50, Source: 98, Length: 2}
		input := 1
		_, err := r.PassThrough(input)
		if err == nil {
			t.Errorf("Wanted error, but gone none for input: %d and range %v", input, r)
		}
	})
}


func TestRangeMapping(t *testing.T) {
	t.Run("Create a RangeMapping from string", func(t *testing.T) {
		got := MakeRangeMapping("0 15 37\n37 52 2\n39 0 15")
		want := RangeMapping{
			[]Range{
				{0, 15, 37},
				{37, 52, 2},
				{39, 0, 15},
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Error("Error constructring range from string")
		}
	})
	t.Run("Pass through", func(t *testing.T) {
		mapping := MakeRangeMapping("50 98 2\n52 50 48")
		testCases := []struct {
			input int
			want int
		} {
			{input: 0, want: 0,},
			{input: 1, want: 1,},
			{input: 49, want: 49,},
			{input: 50, want: 52,},
			{input: 51, want: 53,},
			{input: 97, want: 99,},
			{input: 98, want: 50,},
		}
		for _, testCase := range testCases {
			got := mapping.PassThrough(testCase.input)
			if got != testCase.want {
				t.Errorf("Error for PassThrough: got %d, want %d for input %d", got, testCase.want, testCase.input)
			}
		}
	})
}

//go:embed test_input.txt
var input string
func TestSolve(t *testing.T) {
	almanac := ParseInput(input)
	want := 35
	got := Solve(almanac)
	if want != got {
		t.Errorf("Got %d, want %d", got, want)
	}

}