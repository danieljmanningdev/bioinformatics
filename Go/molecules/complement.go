//
// Daniel J. Manning
//
// Copyright © 2026 Daniel J. Manning.
// SPDX-License-Identifier: MIT
//

package molecules

/*
 	Return the reverse complement of a DNA sequence.

    Complement mapping:
        c(A) = T
        c(T) = A
        c(C) = G
        c(G) = C

    For a DNA sequence:
        s = s₁s₂...sₙ

    the reverse complement is:
        sᶜ = c(sₙ)c(sₙ₋₁)...c(s₁)
*/

func ReverseComplement(s string) string {
	table := map[byte]byte{
		'A': 'T',
		'T': 'A',
		'C': 'G',
		'G': 'C',
	}

	result := make([]byte, 0, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, table[s[i]])
	}

	return string(result)
}
