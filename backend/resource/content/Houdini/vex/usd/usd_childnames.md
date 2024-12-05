---
title: usd_childnames
order: 28
---
| Since | 19.5 |
| --- | --- |

`string [] usd_childnames(<stage>stage, string primpath)`

This function returns the names of the children authored directly in the namespace of the given parent primitive.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

Returns

The names of the children of the given primitive.

Examples

## examples

```vex
// Get the children of the "mother" primitive.
string children[] = usd_child(0, "/geo/mother");

```
