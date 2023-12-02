package configuration_test

import (
	"testing"

	"github.com/tgragnato/aoc23/day-02/configuration"
)

func TestConfiguration_Evaluate(t *testing.T) {
	type fields struct {
		Red   uint
		Blue  uint
		Green uint
	}
	tests := []struct {
		name   string
		fields fields
		line   string
		want   bool
	}{
		{
			name:   "Game 1 (low)",
			fields: fields{Red: 12, Blue: 13, Green: 14},
			line:   "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want:   true,
		},
		{
			name:   "Game 2 (low)",
			fields: fields{Red: 12, Blue: 13, Green: 14},
			line:   "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want:   true,
		},
		{
			name:   "Game 3 (low)",
			fields: fields{Red: 12, Blue: 13, Green: 14},
			line:   "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want:   false,
		},
		{
			name:   "Game 4 (low)",
			fields: fields{Red: 12, Blue: 13, Green: 14},
			line:   "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want:   false,
		},
		{
			name:   "Game 5 (low)",
			fields: fields{Red: 12, Blue: 13, Green: 14},
			line:   "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want:   true,
		},
		{
			name:   "Game 1 (high)",
			fields: fields{Red: 100, Blue: 100, Green: 100},
			line:   "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want:   true,
		},
		{
			name:   "Game 2 (high)",
			fields: fields{Red: 100, Blue: 100, Green: 100},
			line:   "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want:   true,
		},
		{
			name:   "Game 3 (high)",
			fields: fields{Red: 100, Blue: 100, Green: 100},
			line:   "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want:   true,
		},
		{
			name:   "Game 4 (high)",
			fields: fields{Red: 100, Blue: 100, Green: 100},
			line:   "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want:   true,
		},
		{
			name:   "Game 5 (high)",
			fields: fields{Red: 100, Blue: 100, Green: 100},
			line:   "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration.Configuration{
				Red:   tt.fields.Red,
				Blue:  tt.fields.Blue,
				Green: tt.fields.Green,
			}
			if got := c.Evaluate(tt.line); got != tt.want {
				t.Errorf("Configuration.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfiguration_Power(t *testing.T) {
	tests := []struct {
		name string
		line string
		want uint
	}{
		{
			name: "Game 1",
			line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: 48,
		},
		{
			name: "Game 2",
			line: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want: 12,
		},
		{
			name: "Game 3",
			line: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want: 1560,
		},
		{
			name: "Game 4",
			line: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want: 630,
		},
		{
			name: "Game 5",
			line: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want: 36,
		},
	}
	for _, tt := range tests {
		c := &configuration.Configuration{}
		t.Run(tt.name, func(t *testing.T) {
			if got := c.Power(tt.line); got != tt.want {
				t.Errorf("Configuration.Power() = %v, want %v", got, tt.want)
			}
		})
	}
}
