# README



## Results

```bash
$ go test -bench=. -benchmem       
goos: darwin
goarch: arm64
pkg: github.com/apatniv/vivek-akupatni-blog/2026/2026-09-13/golang
cpu: Apple M4 Pro
BenchmarkInserts/size=1000-14      	   34161	     33310 ns/op	   74264 B/op	      20 allocs/op
BenchmarkInserts/size=2000-14      	   17985	     67006 ns/op	  148153 B/op	      29 allocs/op
BenchmarkInserts/size=4000-14      	    8811	    136442 ns/op	  295928 B/op	      46 allocs/op
BenchmarkInserts/size=8000-14      	    4288	    280159 ns/op	  591456 B/op	      79 allocs/op
BenchmarkInserts/size=16000-14     	    2082	    577972 ns/op	 1182586 B/op	     144 allocs/op
BenchmarkInserts/size=32000-14     	     985	   1222532 ns/op	 2364799 B/op	     273 allocs/op
BenchmarkInserts/size=64000-14     	     478	   2501156 ns/op	 4729360 B/op	     530 allocs/op
BenchmarkInserts/size=128000-14    	     241	   4980292 ns/op	 9458454 B/op	    1043 allocs/op
BenchmarkInsertsPreAllocated/size=1000-14         	  121494	      9963 ns/op	   36944 B/op	       5 allocs/op
BenchmarkInsertsPreAllocated/size=2000-14         	   59919	     19932 ns/op	   73888 B/op	       9 allocs/op
BenchmarkInsertsPreAllocated/size=4000-14         	   27537	     43669 ns/op	  147776 B/op	      17 allocs/op
BenchmarkInsertsPreAllocated/size=8000-14         	   12722	     94332 ns/op	  295552 B/op	      33 allocs/op
BenchmarkInsertsPreAllocated/size=16000-14        	    6007	    199731 ns/op	  591106 B/op	      65 allocs/op
BenchmarkInsertsPreAllocated/size=32000-14        	    2818	    428386 ns/op	 1182214 B/op	     129 allocs/op
BenchmarkInsertsPreAllocated/size=64000-14        	    1358	    873989 ns/op	 2364560 B/op	     257 allocs/op
BenchmarkInsertsPreAllocated/size=128000-14       	     692	   1732308 ns/op	 4729109 B/op	     513 allocs/op
PASS

```



| n      | Time Taken (µs) [No Preallocation] & Factor |
| ------ | ------------------------------------------- |
| 1000   | 33.31 (1)                                   |
| 2000   | 67.01 (2.01)                                |
| 4000   | 136.44 (2.04)                               |
| 8000   | 280.16 (2.05)                               |
| 16000  | 577.97 (2.06)                               |
| 32000  | 1222.53 (2.12)                              |
| 64000  | 2501.16 (2.05)                              |
| 128000 | 4980.3 (1.99)                               |

