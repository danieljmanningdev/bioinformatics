# Bioinformatics

A growing Python bioinformatics toolkit built while working through computational biology problems, with an emphasis on understanding the biology, the mathematics behind each operation, and the corresponding implementation.

The project currently focuses on DNA sequence analysis, FASTA parsing, mutation comparison, transcription, reverse complements, GC-content analysis, and simple biological population models.

## Current Features

### Molecular sequence operations

- Count DNA nucleotides in the order `A C G T`
- Transcribe DNA into RNA (`T → U`)
- Generate the reverse complement of a DNA sequence
- Calculate GC-content
- Find the FASTA record with the highest GC-content
- Calculate Hamming distance between equal-length DNA sequences

### Input parsing

- Parse FASTA-formatted text into a mapping of record IDs to sequences
- Join multi-line FASTA sequences into a single sequence per record

### Models

- Model rabbit population growth using a modified Fibonacci recurrence

## Scientific Notation

Where useful, functions are documented alongside the mathematical notation they implement.

For nucleotide counting:

$$
N_X(s) = \sum_{i=1}^{n} \mathbf{1}_{\{s_i = X\}}
$$

for $X \in \{A,C,G,T\}$.

For GC-content:

$$
GC(s) = \frac{N_G(s) + N_C(s)}{|s|} \times 100
$$

For Hamming distance:

$$
d_H(s,t) = \sum_{i=1}^{n} \mathbf{1}_{\{s_i \ne t_i\}}
$$

For the rabbit population model:

$$
F_n = F_{n-1} + kF_{n-2}, \qquad F_1 = F_2 = 1
$$

## Project Structure

```text
.
├── input/
│   └── fasta.py
├── models/
│   └── fibonacci_rabbits.py
├── molecules/
│   ├── complement.py
│   ├── count_nucs.py
│   ├── gc_content.py
│   ├── hamming_distance.py
│   └── transcription.py
├── main.py
├── LICENSE
└── README.md
```

## Example

```python
from molecules.count_nucs import count_nucs
from molecules.transcription import transcribe
from molecules.complement import reverse_complement

s = "AGCTTTTCATTCTGACTGCA"

print(count_nucs(s))
print(transcribe(s))
print(reverse_complement(s))
```

## Learning Goals

This repository is intended to grow alongside continued study of bioinformatics and computational biology. Rather than keeping each exercise as an isolated script, reusable operations are extracted into small modules that can later form the basis of a broader analysis toolkit.

Areas expected to grow over time include:

- sequence analysis
- mutations and genetic variation
- motif discovery
- additional biological file formats
- statistical and probabilistic models
- larger dataset processing
- visualisation and reporting

## Rosalind

Many of the early implementations are based on problems from the [Rosalind](https://rosalind.info/) bioinformatics platform. The repository contains reusable implementations rather than copies of Rosalind problem statements or datasets.

## License

This project is licensed under the MIT License. See [`LICENSE`](LICENSE) for details.
