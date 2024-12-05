---
title: usd_primvarnames
order: 112
---
| Since | 18.0 |
| --- | --- |

`string [] usd_primvarnames(<stage>stage, string primpath)`

This function returns the primvar names that are available on the given primitive.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

Returns

String array containing the names of the primitive’s primvars.

Examples

## examples

```vex
// Get the primvar names from the primitive.
string primvar_names[] = usd_primvarnames(0, "/geo/src_sphere");

```
