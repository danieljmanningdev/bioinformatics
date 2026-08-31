//
// Daniel J. Manning
//
// Copyright © 2026 Daniel J. Manning.
// SPDX-License-Identifier: MIT
//

package molecules

/*
Locate all occurrences of a nucleotide motif within a DNA sequence.

DNA alphabet:

	Σ_DNA = {A, C, G, T}

Let:

	s ∈ Σ_DNA*    be the full DNA sequence
	t ∈ Σ_DNA*    be the target motif

A motif occurrence exists at position i when:

	s[i : i+|t|] = t

where:

	|t| = length of the motif

Rosalind defines sequence positions using 1-based indexing,
whereas Go uses 0-based indexing.

Therefore:

	position = i + 1

The search advances one nucleotide at a time so that
overlapping motif occurrences are also detected.

Example:

	DNA:
	    GATATATGCATATACTT

	Motif:
	    ATAT

	Matches:
	    position 2
	    position 4
	    position 10

	Result:
	    {2, 4, 10}
*/

func FindMotifInDNA(dna string, motif string) []int {
	positions := []int{}

	for i := 0; i <= len(dna)-len(motif); i++ {
		// Take a chunk of DNA the same size as t
		chunk := dna[i : i+len(motif)]

		// If that's the chunk being searched for
		// Remember the position it was found
		if chunk == motif {
			positions = append(positions, i+1)
		}
	}

	return positions
}
