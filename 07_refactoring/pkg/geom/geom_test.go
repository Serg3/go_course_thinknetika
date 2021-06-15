package geom

import "testing"

func TestGeom_Distance(t *testing.T) {
	tests := []struct {
		name         string
		geom         Coords
		wantDistance float64
	}{
		{
			name:         "#1",
			geom:         Coords{X1: 1, Y1: 1, X2: 4, Y2: 5},
			wantDistance: 5,
		},
		// * added edge case test
		{
			name:         "#2",
			geom:         Coords{X1: 0, Y1: 0, X2: 0, Y2: 0},
			wantDistance: 0,
		},
		// * added edge case test
		{
			name:         "#3",
			geom:         Coords{X1: -1, Y1: -1, X2: -4, Y2: -5},
			wantDistance: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := tt.geom.Distance(); gotDistance != tt.wantDistance {
				t.Errorf("Coords.Distance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
