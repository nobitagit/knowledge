# Algorithms

## Measuring

2 algos can be implemented the same, but execute differently because of hardware differences.
Or we may measure only the perf by running 1 or 2 samples.

To get a better insight on the perf of an algo, we want to analyse its curve when the input size (N) increases.

(y) execution time
(x) input size (N)

## Big 0

![big o notation](./images/bigo.png)

- Constant: O(1) v good
- Linear: O(n) generally good
- Logarithmic: O(log n) also good
- Exponential: O(n2) poor perf

An equation to describes how the runtime scales with respect to an input.

### Rules of Big O

#### 1. Each step counts as one

```rb
def do_stuff(input)
  a = input + 1 # 1st step
  b = a * SOME_CONSTANT # 2nd step
end
```

#### 2. Drop the constants

#### 3. Different inputs, different variables

For ex.

```rb
def find_intersections(input1, input2)
  ret = 0
  input1.each do | value1 |
    input2.each do | value2 |
       if value1 == value2
         ret = ret + 1
       end
    end
  end
  ret
end
```

Although we have a nested loop, the **inputs are different**, hence the complexity here is `O(a * b)`, not `O(n²)`.

#### 4. Drop non dominant terms

```rb
def do_stuff(i)
  # This is o(n)
  i.each do | v1 |
    ## do something
  end

  # This is o(n²)
  i.each do | v1 |
    i.each do | v2 |
      # do other stuff
    end
  end
```

We drop the first calculation (`O(n)`) as non-dominant, and we keep the second as the complexity of the function.
