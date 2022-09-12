// You can create modules at this level and they will be
// interpreted as under module birc.au.dk, so to import
// package `gsa` you need `import "birc.au.dk/gsa"`

package gsa

import "strings"

// Extract the edit operations from a pairwise alignment.
//
//  Args:
//      p: The first row in the pairwise alignment.
//      q: The second row in the pairwise alignment.
//
//  Returns:
//      The two strings without gaps and the list of edit operations
//      as a string.
func GetEdits(p, q string) (string, string, string) {
	var gapFreePBuf strings.Builder
	var gapFreeQBuf strings.Builder
	var editsBuf strings.Builder

	// step 1
	if p == "" && q == "" {
		return gapFreePBuf.String(), gapFreeQBuf.String(), editsBuf.String()
	}

	// step 2
	if p[0] != '-' && q[0] != '-' {
		gapFreePBuf.WriteString(p[0:1])
		gapFreeQBuf.WriteString(q[0:1])
		editsBuf.WriteString("M")
	}

	//step 3 + 4
	for index := 1; index < len(p); index++ {
		if p[index] == '-' {
			gapFreeQBuf.WriteString(q[index : index+1])
			editsBuf.WriteString("I")
			continue
		}
		if q[index:index+1] == "-" {
			gapFreePBuf.WriteString(p[index : index+1])
			editsBuf.WriteString("D")
			continue
		}
		gapFreePBuf.WriteString(p[index : index+1])
		gapFreeQBuf.WriteString(q[index : index+1])
		editsBuf.WriteString("M")
	}

	return gapFreePBuf.String(), gapFreeQBuf.String(), editsBuf.String()
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

func LocalAlign_helper(p, x string, i int, edits string) (pRow, xRow string) {
	pRow, xRow = "", ""
	// Align p and q based on edits
	pRow, xRow = Align_help(p, x[i:], edits)
	return pRow, xRow
}
func Align_help(p, q, edits string) (pRow, qRow string) {
	pRow, qRow = "", ""
	pItr, qItr := 0, 0
	for pos, char := range edits {
		char := string(char)

		pEnd := pItr + 1
		qEnd := qItr + 1
		if char == "M" {
			pRow += p[pItr:pEnd]
			qRow += q[qItr:qEnd]
			pItr += 1
			qItr += 1
		}

		if char == "I" {
			pRow += "-"
			qRow += q[qItr:qEnd]
			qItr += 1
		}

		if char == "D" {
			pRow += p[pos:pEnd]
			qRow += "-"
			pItr += 1
		}
	}
	// Align p and q based on edits
	return pRow, qRow
}
func EditDist(p, x string, i int, edits string) int {

	// IDK I CANT IMPORT LocalAlign AAAAAH
	pRow, xRow := LocalAlign_helper(p, x, i, edits)

	count := 0
	for i := range pRow {

		if pRow[i] != xRow[i] {
			count += 1
		}
	}

	return count
}
