---
title: usd_typename
order: 149
---
| Since | 17.5 |
| --- | --- |

`string  usd_typename(<stage>stage, string primpath)`

This function returns the type name of the given primitive.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

Returns

The name of the primitive’s type.

Examples

## examples

```vex
// Get the primitive's type name, eg "Cube".
string type_name = usd_typename(0, "/geo/cube");

```
