## Ruby quirks

These quirks have a more thorough explanation in the readme

- `and` and `or` are not 100% like `||` and `&&` (they have different precedence)
- conditional assignment via `||=` has quirks
- monkey patching stdlib methods is possible
- classes can be extended (see open classes)

## Cargo Operator

This is nice. Check which side of the expression has a higher value.

```rb
1 <=> 3 # -1
2 <=> 2 # 0
4 <=> 1 # 1
```
