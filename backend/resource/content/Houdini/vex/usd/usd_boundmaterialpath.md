---
title: usd_boundmaterialpath
order: 27
---
| Since | 18.0 |
| --- | --- |

`string  usd_boundmaterialpath(<stage>stage, string primpath)`

This function returns the material path for the given primitive. May return an
empty string if no material is bound.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

Returns

The material bound to the given primitive.

Examples

## examples

```vex
// Get the sphere primitive's material.
string matpath = usd_boundmaterialpath(0, "/geo/sphere");

```
