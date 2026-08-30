//
// Daniel J. Manning
//
// Copyright © 2026 Daniel J. Manning.
// SPDX-License-Identifier: MIT
//

package molecules

import "strings"

/*
   Transcribe a DNA sequence into RNA by replacing thymine with uracil.

   For each nucleotide x:

       f(x) = U, if x = T
            = x, otherwise

   Alphabets:
       Σ_DNA = {A, C, G, T}
       Σ_RNA = {A, C, G, U}

   Transformation:
       T → U
*/

func Transcribe(t string) string {
	return strings.ReplaceAll(t, "T", "U")
}
