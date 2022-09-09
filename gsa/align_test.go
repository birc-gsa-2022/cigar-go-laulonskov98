package gsa

import (
	"fmt"
	"testing"
)

func ExampleAlign() {
	pRow, qRow := Align("ACCACAGTCATA", "ACAGAGTACAAA", "MDMMMMMMIMMMM")
	fmt.Println(pRow, qRow)
	// Output: ACCACAGT-CATA A-CAGAGTACAAA
}

func ExampleLocalAlign() {
	pRow, qRow := LocalAlign("ACCACAGTCATA", "CGACAGAGTACAAA", 2, "MDMMMMMMIMMMM")
	fmt.Println(pRow, qRow)
	// Output: ACCACAGT-CATA A-CAGAGTACAAA
}

func TestAlign(t *testing.T) {
	type args struct {
		p     string
		q     string
		edits string
	}
	tests := []struct {
		name     string
		args     args
		wantPRow string
		wantQRow string
	}{
		// Add more test cases here as needed
		{"Basic case",
			args{"ACCACAGTCATA", "ACAGAGTACAAA", "MDMMMMMMIMMMM"},
			"ACCACAGT-CATA", "A-CAGAGTACAAA",
		},
		{"Empty q",
			args{"a", "", "D"},
			"a", "-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPRow, gotQRow := Align(tt.args.p, tt.args.q, tt.args.edits)
			if gotPRow != tt.wantPRow {
				t.Errorf("Align() gotPRow = %v, want %v", gotPRow, tt.wantPRow)
			}
			if gotQRow != tt.wantQRow {
				t.Errorf("Align() gotQRow = %v, want %v", gotQRow, tt.wantQRow)
			}
		})
	}
}

func TestLocalAlign(t *testing.T) {
	type args struct {
		p     string
		x     string
		i     int
		edits string
	}
	tests := []struct {
		name     string
		args     args
		wantPRow string
		wantXRow string
	}{
		// Add more test cases here as needed
		{"Basic case",
			args{"ACCACAGTCATA", "CTACAGAGTACAAA", 2, "MDMMMMMMIMMMM"},
			"ACCACAGT-CATA", "A-CAGAGTACAAA",
		},
		{"Empty q",
			args{"a", "CTA", 2, "D"},
			"a", "-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPRow, gotXRow := LocalAlign(tt.args.p, tt.args.x, tt.args.i, tt.args.edits)
			if gotPRow != tt.wantPRow {
				t.Errorf("LocalAlign() gotPRow = %v, want %v", gotPRow, tt.wantPRow)
			}
			if gotXRow != tt.wantXRow {
				t.Errorf("LocalAlign() gotXRow = %v, want %v", gotXRow, tt.wantXRow)
			}
		})
	}
}
