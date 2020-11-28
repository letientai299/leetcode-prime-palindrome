# Leetcode palindrome prime fasted solutions

Several solutions with 0ms execution time and benchmarks for
[prime-palindrome](https://leetcode.com/problems/prime-palindrome/) problem on
leetcode.

Benchmark result

```
Benchmark_primePalindrome/letientai299-12         	   20000	     95646 ns/op	    1008 B/op	      29 allocs/op
Benchmark_primePalindrome/unknown-12              	    5000	    252136 ns/op	       0 B/op	       0 allocs/op
Benchmark_primePalindrome/tjucoder-12             	     500	   2118324 ns/op	  524816 B/op	   36900 allocs/op
```

- `unknown` guy's solution does zero allocation, and cool trick to generate
  palindrome with odds number of digits, but his solution need to check many
  more palindrome numbers than mine.
- `tjucoder`'s
  [solution](https://leetcode.com/problems/prime-palindrome/discuss/803639/Go-0ms-solution)
  works with `string`, thus require many more allocation and also his
  palindrome generation is not effective.
- Mine is currently the fastest thanks to: generate potental palindrome as list
  of `int` digits (so less allocation and conversion), filter out those
  palindrome that can't be prime.

If you have faster solutions, feel free to submit an MR. Note that:

- `big.NewInt().ProbablyPrime()` has terrible performance in the problem space
  (`int`). I got 4512 ms before switch to simple prime test up to `sqrt(n)`.

## Disclaimer

- I finish my solution without refer to either others solutions.
- This's for fun only (because I write quite a lot of code for this
  problems, so I want to know if it worth it). These code're not for production
  usage and I don't know if the performance of those solutions are still in above
  order if the problem space is bigger (from `int` to `uint64` or `big.Int`).
