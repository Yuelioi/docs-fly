---
title: replace
order: 31
---
`string  replace(string str, string old, string new)`

`string  replace(string str, string old, string new, int count)`

Returns a copy of the string with all occurrences of the string `old` replaced with the string `new`.

Show/hide arguments

`count`

If specified, only the first `count` occurrences are replaced.

Examples

## examples

```vex
string str = "abcdef abcdef abcdef";

// Returns "abcghi abcghi abcghi"
string new_str = replace(str, "def", "ghi");

// Replaces up to 2 occurrences of the string "def".
// Returns "abcghi abcghi abcdef"
new_str = replace(str, "def", "ghi", 2);

```
