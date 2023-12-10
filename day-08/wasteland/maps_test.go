package wasteland_test

import (
	"testing"

	"github.com/tgragnato/aoc23/day-08/wasteland"
)

func TestMaps_ReachZZZ(t *testing.T) {
	type fields struct {
		Left         map[string]string
		Right        map[string]string
		Instructions []bool
	}
	tests := []struct {
		name      string
		fields    fields
		wantSteps uint
	}{
		{
			name: "Example 1",
			fields: fields{
				Left: map[string]string{
					"AAA": "BBB",
					"BBB": "DDD",
					"CCC": "ZZZ",
					"DDD": "DDD",
					"EEE": "EEE",
					"GGG": "GGG",
					"ZZZ": "ZZZ",
				},
				Right: map[string]string{
					"AAA": "CCC",
					"BBB": "EEE",
					"CCC": "GGG",
					"DDD": "DDD",
					"EEE": "EEE",
					"GGG": "GGG",
					"ZZZ": "ZZZ",
				},
				Instructions: []bool{true, false},
			},
			wantSteps: 2,
		},
		{
			name: "Example 2",
			fields: fields{
				Left: map[string]string{
					"AAA": "BBB",
					"BBB": "AAA",
					"ZZZ": "ZZZ",
				},
				Right: map[string]string{
					"AAA": "BBB",
					"BBB": "ZZZ",
					"ZZZ": "ZZZ",
				},
				Instructions: []bool{false, false, true},
			},
			wantSteps: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &wasteland.Maps{
				Left:         tt.fields.Left,
				Right:        tt.fields.Right,
				Instructions: tt.fields.Instructions,
			}
			if gotSteps := m.ReachZZZ(); gotSteps != tt.wantSteps {
				t.Errorf("Maps.ReachZZZ() = %v, want %v", gotSteps, tt.wantSteps)
			}
		})
	}
}

func TestMaps_ReachXXZ(t *testing.T) {
	t.Run("Example 3", func(t *testing.T) {
		m := &wasteland.Maps{
			Left: map[string]string{
				"11A": "11B",
				"11B": "XXX",
				"11Z": "11B",
				"22A": "22B",
				"22B": "22C",
				"22C": "22Z",
				"22Z": "22B",
				"XXX": "XXX",
			},
			Right: map[string]string{
				"11A": "XXX",
				"11B": "11Z",
				"11Z": "XXX",
				"22A": "XXX",
				"22B": "22C",
				"22C": "22Z",
				"22Z": "22B",
				"XXX": "XXX",
			},
			Instructions: []bool{false, true},
		}
		if gotSteps := m.ReachXXZ(); gotSteps != 6 {
			t.Errorf("Maps.ReachXXZ() = %v, want 6", gotSteps)
		}
	})
}
