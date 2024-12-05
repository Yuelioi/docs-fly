---
title: makevalidvarname
order: 19
---
| Since | 19.5 |
| --- | --- |

`string  makevalidvarname(string name)`

`string  makevalidvarname(string name, string safe_chars)`

Variable names in languages such as VEX are only allowed to contain letters, numbers, and underscores, and must not begin with a number.
Node names and attribute names in Houdini have similar requirements.
This function takes any string, and converts it into a string that obeys these restrictions by replacing invalid characters with an underscore.

Show/hide arguments

`name`

String that should be turned into a valid variable name.

`safe_chars`

String specifying any extra characters to allow instead of replacing with underscores.

Examples

## examples

```vex
// Returns "foo_bar"
string s = makevalidvarname("foo:bar");

// Returns "_123"
s = makevalidvarname("123");

// Returns "foo:_bar"
s = makevalidvarname("foo:?bar", ":");

```
