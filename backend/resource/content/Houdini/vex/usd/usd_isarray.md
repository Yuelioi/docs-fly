---
title: usd_isarray
order: 63
---
| Since | 17.5 |
| --- | --- |

`int  usd_isarray(<stage>stage, string primpath, string name)`

This function checks whether the given attribute is an array.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

Attribute name.

Returns

`1` if the attribute exists and is an array, or `0` otherwise.

Examples

## examples

```vex
// Check if attribute "some_attribute" is an array.
int is_array = usd_isarray(0, "/geometry/sphere", "some_attribute");

```
