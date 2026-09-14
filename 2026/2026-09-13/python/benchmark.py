"""
Source of the codebase: Daniel Lemire's blog
Reference:https://lemire.me/blog/2026/09/03/python-sets-and-dictionaries-can-have-quadratic-time-performance

"""

import random
import statistics
import time

REPEATS = 5
KEY_LENGTH = 16


def time_it(fn, *args):
    runs = []
    for _ in range(REPEATS):
        start = time.perf_counter()
        fn(*args)
        runs.append(time.perf_counter() - start)
    return statistics.median(runs)


def set_operations(values):
    s = set(values)  # insertions
    count = sum(v in s for v in values)  # checks
    return count


def website_snippet(sizes):
    print(f"Website one")
    M = (1 << 61) - 1
    for n in sizes:
        values = [i * M for i in range(1, n + 1)]
        taken = time_it(set_operations, values)
        print(f"n={n} taken={taken*1000:.2f}ms")


if __name__ == "__main__":
    rng = random.Random(1234)
    sizes = [1_000, 2_000, 4_000, 8_000, 16_000, 32_000, 64_000]
    website_snippet(sizes)
