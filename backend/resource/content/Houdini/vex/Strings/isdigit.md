---
title: isdigit
order: 15
---
`int  isdigit(string str)`

Returns 1 if all the characters in the string are numeric (0-9). No alphabetic characters,
punctuation or other special characters. If any are present then returns 0.

If passed the string “543.34”, this function will return 0 because “.” is not a digit.
