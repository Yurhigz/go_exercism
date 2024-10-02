package dna
import (//"strings"
		"errors"
        )

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
    errors := errors.New("This is not a nucleotide.")
    h  := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
    for _,letter := range d {
        switch letter {
            case 'A','C','G','T':
            	h[letter] += 1
            default: 
            	return h,errors
        }
    }
	return h, nil
}





package dna

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 5fc501b Remove unneeded nesting (#1798)

var testCases = []struct {
	description   string
	strand        string
	expected      Histogram
	errorExpected bool
}{
	{
		description:   "empty strand",
		strand:        "",
		errorExpected: false,
		expected:      Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0},
	},
	{
		description:   "can count one nucleotide in single-character input",
		strand:        "G",
		errorExpected: false,
		expected:      Histogram{'A': 0, 'C': 0, 'G': 1, 'T': 0},
	},
	{
		description:   "strand with repeated nucleotide",
		strand:        "GGGGGGG",
		errorExpected: false,
		expected:      Histogram{'A': 0, 'C': 0, 'G': 7, 'T': 0},
	},
	{
		description:   "strand with multiple nucleotides",
		strand:        "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC",
		errorExpected: false,
		expected:      Histogram{'A': 20, 'C': 12, 'G': 17, 'T': 21},
	},
	{
		description:   "strand with invalid nucleotides",
		strand:        "AGXXACT",
		errorExpected: true,
		expected:      Histogram{},
	},
}
