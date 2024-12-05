---
title: usd_attribnames
order: 19
---
| Since | 17.5 |
| --- | --- |

`string [] usd_attribnames(<stage>stage, string primpath)`

This function returns the attribute names that are available on the given primitive.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

Returns

String array containing the names of the primitive’s attributes.

Examples

## examples

```vex
// Get the attribute names from the primitive.
string attrib_names[] = usd_attribnames(0, "/geo/sphere");

```
