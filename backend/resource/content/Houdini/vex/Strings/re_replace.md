---
title: re_replace
order: 27
---
`string  re_replace(string regex, string replacement, string input, int maxreplace=0)`

Returns a string where each non-overlapping match of `regex` is replaced with `replacement`. The `replacement` string can reference captured groups from the regex using `$1` syntax.

If `maxreplace` is given and not 0, it specifies the maximum number of replacements to do.
