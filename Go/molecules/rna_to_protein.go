//
// Daniel J. Manning
//
// Copyright © 2026 Daniel J. Manning.
// SPDX-License-Identifier: MIT
//

package molecules

/*
   Translate an mRNA sequence into its corresponding amino acid sequence
   using the standard genetic code.

   RNA alphabet:
       Σ_RNA = {A, C, G, U}

   Protein alphabet:
       Σ_protein =
       {A, C, D, E, F, G, H, I, K, L,
        M, N, P, Q, R, S, T, V, W, Y}

   Codon:
       c = r_i r_(i+1) r_(i+2)

       where:
           r_i ∈ Σ_RNA

   Translation function:
       τ : Σ_RNA³ → Σ_protein ∪ {STOP}

   The mRNA sequence is partitioned into consecutive codons:

       s = c₁ c₂ c₃ ... cₙ

   Protein synthesis:

       P = τ(c₁) τ(c₂) τ(c₃) ... τ(cₖ)

   where cₖ₊₁ is the first termination codon.

   Termination codons:
       UAA → STOP
       UAG → STOP
       UGA → STOP

   Example:

       RNA:
           AUG GCC UUU GAG UAA

       Translation:
           AUG → M
           GCC → A
           UUU → F
           GAG → E
           UAA → STOP

       Protein:
           MAFE
*/

func RNAToProtein(rna string) string {
	codonTable := map[string]string{
		"UUU": "F", "UUC": "F",
		"UUA": "L", "UUG": "L",
		"UCU": "S", "UCC": "S", "UCA": "S", "UCG": "S",
		"UAU": "Y", "UAC": "Y",
		"UAA": "Stop", "UAG": "Stop",
		"UGU": "C", "UGC": "C",
		"UGA": "Stop", "UGG": "W",

		"CUU": "L", "CUC": "L", "CUA": "L", "CUG": "L",
		"CCU": "P", "CCC": "P", "CCA": "P", "CCG": "P",
		"CAU": "H", "CAC": "H",
		"CAA": "Q", "CAG": "Q",
		"CGU": "R", "CGC": "R", "CGA": "R", "CGG": "R",

		"AUU": "I", "AUC": "I", "AUA": "I",
		"AUG": "M",
		"ACU": "T", "ACC": "T", "ACA": "T", "ACG": "T",
		"AAU": "N", "AAC": "N",
		"AAA": "K", "AAG": "K",
		"AGU": "S", "AGC": "S",
		"AGA": "R", "AGG": "R",

		"GUU": "V", "GUC": "V", "GUA": "V", "GUG": "V",
		"GCU": "A", "GCC": "A", "GCA": "A", "GCG": "A",
		"GAU": "D", "GAC": "D",
		"GAA": "E", "GAG": "E",
		"GGU": "G", "GGC": "G", "GGA": "G", "GGG": "G",
	}

	protein := ""

	for i := 0; i+2 < len(rna); i += 3 {
		codon := rna[i : i+3]
		aminoAcid := codonTable[codon]

		if aminoAcid == "Stop" {
			break
		}

		protein += aminoAcid
	}

	return protein
}
