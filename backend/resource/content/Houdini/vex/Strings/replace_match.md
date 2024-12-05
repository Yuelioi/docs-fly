---
title: replace_match
order: 32
---
`string  replace_match(string str, string pattern_from, string pattern_to)`

If the string matches `pattern_from`, it is replaced with `pattern_to` with the matching wildcards substituted in.

The pattern can use wildcards such as `str*` or `str?`, similar to the [match](match.html "This function returns 1 if the subject matches the pattern specified,
or 0 if the subject doesn’t match.") function.
The wildcards may also be referred to with an index (e.g. `(2)`) to allow reordering.
Examples

## examples

```vex
// Returns "carol is my name";
string s = replace_match("bob is my name", "bob*", "carol*");

// Returns "a-b";
s = replace_match("a_to_b", "*_to_*", "*-*");

// Swaps the matched wildcards, returning "b_to_a";
s = replace_match("a_to_b", "*_to_*", "*(1)_to_*(0)");

```
