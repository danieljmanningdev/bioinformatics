def transcribe(t:str) -> str:
    """
    Transcribe a DNA sequence into RNA by replacing thymine with uracil.

    For each nucleotide x:

        f(x) = U, if x = T
             = x, otherwise

    Alphabets:
        Σ_DNA = {A, C, G, T}
        Σ_RNA = {A, C, G, U}

    Transformation:
        T → U

    Args:
        t: DNA sequence to transcribe.

    Returns:
        RNA sequence with each T replaced by U.
    """
    return t.replace("T", "U")