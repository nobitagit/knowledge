# Algorithms

## Measuring

2 algos can be implemented the same, but execute differently because of hardware differences.
Or we may measure only the perf by running 1 or 2 samples.

To get a better insight on the perf of an algo, we want to analyse its curve when the input size (N) increases.

(y) execution time
(x) input size (N)

![big o notation](./images/bigo.png)

- Constant: O(1) v good
- Linear: O(n) generally good
- Logarithmic: O(log n) also good
- Exponential: O(n2) poor perf
