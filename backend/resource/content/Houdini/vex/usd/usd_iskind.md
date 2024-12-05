---
title: usd_iskind
order: 74
---
| Since | 18.0 |
| --- | --- |

`int  usd_iskind(<stage>stage, string primpath, string kind)`

This function checks whether the given primitive is of a given kind, e.g., an
assembly, component, etc.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`kind`

The name of a kind to check.

Returns

1 if the primitive is of a given kind, and 0 otherwise.

Examples

## examples

```vex
// Check if the sphere primitive is of an assembly kind.
int is_assembly = usd_iskind(0, "/geometry/sphere", "assembly");

```
