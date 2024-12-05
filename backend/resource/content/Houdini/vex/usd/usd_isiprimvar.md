---
title: usd_isiprimvar
order: 73
---
| Since | 19.0 |
| --- | --- |

`int  usd_isiprimvar(<stage>stage, string primpath, string name)`

This function checks whether the primitive or its ancestor has a primvar of the
given name.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

Primvar name (without namespace).

Returns

`1` if the primvar exists directly on the given primitive or on its ancestor, or `0` otherwise.

Examples

## examples

```vex
// Check if the sphere primitive or its ancestor has a primvar "some_primvar".
int is_primvar = usd_isiprimvar(0, "/geometry/sphere", "some_primvar");

```
