---
title: usd_specifier
order: 144
---
| Since | 19.0 |
| --- | --- |

`string  usd_specifier(<stage>stage, string primpath)`

This function returns the given primitive’s specifier, e.g., “def”, “class”, etc.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

Returns

The specifier of the given primitive.

Examples

## examples

```vex
// Get the sphere primitive's specifier.
string specifier = usd_specifier(0, "/geo/sphere");

```
