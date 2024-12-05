---
title: usd_parentpath
order: 97
---
| Since | 17.5 |
| --- | --- |

`string  usd_parentpath(<stage>stage, string primpath)`

This function returns the path of the given primitive’s parent.

Note, while this function takes the stage as an argument for consistency, it does not access the stage, but rather it extracts the parent from the path.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

Returns

The path of the primitive’s parent.

Examples

## examples

```vex
// Get the path of the primitive's parent, ie "/geo".
string path = usd_parentpath(0, "/geo/cube");

```
