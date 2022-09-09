package gsa

import (
	"fmt"
	"testing"
)

func ExampleCigarToEdits() {
	edits := CigarToEdits("1M1D6M1I4M")
	fmt.Println(edits)
	// Output: MDMMMMMMIMMMM
}

func ExampleEditsToCigar() {
	cigar := EditsToCigar("MDMMMMMMIMMMM")
	fmt.Println(cigar)
	// Output: 1M1D6M1I4M
}

func TestCigarToEdits(t *testing.T) {
	type args struct {
		cigar string
	}
	tests := []struct {
		name      string
		args      args
		wantEdits string
	}{
		// TODO: Add test cases.
		{
			"Test 1",
			args{"1M1D1I1M1I1D"},
			"MDIMID",
		},
		{
			"Test 2",
			args{"2M2D2I2M2I2D"},
			"MMDDIIMMIIDD",
		},
		{
			"Test 3",
			args{"1M2D3I2M1I2D"},
			"MDDIIIMMIDD",
		},
		{
			"Test 4",
			args{""},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEdits := CigarToEdits(tt.args.cigar); gotEdits != tt.wantEdits {
				t.Errorf("CigarToEdits() = %v, want %v", gotEdits, tt.wantEdits)
			}
		})
	}
}

func TestEditsToCigar(t *testing.T) {
	type args struct {
		edits string
	}
	tests := []struct {
		name      string
		args      args
		wantCigar string
	}{
		// TODO: Add test cases.
		{
			"Test 1",
			args{"MDIMID"},
			"1M1D1I1M1I1D",
		},
		{
			"Test 2",
			args{"MMDDIIMMIIDD"},
			"2M2D2I2M2I2D",
		},
		{
			"Test 3",
			args{"MDDIIIMMIDD"},
			"1M2D3I2M1I2D",
		},
		{
			"Test 4",
			args{""},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCigar := EditsToCigar(tt.args.edits); gotCigar != tt.wantCigar {
				t.Errorf("EditsToCigar() = %v, want %v", gotCigar, tt.wantCigar)
			}
		})
	}
}
