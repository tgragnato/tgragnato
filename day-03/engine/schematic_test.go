package engine_test

import (
	"testing"

	"github.com/tgragnato/aoc23/day-03/engine"
)

var testSchema = [][]rune{
	{'4', '6', '7', '.', '.', '1', '1', '4', '.', '.'},
	{'.', '.', '.', '*', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '3', '5', '.', '.', '6', '3', '3', '.'},
	{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
	{'6', '1', '7', '*', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '+', '.', '5', '8', '.'},
	{'.', '.', '5', '9', '2', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '7', '5', '5', '.'},
	{'.', '.', '.', '$', '.', '*', '.', '.', '.', '.'},
	{'.', '6', '6', '4', '.', '5', '9', '8', '.', '.'},
}

func TestSchematic_IsAdjacent(t *testing.T) {
	tests := []struct {
		name   string
		Schema [][]rune
		start  int
		end    int
		line   int
		want   bool
	}{
		{
			name:   "467",
			Schema: testSchema,
			start:  0,
			end:    2,
			line:   0,
			want:   true,
		},
		{
			name:   "114",
			Schema: testSchema,
			start:  5,
			end:    7,
			line:   0,
			want:   false,
		},
		{
			name:   "35",
			Schema: testSchema,
			start:  2,
			end:    3,
			line:   2,
			want:   true,
		},
		{
			name:   "633",
			Schema: testSchema,
			start:  6,
			end:    8,
			line:   2,
			want:   true,
		},
		{
			name:   "617",
			Schema: testSchema,
			start:  0,
			end:    2,
			line:   4,
			want:   true,
		},
		{
			name:   "58",
			Schema: testSchema,
			start:  7,
			end:    8,
			line:   5,
			want:   false,
		},
		{
			name:   "592",
			Schema: testSchema,
			start:  2,
			end:    4,
			line:   6,
			want:   true,
		},
		{
			name:   "755",
			Schema: testSchema,
			start:  6,
			end:    8,
			line:   7,
			want:   true,
		},
		{
			name:   "664",
			Schema: testSchema,
			start:  1,
			end:    3,
			line:   9,
			want:   true,
		},
		{
			name:   "598",
			Schema: testSchema,
			start:  5,
			end:    7,
			line:   9,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &engine.Schematic{
				Schema: tt.Schema,
			}
			if got := s.IsAdjacent(engine.PartNumber{
				Start: tt.start,
				End:   tt.end,
				Row:   tt.line,
			}); got != tt.want {
				t.Errorf("Schematic.isAdjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchematic_ExtractPartNumber(t *testing.T) {
	tests := []struct {
		name   string
		Schema [][]rune
		start  int
		end    int
		row    int
		want   uint
	}{
		{
			name:   "467",
			Schema: testSchema,
			start:  0,
			end:    2,
			row:    0,
			want:   467,
		},
		{
			name:   "114",
			Schema: testSchema,
			start:  5,
			end:    7,
			row:    0,
			want:   114,
		},
		{
			name:   "35",
			Schema: testSchema,
			start:  2,
			end:    3,
			row:    2,
			want:   35,
		},
		{
			name:   "633",
			Schema: testSchema,
			start:  6,
			end:    8,
			row:    2,
			want:   633,
		},
		{
			name:   "617",
			Schema: testSchema,
			start:  0,
			end:    2,
			row:    4,
			want:   617,
		},
		{
			name:   "58",
			Schema: testSchema,
			start:  7,
			end:    8,
			row:    5,
			want:   58,
		},
		{
			name:   "592",
			Schema: testSchema,
			start:  2,
			end:    4,
			row:    6,
			want:   592,
		},
		{
			name:   "755",
			Schema: testSchema,
			start:  6,
			end:    8,
			row:    7,
			want:   755,
		},
		{
			name:   "664",
			Schema: testSchema,
			start:  1,
			end:    3,
			row:    9,
			want:   664,
		},
		{
			name:   "598",
			Schema: testSchema,
			start:  5,
			end:    7,
			row:    9,
			want:   598,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &engine.Schematic{
				Schema: tt.Schema,
			}
			if got := s.ExtractPartNumber(engine.PartNumber{
				Start: tt.start,
				End:   tt.end,
				Row:   tt.row,
			}); got != tt.want {
				t.Errorf("Schematic.ExtractPartNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchematic_SumPartNumber(t *testing.T) {
	t.Run("part1-test", func(t *testing.T) {
		s := &engine.Schematic{
			Schema: testSchema,
		}
		if gotSum := s.SumPartNumber(); gotSum != 4361 {
			t.Errorf("Schematic.SumPartNumber() = %v, want 4361", gotSum)
		}
	})

}

func TestSchematic_SumGearRatio(t *testing.T) {
	t.Run("part2-test", func(t *testing.T) {
		s := &engine.Schematic{
			Schema: testSchema,
		}
		if got := s.SumGearRatio(); got != 467835 {
			t.Errorf("Schematic.SumGearRatio() = %v, want 467835", got)
		}
	})
}
