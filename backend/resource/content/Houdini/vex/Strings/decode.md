---
title: decode
order: 4
---
| Since | 17.5 |
| --- | --- |

`string  decode(string str)`

Houdini VEX variable names are only allowed to contain letters, numbers, and
underscores, and must not begin with a number. Arbitrary strings can be passed
through the `encode` function to generate a string that obeys these
restriction. This function takes one of these encoded strings, and returns the
original string. A string that has not been encoded will be returned
unmodified.
