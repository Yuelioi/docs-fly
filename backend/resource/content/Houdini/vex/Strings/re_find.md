---
title: re_find
order: 24
---
`int  re_find(string regex, string input)`

`int  re_find(string regex, string input, int start)`

`int  re_find(string regex, string input, int start, int end)`

Returns `1` if `regex` matches in `input`, or `0` otherwise.

`string  re_find(string regex, string input)`

`string  re_find(string regex, string input, int start)`

`string  re_find(string regex, string input, int start, int end)`

Returns the first substring that matches `regex` in `input`.

`int  re_find(int &start_pos[], int &end_pos[], string regex, string input)`

`int  re_find(int &start_pos[], int &end_pos[], string regex, string input, int start)`

`int  re_find(int &start_pos[], int &end_pos[], string regex, string input, int start, int end)`

Returns `1` if `regex` matches in input, or `0` otherwise. Fills the `start_pos` and `end_pos` array variables with the start and end positions of each match.

`string [] re_find(string regex, string input)`

`string [] re_find(string regex, string input, int start)`

`string [] re_find(string regex, string input, int start, int end)`

Returns an array of substrings that match `regex` in `input`.
