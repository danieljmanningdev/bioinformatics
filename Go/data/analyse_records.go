//
// Daniel J. Manning
//
// Copyright © 2026 Daniel J. Manning.
// SPDX-License-Identifier: MIT
//

package data

import "bioinformatics/molecules"

type RecordAnalysis struct {
	ID          string
	Sequence    string
	Nucleotides [4]int
	GCContent   float64
	Motif       string
	Positions   []int
	RNA         string
	Protein     string
}

func AnalyseRecord(id, sequence, motif string) RecordAnalysis {
	nucleotides := molecules.CountNucleotides(sequence)
	gc := molecules.GCContent(sequence)
	positions := molecules.FindMotifInDNA(sequence, motif)

	rna := molecules.Transcribe(sequence)
	protein := molecules.RNAToProtein(rna)

	return RecordAnalysis{
		ID:          id,
		Sequence:    sequence,
		Nucleotides: nucleotides,
		GCContent:   gc,
		Motif:       motif,
		Positions:   positions,
		RNA:         rna,
		Protein:     protein,
	}
}
