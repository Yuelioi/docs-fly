---
title: split
order: 34
---
`string [] split(string s)`

`string [] split(string s, string separators)`

`string [] split(string s, string separators, int maxsplits)`

Splits a string into tokens by removing separator characters from the string
and creating an array entry for each substring bounded by separators. When no
separator string is provided, the string is split on whitespace (spaces, tab,
and return).

The `maxsplits` option limits the number of times the string is split,
this is useful to peel off one token at a time from a larger string.

Note
This differs from Python’s split() in that it takes a list of separators, not a single string to use as a separator.

Note
This operates more as a tokenize method than as a split method.
In particular, if you have repeated separators they will be merged
and only a single split performed.
