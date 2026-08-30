def parse_fasta(data: str) -> dict[str, str]:
    """
    Parse FASTA-formatted text into a mapping of sequence IDs to DNA sequences.

    Each record has the form:

        >id
        s

    where s ∈ Σ^n and Σ = {A, C, G, T}.

    Args:
        data: FASTA-formatted text.

    Returns:
        A dictionary mapping each FASTA record ID to its complete sequence.
    """
    records: dict[str, str] = {}
    current_id: str | None = None

    for line in data.splitlines():
        line = line.strip()

        if line.startswith(">"):
            current_id = line[1:]
            records[current_id] = ""

        elif current_id is not None:
            records[current_id] += line

    return records