//
// Daniel J. Manning
//
// Copyright © 2026 Daniel J. Manning.
// SPDX-License-Identifier: MIT
//

package molecules

import (
	"bioinformatics/input"
	"strings"
)

/*
   Calculate the GC-content of a DNA sequence.

   Let the DNA alphabet be:

       Σ = {A, C, G, T}

   For a DNA sequence s of length n, define:

       GC(s) = ((N_G(s) + N_C(s)) / n) × 100

   where:
       N_G(s) = number of G nucleotides in s
       N_C(s) = number of C nucleotides in s
       n = |s|
*/

func GCContent(s string) float64 {
	gc := strings.Count(s, "G") + strings.Count(s, "C")

	return (float64(gc) / float64(len(s))) * 100
}

/*
   Find the FASTA record with the greatest GC-content.

   For FASTA records R = {r₁, r₂, ..., rₘ}, find:

       r* = arg max GC(s_r)
            r ∈ R

   where:
       s_r is the DNA sequence belonging to record r.
*/

func HighestGC(data string) (string, float64) {
	records := input.ParseFASTA(data)
	highestID := ""
	highestValue := 0.0

	for recordID, sequence := range records {
		gc := GCContent(sequence)

		if gc > highestValue {
			highestValue = gc
			highestID = recordID
		}
	}

	return highestID, highestValue
}
