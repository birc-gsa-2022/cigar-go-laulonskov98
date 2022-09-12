// You can create modules at this level and they will be
// interpreted as under module birc.au.dk, so to import
// package `gsa` you need `import "birc.au.dk/gsa"`

package gsa

import (
	"strings"
)

// Align two sequences from a sequence of edits.
//
//  Args:
//      p: The first sequence to align.
//      q: The second sequence to align
//      edits: The list of edits to apply, given as a string
//
//  Returns:
//      The two rows in the pairwise alignment
func Align(p, q, edits string) (pRow, qRow string) {
	var pBuffer strings.Builder
	var qBuffer strings.Builder

	pItr, qItr := 0, 0
	for _, char := range edits {

		pEnd := pItr + 1
		qEnd := qItr + 1
		if char == 'M' {
			pBuffer.WriteString(p[pItr:pEnd])
			qBuffer.WriteString(q[qItr:qEnd])
			pItr += 1
			qItr += 1
		}

		if char == 'I' {
			pBuffer.WriteString("-")
			qBuffer.WriteString(q[qItr:qEnd])
			qItr += 1
		}

		if char == 'D' {
			pBuffer.WriteString(p[pItr:pEnd])
			qBuffer.WriteString("-")
			pItr += 1
		}
	}
	// Align p and q based on edits
	return pBuffer.String(), qBuffer.String()
}

// Align two sequences from a sequence of edits.
//
//  Args:
//      p: The first sequence to align
//      x: The second sequence to align; we only align locally
//      i: Start position of the alignment in x
//      edits: The list of edits to apply, given as a string
//
//  Returns:
//      The two rows in the pairwise alignment

func LocalAlign(p, x string, i int, edits string) (pRow, xRow string) {
	pRow, xRow = "", ""
	// Align p and q based on edits
	pRow, xRow = Align(p, x[i:], edits)
	return pRow, xRow
}
