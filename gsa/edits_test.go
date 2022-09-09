package gsa

import (
	"fmt"
	"testing"
)

func ExampleGetEdits() {
	p, q, edits := GetEdits("ACCACAGT-CATA", "A-CAGAGTACAAA")
	fmt.Println(p, q, edits)
	// Output: ACCACAGTCATA ACAGAGTACAAA MDMMMMMMIMMMM
}

func ExampleEditDist() {
	fmt.Println(EditDist("accaaagta", "cgacaaatgtcca", 2, "MDMMIMMMMIIM"))
	// Output: 5
}

func TestGetEdits(t *testing.T) {
	type args struct {
		p string
		q string
	}
	tests := []struct {
		name         string
		args         args
		wantGapFreeP string
		wantGapFreeQ string
		wantEdits    string
	}{
		// TODO: Add test cases as needed.
		{
			"Test 1",
			args{"acca-aagt--a", "a-caaatgtcca"},
			"accaaagta", "acaaatgtcca", "MDMMIMMMMIIM",
		},
		{
			"Test 2",
			args{"acgttcga", "aaa---aa"},
			"acgttcga", "aaaaa", "MMMDDDMM",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGapFreeP, gotGapFreeQ, gotEdits := GetEdits(tt.args.p, tt.args.q)
			if gotGapFreeP != tt.wantGapFreeP {
				t.Errorf("GetEdits() gotGapFreeP = %v, want %v", gotGapFreeP, tt.wantGapFreeP)
			}
			if gotGapFreeQ != tt.wantGapFreeQ {
				t.Errorf("GetEdits() gotGapFreeQ = %v, want %v", gotGapFreeQ, tt.wantGapFreeQ)
			}
			if gotEdits != tt.wantEdits {
				t.Errorf("GetEdits() gotEdits = %v, want %v", gotEdits, tt.wantEdits)
			}
		})
	}

}

func TestEditDist(t *testing.T) {
	type args struct {
		p     string
		x     string
		i     int
		edits string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"Test 1",
			args{"accaaagta", "cgacaaatgtcca", 2, "MDMMIMMMMIIM"},
			5,
		},
		{
			"Test 2",
			args{"acgttcga", "aaaaa", 0, "MMMDDDMM"},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EditDist(tt.args.p, tt.args.x, tt.args.i, tt.args.edits); got != tt.want {
				t.Errorf("EditDist() = %v, want %v", got, tt.want)
			}
		})
	}
}
