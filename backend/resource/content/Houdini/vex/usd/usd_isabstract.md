---
title: usd_isabstract
order: 61
---
| Since | 19.0 |
| --- | --- |

`int  usd_isabstract(<stage>stage, string primpath)`

This function checks whether the given primitive is abstract or defined.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

Returns

1 if the primitive is abstract, and 0 otherwise.

Examples

## examples

```vex
// Check if the sphere primitive is abstract.
int is_abstract = usd_isabstract(0, "/geometry/sphere");

```
