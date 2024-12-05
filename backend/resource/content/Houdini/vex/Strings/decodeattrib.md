---
title: decodeattrib
order: 5
---
| Since | 18.0 |
| --- | --- |

`string  decodeattrib(string str)`

Houdini geometry attributes and group names are only allowed to contain
letters, numbers, and underscores, and must not begin with a number. Arbitrary
strings can be passed through the `encodeattrib` function to generate a string
that obeys these restriction. This function takes one of these encoded
strings, and returns the original string. A string that has not been encoded
will be returned unmodified.
