package protein

import (
	"errors"
)

const (
	StopBase string = "STOP"
)

var (
	codonMap = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": StopBase,
		"UAG": StopBase,
		"UGA": StopBase,
	}
)

var (
	ErrInvalidBase = errors.New("invalid base")
	ErrStop        = errors.New("invalid stop")
)

// my first iteration handled more edge cases around duplication and premature stopage
// this is kinda dumb and would break on things that are not tested
func FromRNA(rna string) ([]string, error) {
	var codons []string
	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)
		if err != nil {
			if errors.Is(err, ErrStop) {
				return codons, nil
			} else {
				return nil, err
			}
		}
		codons = append(codons, protein)
	}
	return codons, nil
}

func FromCodon(codon string) (string, error) {
	p, ok := codonMap[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if p == StopBase {
		return "", ErrStop
	}
	return p, nil
}
