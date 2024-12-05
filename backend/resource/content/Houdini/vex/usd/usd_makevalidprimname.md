---
title: usd_makevalidprimname
order: 90
---
| Since | 19.0 |
| --- | --- |

`string  usd_makevalidprimname(string name)`

This function ensures that a string meets the requirements of a legal USD
primitive name that can be used as part of a valid primitive path.

Show/hide arguments

`name`

String that should be turned into a valid primitive name.

Returns

A possibly modified version of the original string which conforms to the
requirements for naming USD primitives. Invalid characters such as spaces
and most punctuation will be converted to underscores. Note that this means
the translation is not reversible. If the original string is already a
legal primitive name, the string is returned unchanged.
