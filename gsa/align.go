// You can create modules at this level and they will be
// interpreted as under module birc.au.dk, so to import
// package `gsa` you need `import "birc.au.dk/gsa"`

package gsa

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
	pRow, qRow = "", ""
	// Align p and q based on edits
	return pRow, qRow
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
	return pRow, xRow
}
