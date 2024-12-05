---
title: decodeparm
order: 6
---
| Since | 18.0 |
| --- | --- |

`string  decodeparm(string str)`

Houdini parameter names are only allowed to contain letters, numbers, hash
characters (for multiparms), and underscores, and must not begin with a
number. Arbitrary strings can be passed through the `encodeparm` function to
generate a string that obeys these restriction. This function takes one of
these encoded strings, and returns the original string. A string that has not
been encoded will be returned unmodified.
