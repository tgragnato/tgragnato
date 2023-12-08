package seed_test

import (
	"reflect"
	"testing"

	"github.com/tgragnato/aoc23/day-05/seed"
)

func TestMap_GetDestination(t *testing.T) {
	type fields struct {
		Destination uint
		Source      uint
		Range       uint
	}
	tests := []struct {
		name    string
		fields  fields
		source  uint
		want    uint
		wantErr bool
	}{
		{
			name: "Case 0",
			fields: fields{
				Destination: 52,
				Source:      50,
				Range:       48,
			},
			source:  0,
			want:    0,
			wantErr: true,
		},
		{
			name: "Case 49",
			fields: fields{
				Destination: 52,
				Source:      50,
				Range:       48,
			},
			source:  49,
			want:    0,
			wantErr: true,
		},
		{
			name: "Case 50",
			fields: fields{
				Destination: 52,
				Source:      50,
				Range:       48,
			},
			source:  50,
			want:    52,
			wantErr: false,
		},
		{
			name: "Case 97",
			fields: fields{
				Destination: 52,
				Source:      50,
				Range:       48,
			},
			source:  97,
			want:    99,
			wantErr: false,
		},
		{
			name: "Case 98",
			fields: fields{
				Destination: 52,
				Source:      50,
				Range:       48,
			},
			source:  98,
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &seed.Map{
				Destination: tt.fields.Destination,
				Source:      tt.fields.Source,
				Range:       tt.fields.Range,
			}
			got, err := m.GetDestination(tt.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Map.GetDestination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Map.GetDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}

var testSet = seed.MapSet{
	[]seed.Map{
		{
			Destination: 50,
			Source:      98,
			Range:       2,
		},
		{
			Destination: 52,
			Source:      50,
			Range:       48,
		},
	},
}

func TestMapSet_GetDestination(t *testing.T) {
	tests := []struct {
		name   string
		Set    seed.MapSet
		source uint
		want   uint
	}{
		{
			name:   "Case 0",
			Set:    testSet,
			source: 0,
			want:   0,
		},
		{
			name:   "Case 49",
			Set:    testSet,
			source: 49,
			want:   49,
		},
		{
			name:   "Case 50",
			Set:    testSet,
			source: 50,
			want:   52,
		},
		{
			name:   "Case 97",
			Set:    testSet,
			source: 97,
			want:   99,
		},
		{
			name:   "Case 98",
			Set:    testSet,
			source: 98,
			want:   50,
		},
		{
			name:   "Case 99",
			Set:    testSet,
			source: 99,
			want:   51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Set.GetDestination(tt.source); got != tt.want {
				t.Errorf("MapSet.GetDestination() = %v, %v => %v", got, tt.source, tt.want)
			}
		})
	}
}

func TestParseSeeds(t *testing.T) {
	t.Run("seeds: 79 14 55 13", func(t *testing.T) {
		if got := seed.ParseSeeds("seeds: 79 14 55 13"); !reflect.DeepEqual(got, []uint{79, 14, 55, 13}) {
			t.Errorf("ParseSeeds(seeds: 79 14 55 13) = %v", got)
		}
	})
}

func TestRecursiveMapSet_GetLowestLocation(t *testing.T) {
	t.Run("Simplest example given", func(t *testing.T) {
		r := &seed.RecursiveMapSet{
			RecursiveSet: []seed.MapSet{testSet},
			Seeds:        map[uint]uint{79: 0, 14: 0, 55: 0, 13: 0},
		}
		if got := r.GetLowestLocation(); got != 13 {
			t.Errorf("RecursiveMapSet.GetDestination() = %v, want 13", got)
		}
	})
}
