package extraction_test

import (
	"testing"

	"github.com/tgragnato/aoc23/day-01/extraction"
)

func TestExtractValues1(t *testing.T) {
	tests := []struct {
		name string
		line string
		want uint
	}{
		{
			name: "Line 1",
			line: "1abc2",
			want: 12,
		},
		{
			name: "Line 2",
			line: "pqr3stu8vwx",
			want: 38,
		},
		{
			name: "Line 3",
			line: "a1b2c3d4e5f",
			want: 15,
		},
		{
			name: "Line 4",
			line: "treb7uchet",
			want: 77,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extraction.ExtractValues1(tt.line); got != tt.want {
				t.Errorf("ExtractValues1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractValues2(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		line string
		want uint
	}{
		{
			name: "Line 1",
			line: "two1nine",
			want: 29,
		},
		{
			name: "Line 2",
			line: "eightwothree",
			want: 83,
		},
		{
			name: "Line 3",
			line: "abcone2threexyz",
			want: 13,
		},
		{
			name: "Line 4",
			line: "xtwone3four",
			want: 24,
		},
		{
			name: "Line 5",
			line: "4nineeightseven2",
			want: 42,
		},
		{
			name: "Line 6",
			line: "zoneight234",
			want: 14,
		},
		{
			name: "Line 7",
			line: "7pqrstsixteen",
			want: 76,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extraction.ExtractValues2(tt.line); got != tt.want {
				t.Errorf("ExtractValues2() = %v, want %v", got, tt.want)
			}
		})
	}
}
