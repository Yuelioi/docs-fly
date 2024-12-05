---
title: re_split
order: 28
---
`string [] re_split(string regex, string input, int maxsplits=0)`

Splits the `input` string into substrings between matches of `regex`.

If `maxsplits` is given and not 0, it specifies the maximum number of splits to do.

NOTE: This does not match Python and other split methods as it behaves
differently with trailing delimiters. `a,b,` will only split into two
strings. The regex string can be augmented with `|($(?!\s))` to split
three tokens. In VEX you frequently will need \\s. In a wrangle another
escape is needed and \\\\s is required.
