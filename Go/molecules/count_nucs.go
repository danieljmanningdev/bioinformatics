//
// Daniel J. Manning
//
// Copyright © 2026 Daniel J. Manning.
// SPDX-License-Identifier: MIT
//

package molecules

// CountNucleotides counts the occurrences of A, C, G, and T in a DNA sequence.
//
// For each nucleotide X ∈ {A, C, G, T}:
//
//	N_X(s) = Σᵢ 1[sᵢ = X]
//
// Returns counts in the order A, C, G, T.
func CountNucleotides(s string) [4]int {
	counts := make(map[rune]int)

	for _, nucleotide := range s {
		counts[nucleotide]++
	}

	return [4]int{
		counts['A'],
		counts['C'],
		counts['G'],
		counts['T'],
	}
}
