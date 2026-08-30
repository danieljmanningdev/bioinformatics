def hamming_distance(s: str, t: str) -> int:
    """
    Count the number of positions where two equal-length DNA sequences differ.

    d_H(s, t) = Σᵢ 1[sᵢ ≠ tᵢ]

    Args:
        s: First DNA sequence.
        t: Second DNA sequence.

    Returns:
        Hamming distance between s and t.
    """
    count = 0

    for a, b in zip(s, t):
        if a != b:
            count += 1

    return count