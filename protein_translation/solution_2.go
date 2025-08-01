// Package protein has functionality to translate RNA sequences into proteins.
package protein

import (
	"errors"
)

var (
	// ErrStop indicates the translation was stopped.
	ErrStop = errors.New("stop")

	// ErrInvalidBase indicates the codon was invalid.
	ErrInvalidBase = errors.New("invalid codon")

	noProteins = []string{}
)

const codonLength = 3

// FromCodon either translates a codon to a protein or returns an error for a stop codon or invalid codon.
func FromCodon(codon string) (protein string, err error) {
	switch codon {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	case "UAA", "UAG", "UGA":
		err = ErrStop
	default:
		err = ErrInvalidBase
	}

	return protein, err
}

// FromRNA maps RNA codons to their matching proteins and returns either the list of proteins
// or returns an error for an invalid codon.
func FromRNA(rna string) (proteins []string, err error) {
	for i := 0; i < len(rna); i += codonLength {
		protein, err := FromCodon(rna[i : i+codonLength])
		if err == ErrStop {
			break
		}
		if err == ErrInvalidBase {
			return noProteins, err
		}
		proteins = append(proteins, protein)
	}

	return proteins, err
}