### Description

This salution calculates the number of ways to decode number string to letters given the following mappings:
Write a function or algorithm that takes a string of digits and returns the number of ways it
can be decoded back into its original message.

'A' -> 1

'B' -> 2

...

'Z' -> 26

### Solution
The solution accumulates the number of ways to decode for each of the two possible previous ways
to find out the final number.

### Complexity
- Algorithmic complexity: O(n)
- Memory complexity: O(1)

## Run Tests

Unit tests
```sh
go test -v
```

Benchmark
```sh
go test -v -bench="."
```# num2letter
# num2letter
