---
title: usd_name
order: 96
---
| Since | 17.5 |
| --- | --- |

`string  usd_name(<stage>stage, string primpath)`

This function returns the name of the given primitive.

Note, while this function takes the stage as an argument for consistency, it does not access the stage, but rather it extracts the name from the path.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

Returns

The name of the primitive.

Examples

## examples

```vex
// Get the primitive name, ie "cube".
string name = usd_name(0, "/geo/cube");

```
