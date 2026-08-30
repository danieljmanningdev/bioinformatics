def count_nucs(s: str) -> list[int]:
    """
    Count the occurrences of A, C, G, and T in a DNA sequence.

    For each nucleotide X ∈ {A, C, G, T}:

        N_X(s) = Σᵢ 1[sᵢ = X]

    Args:
        s: DNA sequence to analyse.

    Returns:
        Counts of A, C, G, and T in that order.
    """
    a = 0
    c = 0
    g = 0
    t = 0

    for nucleotide in s:
        if nucleotide == "A":
            a += 1
        elif nucleotide == "C":
            c += 1
        elif nucleotide == "G":
            g += 1
        elif nucleotide == "T":
            t += 1

    return [a, c, g, t]