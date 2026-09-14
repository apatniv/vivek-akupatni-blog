# README

## Source of the codebase
Source of the codebase: [Daniel Lemire's blog on Python sets and dictionaries can have quadratic-time performance](https://lemire.me/blog/2026/09/03/python-sets-and-dictionaries-can-have-quadratic-time-performance)



## Benchmark results on Apple M4 pro

```bash
$ python3 --version
Python 3.14.4

$ python3 benchmark.py     
Website one
n=1000 taken=9.44ms
n=2000 taken=28.78ms
n=4000 taken=127.24ms
n=8000 taken=475.59ms
n=16000 taken=2079.17ms
n=32000 taken=8023.70ms
n=64000 taken=34462.33ms
```



| n     | Time Taken (ms) | Factor |
| ----- | --------------- | ------ |
| 1000  | 9.44            | 1.00   |
| 2000  | 28.78           | 3.05   |
| 4000  | 127.24          | 4.42   |
| 8000  | 475.59          | 3.74   |
| 16000 | 2079.17         | 4.37   |
| 32000 | 8023.70         | 3.86   |
| 64000 | 34462.33        | 4.30   |

> Doubling the size increases the time taken by roughly 4 times. Suggesting, that actual complexity of the operation(s) to be $O(n^2)$ rather than 0(n).