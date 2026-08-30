"""
Rosalind: Rabbits and Recurrence Relations.

Model rabbit population growth using a modified Fibonacci recurrence.
"""


def rabbit_pairs(n: int, k: int) -> int:
    """
    Calculate the total number of rabbit pairs after n months.

    Recurrence relation:

        F_n = F_(n-1) + kF_(n-2)

    Initial conditions:

        F_1 = 1
        F_2 = 1

    Args:
        n: Number of months.
        k: Number of rabbit pairs produced by each mature pair.

    Returns:
        Total number of rabbit pairs after n months.
    """
    previous = 1
    two_months_ago = 1

    for _ in range(3, n + 1):
        current = previous + k * two_months_ago

        two_months_ago = previous
        previous = current

    return previous