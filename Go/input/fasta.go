//
// Daniel J. Manning
//
// Copyright © 2026 Daniel J. Manning.
// SPDX-License-Identifier: MIT
//

package input

import "strings"

/*
   Parse FASTA-formatted text into a mapping of sequence IDs to DNA sequences.

   Each record has the form:

       >id
       s

   where s ∈ Σ^n and Σ = {A, C, G, T}.
*/

func ParseFASTA(data string) map[string]string {
	records := map[string]string{}
	currentID := ""

	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, ">") {
			currentID = line[1:]
			records[currentID] = ""
		} else if currentID != "" {
			records[currentID] += line
		}
	}

	return records
}
