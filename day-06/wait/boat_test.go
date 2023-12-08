package wait_test

import (
	"testing"

	"github.com/tgragnato/aoc23/day-06/wait"
)

func TestRace_GetWinningTimes(t *testing.T) {
	tests := []struct {
		name string
		Race wait.Race
		want uint
	}{
		{
			name: "Test 1",
			Race: wait.Race{
				Boat: wait.Boat{
					Starting: 0,
					Increase: 1,
				},
				Time:     7,
				Distance: 9,
			},
			want: 4,
		},
		{
			name: "Test 2",
			Race: wait.Race{
				Boat: wait.Boat{
					Starting: 0,
					Increase: 1,
				},
				Time:     15,
				Distance: 40,
			},
			want: 8,
		},
		{
			name: "Test 3",
			Race: wait.Race{
				Boat: wait.Boat{
					Starting: 0,
					Increase: 1,
				},
				Time:     30,
				Distance: 200,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Race.GetWinningTimes(); got != tt.want {
				t.Errorf("Race.GetWinningMultiplier() = %v, want %v", got, tt.want)
			}
		})
	}
}
