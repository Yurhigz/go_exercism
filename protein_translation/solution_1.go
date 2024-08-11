package protein

import (
    "errors"
)
var ErrInvalidBase = errors.New("ErrInvalidBase")
var ErrStop = errors.New("ErrStop")

var codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine",
	"UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	
    var codons []string
    
    for i := 0 ; i <= len(rna) - 3 ; i+= 3 {
        aminoAcid, ok := FromCodon(rna[i:i+3])
        if ok == ErrInvalidBase   {
            return codons, ErrInvalidBase
        } else if ok == ErrStop {
            return codons, nil
        } 
        codons = append(codons,aminoAcid)
    }
    return codons,nil
}

func FromCodon(codon string) (string, error) {
    aminoAcid, ok := codons[codon]
    if !ok {
        return "", ErrInvalidBase
    } else if aminoAcid == "STOP" {
        return "", ErrStop
    } else {
        return aminoAcid, nil
    }
}
