from input.fasta import parse_fasta


def gc_content(s: str) -> float:
    """
    Calculate the GC-content of a DNA sequence.

    Let the DNA alphabet be:

        Σ = {A, C, G, T}

    For a DNA sequence s of length n, define:

        GC(s) = ((N_G(s) + N_C(s)) / n) × 100

    where:
        N_G(s) = number of G nucleotides in s
        N_C(s) = number of C nucleotides in s
        n = |s|

    Args:
        s: DNA sequence.

    Returns:
        GC-content of s as a percentage.
    """
    return ((s.count("G") + s.count("C")) / len(s)) * 100


def highest_gc(data: str) -> tuple[str, float]:
    """
    Find the FASTA record with the greatest GC-content.

    For FASTA records R = {r₁, r₂, ..., rₘ}, find:

        r* = arg max GC(s_r)
             r ∈ R

    where:
        s_r is the DNA sequence belonging to record r.

    Args:
        data: FASTA-formatted DNA records.

    Returns:
        A tuple containing:
            - the ID of the record with maximum GC-content
            - its GC-content percentage
    """
    records = parse_fasta(data)

    highest_id = ""
    highest_value = 0.0

    for record_id, sequence in records.items():
        gc = gc_content(sequence)

        if gc > highest_value:
            highest_value = gc
            highest_id = record_id

    return highest_id, highest_value