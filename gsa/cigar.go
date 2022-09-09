// You can create modules at this level and they will be
// interpreted as under module birc.au.dk, so to import
// package `gsa` you need `import "birc.au.dk/gsa"`

package gsa

import (
	"strconv"
	"strings"
)

// Expand the compressed CIGAR encoding into the full list of edits.
//
//  Args:
//      cigar: A CIGAR string
//
//  Returns:
//      The edit operations the CIGAR string describes.
func CigarToEdits(cigar string) (edits string) {
	edits = ""
	for pos, char := range cigar {
		char := string(char)

		//check if string is a number
		if num, err := strconv.Atoi(char); err == nil {
			edits += strings.Repeat(string(cigar[pos+1]), num)
		}
	}
	return edits
}

// Encode a sequence of edits as a CIGAR.
//
//  Args:
//      edits: A sequence of edit operations
//
//  Returns:
//      The CIGAR encoding of edits.
func EditsToCigar(edits string) (cigar string) {
	if edits == "" {
		return ""
	}
	cigar = ""
	counter := 1
	last := string(edits[0])

	for _, char := range edits[1:] {
		char := string(char)
		if char == last {
			counter += 1
		} else {
			cigar += strconv.Itoa(counter) + last
			counter = 1
		}
		last = char
	}
	cigar += strconv.Itoa(counter) + last
	return cigar
}
