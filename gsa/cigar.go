// You can create modules at this level and they will be
// interpreted as under module birc.au.dk, so to import
// package `gsa` you need `import "birc.au.dk/gsa"`

package gsa

// Expand the compressed CIGAR encoding into the full list of edits.
//
//  Args:
//      cigar: A CIGAR string
//
//  Returns:
//      The edit operations the CIGAR string describes.
func CigarToEdits(cigar string) (edits string) {
	edits = ""
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
	cigar = ""
	return cigar
}
