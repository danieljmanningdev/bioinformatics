//
// Daniel J. Manning
//
// Copyright © 2026 Daniel J. Manning.
// SPDX-License-Identifier: MIT
//

package molecules

/*
	Count the number of positions where two equal-length DNA sequences differ.

    d_H(s, t) = Σᵢ 1[sᵢ ≠ tᵢ]
*/

func HammingDistance(s, t string) int {
	count := 0

	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			count++
		}
	}

	return count
}
