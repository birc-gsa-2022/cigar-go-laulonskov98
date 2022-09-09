// You can create modules at this level and they will be
// interpreted as under module birc.au.dk, so to import
// package `gsa` you need `import "birc.au.dk/gsa"`

package gsa

// Extract the edit operations from a pairwise alignment.
//
//  Args:
//      p: The first row in the pairwise alignment.
//      q: The second row in the pairwise alignment.
//
//  Returns:
//      The two strings without gaps and the list of edit operations
//      as a string.
func GetEdits(p, q string) (gapFreeP, gapFreeQ, edits string) {
	gapFreeP, gapFreeQ, edits = "", "", ""
	return gapFreeP, gapFreeQ, edits
}

//  Get the distance between p and the string that starts at x[i:]
//  using the edits.
//
//  Args:
//      p: The read string we have mapped against x
//      x: The longer string we have mapped against
//      i: The location where we have an approximative match
//      edits: The list of edits to apply, given as a string
//
//  Returns:
//      The distance from p to x[i:?] described by edits
func EditDist(p, x string, i int, edits string) int {
	return -1
}
