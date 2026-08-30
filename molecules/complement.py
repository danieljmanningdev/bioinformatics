def reverse_complement(s: str) -> str:
    """
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

    Args:
        s: DNA sequence.

    Returns:
        Reverse complement of s.
    """
    table = str.maketrans("ATCG", "TAGC")
    return s.translate(table)[::-1]