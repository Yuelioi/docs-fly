---
title: atoi
order: 2
---
`int  atoi(string str)`

`int  atoi(string str, int base)`

Converts the string argument to an integer value. If a base between 2 and 36 inclusive is supplied, the string is converted in base `base`.

- This function ignores whitespace around the number.
- Returns `0` if the string does not contain a number, or if an invalid base is specified.
- The string can contain exponential notation (for example `"1.25e+5"`).
