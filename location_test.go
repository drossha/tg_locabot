package main

import "testing"

func TestFindCoordinates(t *testing.T) {
	type args struct {
		str string
	}
	var tests = []struct {
		name    string
		text    string
		wantLat string
		wantLon string
	}{
		struct {
			name    string
			text    string
			wantLat string
			wantLon string
		}{
			name:    "Only geo",
			text:    "11.111 22.222",
			wantLat: "11.111",
			wantLon: "22.222",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLat, gotLon := FindCoordinates(tt.text)
			if gotLat != tt.wantLat {
				t.Errorf("FindCoordinates() gotLat = %v, want %v", gotLat, tt.wantLat)
			}
			if gotLon != tt.wantLon {
				t.Errorf("FindCoordinates() gotLon = %v, want %v", gotLon, tt.wantLon)
			}
		})
	}
}
