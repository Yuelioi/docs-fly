---
title: find
order: 13
---
`int  find(string haystack, string needle)`

`int  find(string haystack, string needle, int start)`

`int  find(string haystack, string needle, int start, int end)`

Returns the position of the first occurrence of the `needle` string within the `haystack` string. You can limit the result to the first substring that starts at or after a `start` position, and at or before an `end` position.

You can find each occurrence in a loop by setting `start` at each iteration to the end of the previous match.

Returns a negative number if the substring is not found.

`int [] find(string haystack, string needle)`

`int [] find(string haystack, string needle, int start)`

`int [] find(string haystack, string needle, int start, int end)`

Returns a list of positions of occurrences of the `needle` string within the `haystack` string. You can limit the results to substrings that start at or after a `start` position, and before an `end` position.

Returns an empty array if the substring is not found.

`int  find(<type>array[], <type>target)`

`int  find(<type>array[], <type>target, int start)`

`int  find(<type>array[], <type>target, int start, int end)`

Returns the position of the first occurrence of the `target` value within the `array`. You can limit the result to the first occurrence at or after a `start` position, and at or before an `end` position.

You can find each occurrence in a loop by setting `start` at each iteration to the next position after the previous match.

Returns a negative number if the target is not found.

`int [] find(<type>array[], <type>target)`

`int [] find(<type>array[], <type>target, int start)`

`int [] find(<type>array[], <type>target, int start, int end)`

Returns a list of positions of occurrences of the `target` value within the `array`. You can limit the results to occurrences at or after a `start` position, and before an `end` position.

- When you specify an `end` position, it means the matching substring must *start* before the `end`.
- The scalar versions return `-len(haystack)-1` to indicate no matches. This value is intended to cause an error if you try to use it as an index into the string/array.
- Searching for the empty string always fails. This differs from Python.
- You cannot use negative indices for the `start` and `end` arguments.
